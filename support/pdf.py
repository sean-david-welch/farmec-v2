import io
import logging
import re
import zipfile
from collections.abc import Generator
from contextlib import contextmanager
from datetime import date
from typing import Callable

from django.contrib import messages
from django.contrib.admin import ModelAdmin
from django.db.models import Model, QuerySet
from django.http import HttpRequest, HttpResponse
from django.template.loader import render_to_string
from django.conf import settings
from weasyprint import CSS, HTML

logger = logging.getLogger(__name__)


@contextmanager
def collect_image_failures() -> Generator[list[str]]:
    """
    Collect WeasyPrint "Failed to load image" log messages raised inside the block.

    WeasyPrint skips images it cannot fetch and logs an error instead of raising.
    The handler is attached to the global ``weasyprint`` logger, which is safe
    under gunicorn's sync workers (one request per process at a time).

    :returns: A list that fills with one message per image that failed to load.
    """
    failures: list[str] = []

    class FailureHandler(logging.Handler):
        def emit(self, record: logging.LogRecord) -> None:
            message: str = record.getMessage()
            if 'Failed to load image' in message:
                failures.append(message)

    handler = FailureHandler(level=logging.WARNING)
    weasyprint_logger = logging.getLogger('weasyprint')
    weasyprint_logger.addHandler(handler)
    try:
        yield failures
    finally:
        weasyprint_logger.removeHandler(handler)


class PDFDownloadAction:
    """
    Reusable Django admin action for generating PDF downloads.

    Instantiate once per model, passing a template path, context builder,
    and filename builder. The instance is callable and can be passed directly
    to a ModelAdmin's ``actions`` list.

    Single record selections produce a PDF attachment; multiple selections
    produce a ZIP archive containing one PDF per record.

    Example::

        download_pdf = PDFDownloadAction(
            template='support/pdf/warranty_claim.html',
            context_fn=lambda obj: {'claim': obj, 'parts': obj.partsrequired_set.all()},
            filename_fn=lambda obj: f'warranty_{obj.owner_name}_{obj.machine_model}',
            zip_filename='warranty_claims.zip',
        )
    """

    short_description: str = 'Download as PDF'

    def __init__(self, template: str, context_fn: Callable[[Model], dict], filename_fn: Callable[[Model], str], zip_filename: str) -> None:
        """
        Initialise the action.

        :param template: Django template path used to render each PDF.
        :param context_fn: Callable that receives a model instance and returns
            the template context dict. ``generated_date`` is injected
            automatically and does not need to be included.
        :param filename_fn: Callable that receives a model instance and returns
            a raw filename string (without extension). Special characters
            are sanitised automatically.
        :param zip_filename: Filename for the ZIP archive when multiple records
            are selected (e.g. ``'warranty_claims.zip'``).
        """
        self.template: str = template
        self.context_fn: Callable[[Model], dict] = context_fn
        self.filename_fn: Callable[[Model], str] = filename_fn
        self.zip_filename: str = zip_filename
        base: str = re.sub(r'\W+', '_', zip_filename.replace('.zip', ''))
        self.__name__: str = f'download_{base}_pdf'

    def render_pdf(self, obj: Model) -> bytes:
        """
        Render a single model instance to PDF bytes.

        Builds the template context via ``context_fn``, injects
        ``generated_date``, renders the template string, then converts
        it to PDF using WeasyPrint.

        :param obj: The model instance to render.
        :returns: Raw PDF bytes.
        """
        context: dict = self.context_fn(obj)
        context['generated_date'] = date.today()
        html_string: str = render_to_string(self.template, context)
        css: CSS = CSS(filename=str(settings.BASE_DIR / 'theme' / 'static' / 'css' / 'pdf.css'))
        return HTML(string=html_string).write_pdf(stylesheets=[css])

    def build_filename(self, obj: Model) -> str:
        """
        Build a filesystem-safe PDF filename for a model instance.

        Passes the instance through ``filename_fn``, then replaces any
        non-alphanumeric characters with underscores and lowercases the result.

        :param obj: The model instance.
        :returns: A sanitised filename string ending in ``.pdf``.
        """
        raw: str = self.filename_fn(obj)
        slug: str = re.sub(r'\W+', '_',raw.strip()).strip('_').lower()
        return f'{slug}.pdf'

    def __call__(self, modeladmin: ModelAdmin, request: HttpRequest, queryset: QuerySet) -> HttpResponse:
        """
        Execute the admin action.

        Generates a single PDF response for one selected record, or a ZIP
        archive for multiple selected records.

        :param modeladmin: The originating ModelAdmin instance (unused).
        :param request: The current HTTP request (unused).
        :param queryset: The queryset of selected records.
        :returns: An ``HttpResponse`` with either ``application/pdf`` or
            ``application/zip`` content type and a ``Content-Disposition``
            attachment header.
        """
        with collect_image_failures() as failures:
            if queryset.count() == 1:
                obj: Model = queryset.first()
                pdf: bytes = self.render_pdf(obj)
                response: HttpResponse = HttpResponse(pdf, content_type='application/pdf')
                response['Content-Disposition'] = f'attachment; filename="{self.build_filename(obj)}"'
            else:
                buffer: io.BytesIO = io.BytesIO()
                with zipfile.ZipFile(buffer, 'w') as zf:
                    for obj in queryset:
                        zf.writestr(self.build_filename(obj), self.render_pdf(obj))
                response = HttpResponse(buffer.getvalue(), content_type='application/zip')
                response['Content-Disposition'] = f'attachment; filename="{self.zip_filename}"'
        if failures:
            logger.warning('PDF generated with %d missing image(s): %s', len(failures), failures)
            modeladmin.message_user(
                request,
                f'{len(failures)} image(s) could not be loaded and were left out of the PDF.',
                level=messages.WARNING,
            )
        return response
