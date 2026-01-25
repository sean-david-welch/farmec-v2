from django.db import models


class Blog(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    date = models.TextField(blank=True, null=True)
    main_image = models.TextField(blank=True, null=True)
    subheading = models.TextField(blank=True, null=True)
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Blog'


class Carousel(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    image = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Carousel'


class Exhibition(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    date = models.TextField(blank=True, null=True)
    location = models.TextField(blank=True, null=True)
    info = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Exhibition'


class Timeline(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    date = models.TextField(blank=True, null=True)
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Timeline'
