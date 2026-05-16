from django.contrib import admin
from django.http import HttpRequest, HttpResponse
from unfold.admin import ModelAdmin, TabularInline
from unfold.decorators import action

from .forms import WarrantyclaimForm, PartsrequiredForm, MachineregistrationForm
from .models import Warrantyclaim, WarrantyImage, Partsrequired, Machineregistration
from .pdf import PDFDownloadAction

BASE_READONLY = ('created', 'modified')


class WarrantyImageInline(TabularInline):
    model = WarrantyImage
    extra = 0
    fields = ('image',)


class PartsrequiredInline(TabularInline):
    model = Partsrequired
    form = PartsrequiredForm
    extra = 0
    fields = ('part_number', 'quantity_needed', 'invoice_number', 'description')


@admin.register(Warrantyclaim)
class WarrantyclaimAdmin(ModelAdmin):
    form = WarrantyclaimForm
    inlines = [PartsrequiredInline, WarrantyImageInline]
    actions = ['download_pdf']
    actions_detail = ['download_pdf_detail']
    list_display = ('owner_name', 'dealer', 'machine_model', 'serial_number', 'failure_date', 'parts_count', 'completed_by')
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

    download_pdf: PDFDownloadAction = PDFDownloadAction(
        template='support/pdf/warranty_claim.html',
        context_fn=lambda claim: {'claim': claim, 'parts': claim.partsrequired_set.all(), 'images': claim.images.all()},
        filename_fn=lambda claim: f'warranty_{claim.owner_name}_{claim.machine_model}',
        zip_filename='warranty_claims.zip',
    )

    @admin.display(description='Parts')
    def parts_count(self, obj: Warrantyclaim) -> int:
        return obj.partsrequired_set.count()

    def get_queryset(self, request: HttpRequest):
        return super().get_queryset(request).prefetch_related('partsrequired_set', 'images')

    @action(description='Download PDF', icon='download')
    def download_pdf_detail(self, request: HttpRequest, object_id: str) -> HttpResponse:
        return self.download_pdf(self, request, Warrantyclaim.objects.filter(pk=object_id))


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
    actions = ['download_pdf']
    actions_detail = ['download_pdf_detail']
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

    download_pdf: PDFDownloadAction = PDFDownloadAction(
        template='support/pdf/machine_registration.html',
        context_fn=lambda reg: {'reg': reg},
        filename_fn=lambda reg: f'registration_{reg.owner_name}_{reg.machine_model}',
        zip_filename='machine_registrations.zip',
    )

    @action(description='Download PDF', icon='download')
    def download_pdf_detail(self, request: HttpRequest, object_id: str) -> HttpResponse:
        return self.download_pdf(self, request, Machineregistration.objects.filter(pk=object_id))
