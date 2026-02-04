from django import forms

from .models import Blog, Carousel, Exhibition, Timeline


class ContactForm(forms.Form):
    """Form for contact messages from website visitors."""

    name = forms.CharField(
        max_length=100,
        required=True,
        widget=forms.TextInput(attrs={
            'class': 'form-input',
            'placeholder': 'Your name',
        }),
    )

    email = forms.EmailField(
        required=True,
        widget=forms.EmailInput(attrs={
            'class': 'form-input',
            'placeholder': 'your@email.com',
        }),
    )

    phone = forms.CharField(
        max_length=20,
        required=False,
        widget=forms.TextInput(attrs={
            'class': 'form-input',
            'placeholder': '+353 1 825 9289',
        }),
    )

    message = forms.CharField(
        required=True,
        widget=forms.Textarea(attrs={
            'class': 'form-input form-textarea',
            'placeholder': 'Your message here...',
            'rows': 5,
        }),
    )


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
