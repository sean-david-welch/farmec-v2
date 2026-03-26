from django.http import JsonResponse
from django.views.generic import ListView, DetailView

from farmec.settings import env
from farmec.utils import EmailClient
from .forms import ContactForm
from .models import (
    Blog,
    BlogQuerySet,
    Carousel,
    CarouselQuerySet,
    Exhibition,
    ExhibitionQuerySet,
    Timeline,
    TimelineQuerySet,
)


class HomeView(ListView):
    model: type[Carousel] = Carousel
    template_name: str = 'pages/home.html'
    context_object_name: str = 'carousels'
    queryset: CarouselQuerySet = Carousel.objects.publish().order_by('-created')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['form'] = ContactForm()
        context['google_maps_api_key'] = env('GOOGLE_MAPS_API_KEY', default='')
        return context

    def post(self, request, *args, **kwargs):
        form = ContactForm(request.POST)
        if not form.is_valid():
            return JsonResponse({'errors': form.errors}, status=400)

        client = EmailClient()
        client.send_contact_notification(
            name=form.cleaned_data['name'],
            email=form.cleaned_data['email'],
            message=form.cleaned_data['message'],
        )
        return JsonResponse({'success': True})


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


class TimelineListView(ListView):
    model: type[Timeline] = Timeline
    template_name: str = 'content/timeline_list.html'
    context_object_name: str = 'timelines'
    queryset: TimelineQuerySet = Timeline.objects.publish().order_by('-created')
