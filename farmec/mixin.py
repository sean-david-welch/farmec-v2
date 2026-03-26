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
        """
        HTMX details for the current request.

        :returns: :class:`~django_htmx.middleware.HtmxDetails` for the current request.
        """
        return self.request.htmx

    def htmx_redirect(self, url: str | None = None) -> HttpResponse:
        """
        Return a response that instructs HTMX to perform a client-side redirect.

        :param url: URL to redirect to. Defaults to the current :attr:`request.path` if not provided.
        """
        if url is None:
            url = self.request.path
        return HttpResponseClientRedirect(url)

    def handle_htmx(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        """
        Handle an incoming HTMX request. Override in subclasses to provide custom behaviour.

        :param request: The current HTTP request.
        :returns: HTTP 400 by default; subclasses should return an appropriate :class:`~django.http.HttpResponse`.
        """
        if self.request:
            return HttpResponse(status=400)
        else:
            return HttpResponse(status=400)

    def render_htmx_response(self, template_name: str, extra_context: dict | None = None) -> HttpResponse:
        """
        Render a template with the view's context and return it as an HTMX partial response.

        :param template_name: Path to the template to render. Supports ``#partial-name`` suffixes for django-template-partials
        :param extra_context: Additional context variables merged on top of the view's context.
        """
        get_context_data = getattr(self, 'get_context_data', lambda **ctx: ctx)
        context = get_context_data()
        if extra_context:
            context |= extra_context
        return render(request=self.request, template_name=template_name, context=context)

    def dispatch(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        """
        Route HTMX requests to :meth:`handle_htmx` and standard requests to the normal dispatch chain.

        :param request: The current HTTP request.
        :returns: Response from :meth:`handle_htmx` for HTMX requests, otherwise the result of ``super().dispatch()``.
        """
        if self.htmx:
            return self.handle_htmx(request, *args, **kwargs)
        return super().dispatch(request, *args, **kwargs)
