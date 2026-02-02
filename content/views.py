from django.views.generic import ListView, DetailView

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


class HomeView(ListView):
    model: type[Carousel] = Carousel
    template_name: str = 'pages/home.html'
    context_object_name: str = 'carousels'
    queryset: CarouselQuerySet = Carousel.objects.publish().order_by('-created')


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
