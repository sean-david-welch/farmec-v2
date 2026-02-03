from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import EmployeeForm
from .models import Employee


@admin.register(Employee)
class EmployeeAdmin(ModelAdmin):
    form = EmployeeForm
    list_display = ("name", "role", "email")
    search_fields = ("name", "role", "email")
