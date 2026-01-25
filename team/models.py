from django.db import models


class Employee(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    email = models.TextField()
    role = models.TextField()
    profile_image = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'Employee'
