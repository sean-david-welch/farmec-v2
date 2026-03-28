from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import SupplierForm, MachineForm, ProductForm, SparepartsForm, VideoForm
from .models import Supplier, Machine, Product, Spareparts, Video


@admin.register(Supplier)
class SupplierAdmin(ModelAdmin):
    form = SupplierForm
    list_display = ("name", "id", "slug")
    search_fields = ("name", "description")


@admin.register(Machine)
class MachineAdmin(ModelAdmin):
    form = MachineForm
    list_display = ("name", "supplier", "slug")
    search_fields = ("name", "description")
    list_filter = ("supplier",)


@admin.register(Product)
class ProductAdmin(ModelAdmin):
    form = ProductForm
    list_display = ("name", "machine", "slug")
    search_fields = ("name", "description")


@admin.register(Spareparts)
class SparepartsAdmin(ModelAdmin):
    form = SparepartsForm
    list_display = ("name", "supplier", "slug")
    search_fields = ("name",)


@admin.register(Video)
class VideoAdmin(ModelAdmin):
    form = VideoForm
    list_display = ("title", "supplier", "video_id")
    search_fields = ("title", "description")
