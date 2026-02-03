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


class CarouselForm(forms.ModelForm):
    """Form for creating and updating Carousel instances."""

    class Meta:
        model = Carousel
        fields: list[str] = [
            'name',
            'image',
        ]


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


class TimelineForm(forms.ModelForm):
    """Form for creating and updating Timeline instances."""

    class Meta:
        model = Timeline
        fields: list[str] = [
            'title',
            'date',
            'body',
        ]
