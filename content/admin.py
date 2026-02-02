from django.contrib import admin
from unfold.admin import ModelAdmin
from .models import Blog, Carousel, Exhibition, Timeline

@admin.register(Blog)
class BlogAdmin(ModelAdmin):
    list_display = ("title", "date", "slug")
    search_fields = ("title", "subheading", "body")
    list_filter = ("date",)


@admin.register(Carousel)
class CarouselAdmin(ModelAdmin):
    list_display = ("name", "id")
    search_fields = ("name",)


@admin.register(Exhibition)
class ExhibitionAdmin(ModelAdmin):
    list_display = ("title", "date", "location")
    search_fields = ("title", "location", "info")
    list_filter = ("date",)


@admin.register(Timeline)
class TimelineAdmin(ModelAdmin):
    list_display = ("title", "date")
    search_fields = ("title", "body")
    list_filter = ("date",)
