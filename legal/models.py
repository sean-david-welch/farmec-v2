from django.db import models


class Privacy(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Privacy'


class Terms(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Terms'
