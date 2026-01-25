from django.db import models


class Supplier(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    logo_image = models.TextField(blank=True, null=True)
    marketing_image = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    social_facebook = models.TextField(blank=True, null=True)
    social_twitter = models.TextField(blank=True, null=True)
    social_instagram = models.TextField(blank=True, null=True)
    social_youtube = models.TextField(blank=True, null=True)
    social_linkedin = models.TextField(blank=True, null=True)
    social_website = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'Supplier'


class Machine(models.Model):
    id = models.TextField(primary_key=True)
    supplier = models.TextField(blank=True, null=True)
    name = models.TextField()
    machine_image = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    machine_link = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'Machine'


class Product(models.Model):
    id = models.TextField(primary_key=True)
    machine = models.TextField(blank=True, null=True)
    name = models.TextField()
    product_image = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    product_link = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'Product'


class Spareparts(models.Model):
    id = models.TextField(primary_key=True)
    supplier = models.TextField(blank=True, null=True)
    name = models.TextField()
    parts_image = models.TextField(blank=True, null=True)
    spare_parts_link = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'SpareParts'


class Lineitems(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    price = models.FloatField()
    image = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'LineItems'


class Video(models.Model):
    id = models.TextField(primary_key=True)
    supplier = models.TextField(blank=True, null=True)
    web_url = models.TextField(blank=True, null=True)
    title = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    video_id = models.TextField(blank=True, null=True)
    thumbnail_url = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = True
        db_table = 'Video'
