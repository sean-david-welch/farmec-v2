from django.db import models
from base_model import BaseModel


class Privacy(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Privacy'


class Terms(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Terms'
