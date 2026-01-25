from django.db import models
from base_model import BaseModel


class Employee(BaseModel):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    email = models.TextField()
    role = models.TextField()
    profile_image = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Employee'
