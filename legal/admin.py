from django.contrib import admin
from unfold.admin import ModelAdmin
from .models import Privacy, Terms

@admin.register(Privacy)
class PrivacyAdmin(ModelAdmin):
    list_display = ("title",)


@admin.register(Terms)
class TermsAdmin(ModelAdmin):
    list_display = ("title",)
