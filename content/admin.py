from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import BlogForm, CarouselForm, ExhibitionForm, TimelineForm
from .models import Blog, Carousel, Exhibition, Timeline

BASE_READONLY = ('uid', 'created', 'modified')


@admin.register(Blog)
class BlogAdmin(ModelAdmin):
    form = BlogForm
    list_display = ('title', 'date', 'publish')
    search_fields = ('title', 'subheading', 'body')
    list_filter = ('date',)
    readonly_fields = BASE_READONLY
    ordering = ('-date',)


@admin.register(Carousel)
class CarouselAdmin(ModelAdmin):
    form = CarouselForm
    list_display = ('name', 'order', 'publish')
    search_fields = ('name',)
    readonly_fields = BASE_READONLY
    ordering = ('order',)


@admin.register(Exhibition)
class ExhibitionAdmin(ModelAdmin):
    form = ExhibitionForm
    list_display = ('title', 'date', 'location', 'publish')
    search_fields = ('title', 'location', 'info')
    list_filter = ('date',)
    readonly_fields = BASE_READONLY
    ordering = ('-date',)


@admin.register(Timeline)
class TimelineAdmin(ModelAdmin):
    form = TimelineForm
    list_display = ('title', 'date', 'publish')
    search_fields = ('title', 'body')
    readonly_fields = BASE_READONLY
    ordering = ('-date',)
