from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import PrivacyForm, TermsForm
from .models import Privacy, Terms

BASE_READONLY = ('uid', 'created', 'modified')


@admin.register(Privacy)
class PrivacyAdmin(ModelAdmin):
    form = PrivacyForm
    list_display = ('title', 'order', 'publish')
    list_editable = ('order',)
    readonly_fields = BASE_READONLY
    ordering = ('order',)


@admin.register(Terms)
class TermsAdmin(ModelAdmin):
    form = TermsForm
    list_display = ('title', 'order', 'publish')
    list_editable = ('order',)
    readonly_fields = BASE_READONLY
    ordering = ('order',)
