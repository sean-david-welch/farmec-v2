from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import SupplierForm, MachineForm, ProductForm, SparepartsForm, VideoForm
from .models import Supplier, Machine, Product, Spareparts, Video

@admin.register(Supplier)
class SupplierAdmin(ModelAdmin):
    form = SupplierForm
    list_display = ('name', 'slug', 'order', 'publish')
    list_editable = ('order',)
    search_fields = ('name', 'description')
    date_hierarchy = 'created'
    ordering = ('order', 'name')
    fieldsets = (
        (None, {'fields': ('name', 'slug', 'order', 'publish', 'description', 'logo_image', 'marketing_image')}),
        ('Social & Links', {'fields': ('social_website', 'social_facebook', 'social_instagram', 'social_youtube', 'social_linkedin', 'social_twitter')}),
        ('SEO', {'fields': ('meta_title', 'meta_description'), 'classes': ('collapse',)}),
    )


@admin.register(Machine)
class MachineAdmin(ModelAdmin):
    form = MachineForm
    list_display = ('name', 'supplier', 'slug', 'order', 'publish')
    list_editable = ('order',)
    search_fields = ('name', 'description')
    list_filter = ('supplier',)
    date_hierarchy = 'created'
    ordering = ('order', 'supplier__name', 'name')
    fieldsets = (
        (None, {'fields': ('supplier', 'name', 'slug', 'order', 'publish', 'description', 'machine_image', 'machine_link')}),
        ('SEO', {'fields': ('meta_title', 'meta_description'), 'classes': ('collapse',)}),
    )


@admin.register(Product)
class ProductAdmin(ModelAdmin):
    form = ProductForm
    list_display = ('name', 'machine', 'slug', 'order', 'publish')
    list_editable = ('order',)
    search_fields = ('name', 'description')
    list_filter = ('machine__supplier', 'machine')
    date_hierarchy = 'created'
    ordering = ('order', 'machine__name', 'name')


@admin.register(Spareparts)
class SparepartsAdmin(ModelAdmin):
    form = SparepartsForm
    list_display = ('name', 'supplier', 'slug', 'order', 'publish')
    list_editable = ('order',)
    search_fields = ('name',)
    list_filter = ('supplier',)
    date_hierarchy = 'created'
    ordering = ('order', 'supplier__name', 'name')


@admin.register(Video)
class VideoAdmin(ModelAdmin):
    form = VideoForm
    list_display = ('title', 'supplier', 'video_id', 'order', 'publish')
    list_editable = ('order',)
    search_fields = ('title', 'description')
    list_filter = ('supplier',)
    date_hierarchy = 'created'
    readonly_fields = ('title', 'description', 'video_id', 'thumbnail_url')
    ordering = ('order', 'supplier__name', 'title')
