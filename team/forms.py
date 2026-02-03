from django import forms

from .models import Employee


class EmployeeForm(forms.ModelForm):
    """Form for creating and updating Employee instances."""

    class Meta:
        model = Employee
        fields: list[str] = [
            'name',
            'email',
            'role',
            'profile_image',
        ]
