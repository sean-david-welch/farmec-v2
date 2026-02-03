from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import BlogForm, CarouselForm, ExhibitionForm, TimelineForm
from .models import Blog, Carousel, Exhibition, Timeline


@admin.register(Blog)
class BlogAdmin(ModelAdmin):
    form = BlogForm
    list_display = ("title", "date", "slug")
    search_fields = ("title", "subheading", "body")
    list_filter = ("date",)


@admin.register(Carousel)
class CarouselAdmin(ModelAdmin):
    form = CarouselForm
    list_display = ("name", "id")
    search_fields = ("name",)


@admin.register(Exhibition)
class ExhibitionAdmin(ModelAdmin):
    form = ExhibitionForm
    list_display = ("title", "date", "location")
    search_fields = ("title", "location", "info")
    list_filter = ("date",)


@admin.register(Timeline)
class TimelineAdmin(ModelAdmin):
    form = TimelineForm
    list_display = ("title", "date")
    search_fields = ("title", "body")
    list_filter = ("date",)
