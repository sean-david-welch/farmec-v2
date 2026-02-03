from typing import Any
from django import forms

from .models import Privacy, Terms


class PrivacyForm(forms.ModelForm):
    """Form for creating and updating Privacy instances."""

    class Meta:
        model = Privacy
        fields: list[str] = [
            'title',
            'body',
        ]
        widgets: dict[str, Any] = {
            'title': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Page title',
            }),
            'body': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 10,
                'placeholder': 'Legal privacy policy text',
            }),
        }


class TermsForm(forms.ModelForm):
    """Form for creating and updating Terms instances."""

    class Meta:
        model = Terms
        fields: list[str] = [
            'title',
            'body',
        ]
        widgets: dict[str, Any] = {
            'title': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Page title',
            }),
            'body': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 10,
                'placeholder': 'Legal terms and conditions text',
            }),
        }
