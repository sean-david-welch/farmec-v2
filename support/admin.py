import io
import zipfile
from datetime import date

from django.contrib import admin
from django.http import HttpResponse
from django.template.loader import render_to_string
from unfold.admin import ModelAdmin, TabularInline
from weasyprint import HTML

from .forms import WarrantyclaimForm, PartsrequiredForm, MachineregistrationForm
from .models import Warrantyclaim, Partsrequired, Machineregistration

BASE_READONLY = ('created', 'modified')


def _render_pdf(template, context):
    html = render_to_string(template, context)
    return HTML(string=html).write_pdf()


def _slug(value):
    """Produce a filesystem-safe slug from a string."""
    import re
    return re.sub(r'[^\w]+', '_', value.strip()).strip('_').lower()


def download_warranty_pdf(modeladmin, request, queryset):
    generated = date.today()
    if queryset.count() == 1:
        claim = queryset.select_related().first()
        parts = claim.partsrequired_set.all()
        pdf = _render_pdf('support/pdf/warranty_claim.html', {'claim': claim, 'parts': parts, 'generated_date': generated})
        response = HttpResponse(pdf, content_type='application/pdf')
        filename = f"warranty_{_slug(claim.owner_name)}_{_slug(claim.machine_model)}.pdf"
        response['Content-Disposition'] = f'attachment; filename="{filename}"'
        return response

    buffer = io.BytesIO()
    with zipfile.ZipFile(buffer, 'w') as zf:
        for claim in queryset:
            parts = claim.partsrequired_set.all()
            pdf = _render_pdf('support/pdf/warranty_claim.html', {'claim': claim, 'parts': parts, 'generated_date': generated})
            filename = f"warranty_{_slug(claim.owner_name)}_{_slug(claim.machine_model)}.pdf"
            zf.writestr(filename, pdf)
    response = HttpResponse(buffer.getvalue(), content_type='application/zip')
    response['Content-Disposition'] = 'attachment; filename="warranty_claims.zip"'
    return response

download_warranty_pdf.short_description = 'Download as PDF'


def download_registration_pdf(modeladmin, request, queryset):
    generated = date.today()
    if queryset.count() == 1:
        reg = queryset.first()
        pdf = _render_pdf('support/pdf/machine_registration.html', {'reg': reg, 'generated_date': generated})
        response = HttpResponse(pdf, content_type='application/pdf')
        filename = f"registration_{_slug(reg.owner_name)}_{_slug(reg.machine_model)}.pdf"
        response['Content-Disposition'] = f'attachment; filename="{filename}"'
        return response

    buffer = io.BytesIO()
    with zipfile.ZipFile(buffer, 'w') as zf:
        for reg in queryset:
            pdf = _render_pdf('support/pdf/machine_registration.html', {'reg': reg, 'generated_date': generated})
            filename = f"registration_{_slug(reg.owner_name)}_{_slug(reg.machine_model)}.pdf"
            zf.writestr(filename, pdf)
    response = HttpResponse(buffer.getvalue(), content_type='application/zip')
    response['Content-Disposition'] = 'attachment; filename="machine_registrations.zip"'
    return response

download_registration_pdf.short_description = 'Download as PDF'


class PartsrequiredInline(TabularInline):
    model = Partsrequired
    form = PartsrequiredForm
    extra = 0
    fields = ('part_number', 'quantity_needed', 'invoice_number', 'description')


@admin.register(Warrantyclaim)
class WarrantyclaimAdmin(ModelAdmin):
    form = WarrantyclaimForm
    inlines = [PartsrequiredInline]
    actions = [download_warranty_pdf]
    list_display = ('owner_name', 'dealer', 'machine_model', 'serial_number', 'failure_date', 'parts_count', 'completed_by')

    @admin.display(description='Parts')
    def parts_count(self, obj):
        return obj.partsrequired_set.count()
    search_fields = ('owner_name', 'machine_model', 'serial_number', 'dealer')
    list_filter = ('dealer',)
    readonly_fields = BASE_READONLY
    ordering = ('-created',)

    def get_queryset(self, request):
        return super().get_queryset(request).prefetch_related('partsrequired_set')
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
    actions = [download_registration_pdf]
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
