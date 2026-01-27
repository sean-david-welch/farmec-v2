from django.db import models
from django.utils.translation import gettext_lazy as _
from base_model import BaseModel, BaseQuerySet


class EmployeeQuerySet(BaseQuerySet):
    pass


class Employee(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    name = models.CharField(max_length=255, verbose_name=_('name'), help_text=_('Employee full name'))
    email = models.EmailField(verbose_name=_('email'), help_text=_('Work email address'))
    role = models.CharField(max_length=255, verbose_name=_('role'), help_text=_('Job title or position'))
    profile_image = models.URLField(blank=True, null=True, verbose_name=_('profile image'), help_text=_('URL to employee profile photo'))

    objects = EmployeeQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Employee'

    def __str__(self):
        return self.name
