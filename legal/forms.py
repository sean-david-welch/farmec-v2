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


class TermsForm(forms.ModelForm):
    """Form for creating and updating Terms instances."""

    class Meta:
        model = Terms
        fields: list[str] = [
            'title',
            'body',
        ]
