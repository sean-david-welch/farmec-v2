from typing import Any
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
        widgets: dict[str, Any] = {
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Employee full name',
            }),
            'email': forms.EmailInput(attrs={
                'class': 'form-control',
                'placeholder': 'work@example.com',
            }),
            'role': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Job title or position',
            }),
            'profile_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/profile.png',
            }),
        }
