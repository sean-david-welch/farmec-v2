from typing import Any
from django import forms

from .models import Blog, Carousel, Exhibition, Timeline


class BlogForm(forms.ModelForm):
    """Form for creating and updating Blog instances."""

    class Meta:
        model = Blog
        fields: list[str] = [
            'title',
            'date',
            'main_image',
            'subheading',
            'body',
            'slug',
        ]
        widgets: dict[str, Any] = {
            'title': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Blog post headline',
            }),
            'date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'main_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/image.png',
            }),
            'subheading': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Optional subtitle or summary',
            }),
            'body': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 10,
                'placeholder': 'Blog post content',
            }),
            'slug': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'blog-post-title',
            }),
        }


class CarouselForm(forms.ModelForm):
    """Form for creating and updating Carousel instances."""

    class Meta:
        model = Carousel
        fields: list[str] = [
            'name',
            'image',
        ]
        widgets: dict[str, Any] = {
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Carousel slide name',
            }),
            'image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/slide.png',
            }),
        }


class ExhibitionForm(forms.ModelForm):
    """Form for creating and updating Exhibition instances."""

    class Meta:
        model = Exhibition
        fields: list[str] = [
            'title',
            'date',
            'location',
            'info',
        ]
        widgets: dict[str, Any] = {
            'title': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Exhibition or event name',
            }),
            'date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'location': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Venue or location name',
            }),
            'info': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 6,
                'placeholder': 'Event details and description',
            }),
        }


class TimelineForm(forms.ModelForm):
    """Form for creating and updating Timeline instances."""

    class Meta:
        model = Timeline
        fields: list[str] = [
            'title',
            'date',
            'body',
        ]
        widgets: dict[str, Any] = {
            'title': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Timeline event title',
            }),
            'date': forms.DateInput(attrs={
                'class': 'form-control',
                'type': 'date',
            }),
            'body': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 6,
                'placeholder': 'Event description and details',
            }),
        }
