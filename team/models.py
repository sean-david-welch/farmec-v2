from django.db import models
from base_model import BaseModel


class Employee(BaseModel):
    id = models.TextField(primary_key=True)
    name = models.CharField(max_length=255, verbose_name='name')
    email = models.EmailField(verbose_name='email')
    role = models.CharField(max_length=255, verbose_name='role')
    profile_image = models.URLField(blank=True, null=True, verbose_name='profile image')

    class Meta:
        managed = True
        db_table = 'Employee'

    def __str__(self):
        return self.name
