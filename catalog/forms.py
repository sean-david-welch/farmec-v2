from django import forms

from .models import Supplier, Machine, Product, Spareparts, Video


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
        ]


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
        ]


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
        ]


class SparepartsForm(forms.ModelForm):
    """Form for creating and updating Spareparts instances."""

    class Meta:
        model = Spareparts
        fields: list[str] = [
            'supplier',
            'name',
            'parts_image',
            'spare_parts_link',
        ]



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
