import os

from django.views.generic import ListView, DetailView

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
        context['google_maps_api_key'] = 'REMOVED'
        return context


class BlogListView(ListView):
    model: type[Blog] = Blog
    template_name: str = 'content/blog_list.html'
    context_object_name: str = 'blogs'
    queryset: BlogQuerySet = Blog.objects.publish().order_by('-created')


class BlogDetailView(DetailView):
    model: type[Blog] = Blog
    template_name: str = 'content/blog_detail.html'
    context_object_name: str = 'blog'
    slug_field: str = 'slug'
    slug_url_kwarg: str = 'slug'
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
