from django.contrib import admin
from unfold.admin import ModelAdmin

from .forms import EmployeeForm
from .models import Employee

BASE_READONLY = ('uid', 'created', 'modified')


@admin.register(Employee)
class EmployeeAdmin(ModelAdmin):
    form = EmployeeForm
    list_display = ('name', 'role', 'email', 'publish')
    search_fields = ('name', 'role', 'email')
    readonly_fields = BASE_READONLY
    ordering = ('name',)
