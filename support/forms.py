from typing import Any
from django import forms

from .models import Warrantyclaim, Partsrequired, Machineregistration


class WarrantyclaimForm(forms.ModelForm):
    """Form for creating and updating Warrantyclaim instances."""

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
        widgets: dict[str, Any] = {
            'dealer': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Authorized dealer name',
            }),
            'dealer_contact': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Dealer contact person or phone',
            }),
            'owner_name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine owner/customer name',
            }),
            'machine_model': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine model name/number',
            }),
            'serial_number': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine serial number',
            }),
            'install_date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'failure_date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'repair_date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'failure_details': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 4,
                'placeholder': 'Description of the failure/damage',
            }),
            'repair_details': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 4,
                'placeholder': 'Description of work performed',
            }),
            'labour_hours': forms.NumberInput(attrs={
                'class': 'form-control',
                'placeholder': '0.00',
                'step': '0.01',
            }),
            'completed_by': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Technician or person who completed repair',
            }),
        }


class PartsrequiredForm(forms.ModelForm):
    """Form for creating and updating Partsrequired instances."""

    class Meta:
        model = Partsrequired
        fields: list[str] = [
            'warranty',
            'part_number',
            'quantity_needed',
            'invoice_number',
            'description',
        ]
        widgets: dict[str, Any] = {
            'warranty': forms.Select(attrs={
                'class': 'form-control',
            }),
            'part_number': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Supplier part number or SKU',
            }),
            'quantity_needed': forms.NumberInput(attrs={
                'class': 'form-control',
                'placeholder': '1',
                'min': '1',
            }),
            'invoice_number': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Supplier invoice reference',
            }),
            'description': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 3,
                'placeholder': 'Part name and specifications',
            }),
        }


class MachineregistrationForm(forms.ModelForm):
    """Form for creating and updating Machineregistration instances."""

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
        widgets: dict[str, Any] = {
            'dealer_name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Authorized dealer business name',
            }),
            'dealer_address': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Dealer location address',
            }),
            'owner_name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine owner/customer name',
            }),
            'owner_address': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Owner location address',
            }),
            'machine_model': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine model name/number',
            }),
            'serial_number': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine serial number',
            }),
            'install_date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'invoice_number': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Sales invoice reference',
            }),
            'complete_supply': forms.CheckboxInput(attrs={
                'class': 'form-check-input',
            }),
            'pdi_complete': forms.CheckboxInput(attrs={
                'class': 'form-check-input',
            }),
            'pto_correct': forms.CheckboxInput(attrs={
                'class': 'form-check-input',
            }),
            'machine_test_run': forms.CheckboxInput(attrs={
                'class': 'form-check-input',
            }),
            'safety_induction': forms.CheckboxInput(attrs={
                'class': 'form-check-input',
            }),
            'operator_handbook': forms.CheckboxInput(attrs={
                'class': 'form-check-input',
            }),
            'date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'completed_by': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Person who completed registration',
            }),
        }
