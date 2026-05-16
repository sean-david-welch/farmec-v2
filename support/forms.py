from typing import Any

from django.core.files.uploadedfile import UploadedFile
from django import forms
from .models import Warrantyclaim, Partsrequired, Machineregistration


class MultipleFileInput(forms.FileInput):
    allow_multiple_selected = True

    def value_from_datadict(self, data: dict[str, Any], files: Any, name: str) -> list[UploadedFile]:
        return files.getlist(name)


class MultipleFileField(forms.FileField):
    widget = MultipleFileInput

    def clean(self, data: list[UploadedFile], initial: Any = None) -> list[UploadedFile]:
        if not data:
            if self.required:
                raise forms.ValidationError(self.error_messages['required'])
            return []
        files: list[UploadedFile] = [super().clean(f, initial) for f in data]
        if len(files) < 4:
            raise forms.ValidationError('Please upload at least 4 images.')
        return files


class WarrantyclaimForm(forms.ModelForm):
    """Form for creating and updating Warrantyclaim instances."""
    install_date = forms.DateField(widget=forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}))
    failure_date = forms.DateField(widget=forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}))
    repair_date = forms.DateField(widget=forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}))
    failure_details = forms.CharField(widget=forms.Textarea(attrs={'rows': 3}))
    repair_details = forms.CharField(widget=forms.Textarea(attrs={'rows': 3}))
    dealer_contact = forms.EmailField(
        max_length=255, required=False, label='Dealer Contact (Email Address)',
    )
    labour_hours = forms.DecimalField(max_digits=8, decimal_places=2, widget=forms.NumberInput(attrs={'step': '0.5', 'min': '0'}))
    completed_by = forms.CharField(max_length=255)
    warranty_images = MultipleFileField(
        required=False, label='Warranty Images', help_text='Please upload at least four images of the the machine including all sides as well as the serial number of the machine',
    )

    class Meta:
        model = Warrantyclaim
        fields: list[str] = [
            'dealer',
            'dealer_contact',
            'owner_name',
            'machine_model',
            'serial_number',
            'install_date',
            'failure_date',
            'repair_date',
            'failure_details',
            'repair_details',
            'labour_hours',
            'completed_by',
        ]


class PartsrequiredForm(forms.ModelForm):
    """Form for creating and updating Partsrequired instances."""
    part_number = forms.CharField(max_length=100)
    invoice_number = forms.CharField(max_length=100)
    description = forms.CharField(widget=forms.Textarea())

    class Meta:
        model = Partsrequired
        fields: list[str] = [
            'warranty',
            'part_number',
            'quantity_needed',
            'invoice_number',
            'description',
        ]


class MachineregistrationForm(forms.ModelForm):
    """Form for creating and updating Machineregistration instances."""
    dealer_address = forms.CharField(max_length=500)
    owner_address = forms.CharField(max_length=500)
    install_date = forms.DateField(widget=forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}))
    invoice_number = forms.CharField(max_length=100)
    date = forms.DateField(widget=forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}))
    completed_by = forms.CharField(max_length=255)
    complete_supply = forms.BooleanField(required=False)
    pdi_complete = forms.BooleanField(required=False)
    pto_correct = forms.BooleanField(required=False)
    machine_test_run = forms.BooleanField(required=False)
    safety_induction = forms.BooleanField(required=False)
    operator_handbook = forms.BooleanField(required=False)

    class Meta:
        model = Machineregistration
        fields: list[str] = [
            'dealer_name',
            'dealer_address',
            'owner_name',
            'owner_address',
            'machine_model',
            'serial_number',
            'install_date',
            'invoice_number',
            'complete_supply',
            'pdi_complete',
            'pto_correct',
            'machine_test_run',
            'safety_induction',
            'operator_handbook',
            'date',
            'completed_by',
        ]
