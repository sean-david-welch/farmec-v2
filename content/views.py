import json
import urllib.parse
import urllib.request
from typing import Any

from django.http import HttpRequest, HttpResponse
from django.views.generic import ListView, DetailView

from farmec.mixin import HTMXViewMixin
from django.conf import settings
from farmec.utils import EmailClient
from .forms import ContactForm
from .models import (
    Blog,
    BlogQuerySet,
    Carousel,
    CarouselQuerySet,
    Exhibition,
    ExhibitionQuerySet,
)


class HomeView(HTMXViewMixin, ListView):
    model: type[Carousel] = Carousel
    template_name: str = 'pages/home.html'
    context_object_name: str = 'carousels'
    queryset: CarouselQuerySet = Carousel.objects.publish().order_by('-created')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['form'] = ContactForm()
        context['google_maps_api_key'] = settings.GOOGLE_MAPS_API_KEY
        context['cloudflare_site_key'] = settings.CLOUDFLARE_SITE_KEY
        return context

    def verify_turnstile(self, token: str) -> bool:
        """
        Verify a Cloudflare Turnstile token against the siteverify API.

        :param token: The ``cf-turnstile-response`` token from the POST request.
        :returns: ``True`` if the token is valid, ``False`` otherwise.
        """
        data: bytes = urllib.parse.urlencode({
            'secret': settings.CLOUDFLARE_SECRET_KEY,
            'response': token,
        }).encode()
        req = urllib.request.Request('https://challenges.cloudflare.com/turnstile/v0/siteverify', data=data)
        with urllib.request.urlopen(req) as response:
            result: dict = json.loads(response.read())
        return result.get('success', False)

    def handle_htmx(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        base: dict[str, Any] = {'form': ContactForm(), 'cloudflare_site_key': settings.CLOUDFLARE_SITE_KEY}
        token: str = request.POST.get('cf-turnstile-response', '')
        if not self.verify_turnstile(token):
            return self.render_htmx_response(
                'includes/contact.html#contact_form', include_base_context=False, extra_context={**base, 'turnstile_error': True},
            )
        form: ContactForm = ContactForm(request.POST)
        if not form.is_valid():
            return self.render_htmx_response(
                'includes/contact.html#contact_form', include_base_context=False, extra_context={**base, 'form': form},
            )
        EmailClient().send_contact_notification(
            name=form.cleaned_data['name'],
            email=form.cleaned_data['email'],
            message=form.cleaned_data['message'],
        )
        return self.render_htmx_response(
            'includes/contact.html#contact_form', include_base_context=False, extra_context=base, message="Message sent! We'll be in touch shortly.",
        )


class BlogListView(ListView):
    model: type[Blog] = Blog
    template_name: str = 'content/blog_list.html'
    context_object_name: str = 'blogs'
    queryset: BlogQuerySet = Blog.objects.publish().order_by('-created')


class BlogDetailView(DetailView):
    model: type[Blog] = Blog
    template_name: str = 'content/blog_detail.html'
    context_object_name: str = 'blog'
    pk_url_kwarg: str = 'id'
    queryset: BlogQuerySet = Blog.objects.publish()


class ExhibitionListView(ListView):
    model: type[Exhibition] = Exhibition
    template_name: str = 'content/exhibition_list.html'
    context_object_name: str = 'exhibitions'
    queryset: ExhibitionQuerySet = Exhibition.objects.publish().order_by('-created')
