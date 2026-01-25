from django.db import models
from base_model import BaseModel


class Blog(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.CharField(max_length=500, verbose_name='title')
    date = models.DateField(blank=True, null=True, verbose_name='date')
    main_image = models.URLField(blank=True, null=True, verbose_name='main image')
    subheading = models.CharField(max_length=500, blank=True, null=True, verbose_name='subheading')
    body = models.TextField(blank=True, null=True, verbose_name='body')
    slug = models.SlugField(max_length=500, blank=True, null=True, db_index=True, verbose_name='slug')

    class Meta:
        managed = True
        db_table = 'Blog'

    def __str__(self):
        return self.title


class Carousel(BaseModel):
    id = models.TextField(primary_key=True)
    name = models.CharField(max_length=255, verbose_name='name')
    image = models.URLField(blank=True, null=True, verbose_name='image')

    class Meta:
        managed = True
        db_table = 'Carousel'

    def __str__(self):
        return self.name


class Exhibition(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.CharField(max_length=255, verbose_name='title')
    date = models.DateField(blank=True, null=True, verbose_name='date')
    location = models.CharField(max_length=255, blank=True, null=True, verbose_name='location')
    info = models.TextField(blank=True, null=True, verbose_name='information')

    class Meta:
        managed = True
        db_table = 'Exhibition'

    def __str__(self):
        return self.title


class Timeline(BaseModel):
    id = models.TextField(primary_key=True)
    title = models.CharField(max_length=255, verbose_name='title')
    date = models.DateField(blank=True, null=True, verbose_name='date')
    body = models.TextField(blank=True, null=True, verbose_name='body')

    class Meta:
        managed = True
        db_table = 'Timeline'

    def __str__(self):
        return self.title
