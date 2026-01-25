from django.db import models
from base_model import BaseModel


class Privacy(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.CharField(max_length=255, verbose_name='title')
    body = models.TextField(blank=True, null=True, verbose_name='body', help_text='Legal privacy policy text')

    class Meta:
        managed = True
        db_table = 'Privacy'

    def __str__(self):
        return self.title


class Terms(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.CharField(max_length=255, verbose_name='title')
    body = models.TextField(blank=True, null=True, verbose_name='body', help_text='Legal terms and conditions text')

    class Meta:
        managed = True
        db_table = 'Terms'

    def __str__(self):
        return self.title
