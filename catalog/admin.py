from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import SupplierForm, MachineForm, ProductForm, SparepartsForm, VideoForm
from .models import Supplier, Machine, Product, Spareparts, Video

BASE_READONLY = ('uid', 'created', 'modified')


@admin.register(Supplier)
class SupplierAdmin(ModelAdmin):
    form = SupplierForm
    list_display = ('name', 'slug', 'publish')
    search_fields = ('name', 'description')
    readonly_fields = BASE_READONLY
    ordering = ('name',)


@admin.register(Machine)
class MachineAdmin(ModelAdmin):
    form = MachineForm
    list_display = ('name', 'supplier', 'slug', 'publish')
    search_fields = ('name', 'description')
    list_filter = ('supplier',)
    readonly_fields = BASE_READONLY
    ordering = ('supplier__name', 'name')


@admin.register(Product)
class ProductAdmin(ModelAdmin):
    form = ProductForm
    list_display = ('name', 'machine', 'slug', 'publish')
    search_fields = ('name', 'description')
    list_filter = ('machine__supplier', 'machine')
    readonly_fields = BASE_READONLY
    ordering = ('machine__name', 'name')


@admin.register(Spareparts)
class SparepartsAdmin(ModelAdmin):
    form = SparepartsForm
    list_display = ('name', 'supplier', 'slug', 'publish')
    search_fields = ('name',)
    list_filter = ('supplier',)
    readonly_fields = BASE_READONLY
    ordering = ('supplier__name', 'name')


@admin.register(Video)
class VideoAdmin(ModelAdmin):
    form = VideoForm
    list_display = ('title', 'supplier', 'video_id', 'publish')
    search_fields = ('title', 'description')
    list_filter = ('supplier',)
    readonly_fields = ('title', 'description', 'video_id', 'thumbnail_url')
    ordering = ('supplier__name', 'title')
