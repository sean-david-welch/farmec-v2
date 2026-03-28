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
        widgets: dict = {
            'install_date': forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}),
            'failure_date': forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}),
            'repair_date': forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}),
            'labour_hours': forms.NumberInput(attrs={'step': '0.5', 'min': '0'}),
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


class MachineregistrationForm(forms.ModelForm):
    """Form for creating and updating Machineregistration instances."""
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
        widgets: dict = {
            'install_date': forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}),
            'date': forms.DateInput(attrs={'type': 'date', 'onfocus': 'this.showPicker()'}),
        }
