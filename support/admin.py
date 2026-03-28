from django.contrib import admin
from unfold.admin import ModelAdmin, StackedInline

from .forms import WarrantyclaimForm, PartsrequiredForm, MachineregistrationForm
from .models import Warrantyclaim, Partsrequired, Machineregistration

BASE_READONLY = ('uid', 'created', 'modified')


class PartsrequiredInline(StackedInline):
    model = Partsrequired
    form = PartsrequiredForm
    extra = 0
    readonly_fields = BASE_READONLY


@admin.register(Warrantyclaim)
class WarrantyclaimAdmin(ModelAdmin):
    form = WarrantyclaimForm
    inlines = [PartsrequiredInline]
    list_display = ('owner_name', 'dealer', 'machine_model', 'serial_number', 'failure_date', 'completed_by')
    search_fields = ('owner_name', 'machine_model', 'serial_number', 'dealer')
    list_filter = ('dealer',)
    readonly_fields = BASE_READONLY
    ordering = ('-created',)


@admin.register(Partsrequired)
class PartsrequiredAdmin(ModelAdmin):
    form = PartsrequiredForm
    list_display = ('part_number', 'quantity_needed', 'invoice_number', 'warranty')
    search_fields = ('part_number', 'invoice_number')
    readonly_fields = BASE_READONLY
    ordering = ('warranty',)


@admin.register(Machineregistration)
class MachineregistrationAdmin(ModelAdmin):
    form = MachineregistrationForm
    list_display = ('owner_name', 'dealer_name', 'machine_model', 'serial_number', 'date', 'completed_by')
    search_fields = ('owner_name', 'machine_model', 'serial_number', 'dealer_name')
    list_filter = ('dealer_name',)
    readonly_fields = BASE_READONLY
    ordering = ('-created',)
