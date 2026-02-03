from typing import Any
from django import forms
from django.db.models import Model

from .models import Supplier, Machine, Product, Spareparts, Lineitems, Video


class SupplierForm(forms.ModelForm):
    """Form for creating and updating Supplier instances."""

    class Meta:
        model = Supplier
        fields: list[str] = [
            'name',
            'logo_image',
            'marketing_image',
            'description',
            'social_facebook',
            'social_twitter',
            'social_instagram',
            'social_youtube',
            'social_linkedin',
            'social_website',
            'slug',
        ]
        widgets: dict[str, Any] = {
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Supplier business name',
            }),
            'logo_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/logo.png',
            }),
            'marketing_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/marketing.png',
            }),
            'description': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 5,
                'placeholder': 'Supplier description and information',
            }),
            'social_facebook': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://facebook.com/supplier',
            }),
            'social_twitter': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://twitter.com/supplier',
            }),
            'social_instagram': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://instagram.com/supplier',
            }),
            'social_youtube': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://youtube.com/supplier',
            }),
            'social_linkedin': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://linkedin.com/company/supplier',
            }),
            'social_website': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com',
            }),
            'slug': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'supplier-name',
            }),
        }


class MachineForm(forms.ModelForm):
    """Form for creating and updating Machine instances."""

    class Meta:
        model = Machine
        fields: list[str] = [
            'supplier',
            'name',
            'machine_image',
            'description',
            'machine_link',
            'slug',
        ]
        widgets: dict[str, Any] = {
            'supplier': forms.Select(attrs={
                'class': 'form-control',
            }),
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Machine model name',
            }),
            'machine_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/machine.png',
            }),
            'description': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 5,
                'placeholder': 'Detailed machine specifications and features',
            }),
            'machine_link': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/machine',
            }),
            'slug': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'machine-name',
            }),
        }


class ProductForm(forms.ModelForm):
    """Form for creating and updating Product instances."""

    class Meta:
        model = Product
        fields: list[str] = [
            'machine',
            'name',
            'product_image',
            'description',
            'product_link',
            'slug',
        ]
        widgets: dict[str, Any] = {
            'machine': forms.Select(attrs={
                'class': 'form-control',
            }),
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Product name',
            }),
            'product_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/product.png',
            }),
            'description': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 5,
                'placeholder': 'Product description and details',
            }),
            'product_link': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/product',
            }),
            'slug': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'product-name',
            }),
        }


class SparepartsForm(forms.ModelForm):
    """Form for creating and updating Spareparts instances."""

    class Meta:
        model = Spareparts
        fields: list[str] = [
            'supplier',
            'name',
            'parts_image',
            'spare_parts_link',
            'slug',
        ]
        widgets: dict[str, Any] = {
            'supplier': forms.Select(attrs={
                'class': 'form-control',
            }),
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Spare part name',
            }),
            'parts_image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/part.png',
            }),
            'spare_parts_link': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/spare-parts',
            }),
            'slug': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'spare-part-name',
            }),
        }


class LineitemsForm(forms.ModelForm):
    """Form for creating and updating Lineitems instances."""

    class Meta:
        model = Lineitems
        fields: list[str] = [
            'name',
            'price',
            'image',
        ]
        widgets: dict[str, Any] = {
            'name': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Product or item name',
            }),
            'price': forms.NumberInput(attrs={
                'class': 'form-control',
                'placeholder': '0.00',
                'step': '0.01',
            }),
            'image': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/item.png',
            }),
        }


class VideoForm(forms.ModelForm):
    """Form for creating and updating Video instances."""

    class Meta:
        model = Video
        fields: list[str] = [
            'supplier',
            'web_url',
            'title',
            'description',
            'video_id',
            'thumbnail_url',
        ]
        widgets: dict[str, Any] = {
            'supplier': forms.Select(attrs={
                'class': 'form-control',
            }),
            'web_url': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://youtube.com/watch?v=...',
            }),
            'title': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Video title',
            }),
            'description': forms.Textarea(attrs={
                'class': 'form-control',
                'rows': 4,
                'placeholder': 'Video description',
            }),
            'video_id': forms.TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'YouTube video ID',
            }),
            'thumbnail_url': forms.URLInput(attrs={
                'class': 'form-control',
                'placeholder': 'https://example.com/thumbnail.png',
            }),
        }
