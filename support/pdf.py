import io
import re
import zipfile
from datetime import date
from typing import Callable

from django.contrib.admin import ModelAdmin
from django.db.models import Model, QuerySet
from django.http import HttpResponse
from django.template.loader import render_to_string
from weasyprint import HTML


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

    short_description = "Download as PDF"

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
        self.template = template
        self.context_fn = context_fn
        self.filename_fn = filename_fn
        self.zip_filename = zip_filename
        self.__name__ = (
            f"download_{re.sub(r'[^\w]+', '_', zip_filename.replace('.zip', ''))}_pdf"
        )

    def render_pdf(self, obj: Model) -> bytes:
        """
        Render a single model instance to PDF bytes.

        Builds the template context via ``context_fn``, injects
        ``generated_date``, renders the template string, then converts
        it to PDF using WeasyPrint.

        :param obj: The model instance to render.
        :returns: Raw PDF bytes.
        """
        context = self.context_fn(obj)
        context["generated_date"] = date.today()
        html_string = render_to_string(self.template, context)
        return HTML(string=html_string).write_pdf()

    def build_filename(self, obj: Model) -> str:
        """
        Build a filesystem-safe PDF filename for a model instance.

        Passes the instance through ``filename_fn``, then replaces any
        non-alphanumeric characters with underscores and lowercases the result.

        :param obj: The model instance.
        :returns: A sanitised filename string ending in ``.pdf``.
        """
        raw = self.filename_fn(obj)
        slug = re.sub(r"[^\w]+", "_", raw.strip()).strip("_").lower()
        return f"{slug}.pdf"

    def __call__(
        self,
        modeladmin: ModelAdmin,
        request,
        queryset: QuerySet,
    ) -> HttpResponse:
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
        if queryset.count() == 1:
            obj = queryset.first()
            pdf = self.render_pdf(obj)
            response = HttpResponse(pdf, content_type="application/pdf")
            response["Content-Disposition"] = (
                f'attachment; filename="{self.build_filename(obj)}"'
            )
            return response

        buffer = io.BytesIO()
        with zipfile.ZipFile(buffer, "w") as zf:
            for obj in queryset:
                zf.writestr(self.build_filename(obj), self.render_pdf(obj))

        response = HttpResponse(buffer.getvalue(), content_type="application/zip")
        response["Content-Disposition"] = f'attachment; filename="{self.zip_filename}"'
        return response
