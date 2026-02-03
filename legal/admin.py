from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import PrivacyForm, TermsForm
from .models import Privacy, Terms


@admin.register(Privacy)
class PrivacyAdmin(ModelAdmin):
    form = PrivacyForm
    list_display = ("title",)


@admin.register(Terms)
class TermsAdmin(ModelAdmin):
    form = TermsForm
    list_display = ("title",)
