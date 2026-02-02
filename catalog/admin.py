from django.contrib import admin
from unfold.admin import ModelAdmin
from catalog.models import Supplier, Machine, Product, Spareparts, Lineitems, Video

@admin.register(Supplier)
class SupplierAdmin(ModelAdmin):
    list_display = ("name", "id", "slug")
    search_fields = ("name", "description")


@admin.register(Machine)
class MachineAdmin(ModelAdmin):
    list_display = ("name", "supplier", "slug")
    search_fields = ("name", "description")
    list_filter = ("supplier",)


@admin.register(Product)
class ProductAdmin(ModelAdmin):
    list_display = ("name", "machine", "slug")
    search_fields = ("name", "description")


@admin.register(Spareparts)
class SparepartsAdmin(ModelAdmin):
    list_display = ("name", "supplier", "slug")
    search_fields = ("name",)


@admin.register(Lineitems)
class LineitemsAdmin(ModelAdmin):
    list_display = ("name", "price")
    search_fields = ("name",)


@admin.register(Video)
class VideoAdmin(ModelAdmin):
    list_display = ("title", "supplier", "video_id")
    search_fields = ("title", "description")
