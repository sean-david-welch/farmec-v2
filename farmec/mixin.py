from django.http import HttpRequest, HttpResponse
from django.shortcuts import render
from django_htmx.http import HttpResponseClientRedirect
from django_htmx.middleware import HtmxDetails


class HTMXViewMixin:
    """
    Mixin for views that need to handle HTMX requests separately from standard requests.

    Dispatches HTMX requests to `handle_htmx`, and provides `render_htmx_response`
    for rendering partial templates with the view's context.
    """

    request: HttpRequest

    @property
    def htmx(self) -> HtmxDetails:
        return self.request.htmx

    def htmx_redirect(self, url: str | None = None) -> HttpResponse:
        if url is None:
            url = self.request.path
        return HttpResponseClientRedirect(url)

    def handle_htmx(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        return HttpResponse(status=400)

    def render_htmx_response(self, template_name: str, extra_context: dict | None = None) -> HttpResponse:
        get_context_data = getattr(self, 'get_context_data', lambda **ctx: ctx)
        context = get_context_data()
        if extra_context:
            context |= extra_context
        return render(request=self.request, template_name=template_name, context=context)

    def dispatch(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        if self.htmx:
            return self.handle_htmx(request, *args, **kwargs)
        return super().dispatch(request, *args, **kwargs)
