from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import WarrantyclaimForm, PartsrequiredForm, MachineregistrationForm
from .models import Warrantyclaim, Partsrequired, Machineregistration


@admin.register(Warrantyclaim)
class WarrantyclaimAdmin(ModelAdmin):
    form = WarrantyclaimForm
    list_display = ("owner_name", "machine_model", "serial_number", "install_date")
    search_fields = ("owner_name", "machine_model", "serial_number", "dealer")
    list_filter = ("install_date", "dealer")


@admin.register(Partsrequired)
class PartsrequiredAdmin(ModelAdmin):
    form = PartsrequiredForm
    list_display = ("part_number", "quantity_needed", "warranty")
    search_fields = ("part_number", "invoice_number")


@admin.register(Machineregistration)
class MachineregistrationAdmin(ModelAdmin):
    form = MachineregistrationForm
    list_display = ("owner_name", "machine_model", "serial_number", "date")
    search_fields = ("owner_name", "machine_model", "serial_number", "dealer_name")
    list_filter = ("date", "dealer_name")
