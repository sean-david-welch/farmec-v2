from django.contrib import admin
from unfold.admin import ModelAdmin, TabularInline

from .forms import WarrantyclaimForm, PartsrequiredForm, MachineregistrationForm
from .models import Warrantyclaim, Partsrequired, Machineregistration

BASE_READONLY = ('created', 'modified')


class PartsrequiredInline(TabularInline):
    model = Partsrequired
    form = PartsrequiredForm
    extra = 0
    fields = ('part_number', 'quantity_needed', 'invoice_number', 'description')


@admin.register(Warrantyclaim)
class WarrantyclaimAdmin(ModelAdmin):
    form = WarrantyclaimForm
    inlines = [PartsrequiredInline]
    list_display = ('owner_name', 'dealer', 'machine_model', 'serial_number', 'failure_date', 'completed_by')
    search_fields = ('owner_name', 'machine_model', 'serial_number', 'dealer')
    list_filter = ('dealer',)
    readonly_fields = BASE_READONLY
    ordering = ('-created',)
    fieldsets = (
        ('Dealer', {'fields': ('dealer', 'dealer_contact')}),
        ('Machine', {'fields': ('owner_name', 'machine_model', 'serial_number')}),
        ('Dates', {'fields': ('install_date', 'failure_date', 'repair_date')}),
        ('Details', {'fields': ('failure_details', 'repair_details', 'labour_hours', 'completed_by')}),
        ('Record', {'fields': ('created', 'modified'), 'classes': ('collapse',)}),
    )


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
    fieldsets = (
        ('Dealer', {'fields': ('dealer_name', 'dealer_address')}),
        ('Owner', {'fields': ('owner_name', 'owner_address')}),
        ('Machine', {'fields': ('machine_model', 'serial_number', 'install_date', 'invoice_number')}),
        ('Pre-Delivery Checklist', {'fields': (
            'complete_supply', 'pdi_complete', 'pto_correct',
            'machine_test_run', 'safety_induction', 'operator_handbook',
        )}),
        ('Completion', {'fields': ('date', 'completed_by')}),
        ('Record', {'fields': ('created', 'modified'), 'classes': ('collapse',)}),
    )
