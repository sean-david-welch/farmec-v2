from django.db import models
from base_model import BaseModel


class Supplier(BaseModel):
    id = models.TextField(primary_key=True)
    name = models.CharField(max_length=255, verbose_name='name')
    logo_image = models.URLField(blank=True, null=True, verbose_name='logo image')
    marketing_image = models.URLField(blank=True, null=True, verbose_name='marketing image')
    description = models.TextField(blank=True, null=True, verbose_name='description')
    social_facebook = models.URLField(blank=True, null=True, verbose_name='facebook')
    social_twitter = models.URLField(blank=True, null=True, verbose_name='twitter')
    social_instagram = models.URLField(blank=True, null=True, verbose_name='instagram')
    social_youtube = models.URLField(blank=True, null=True, verbose_name='youtube')
    social_linkedin = models.URLField(blank=True, null=True, verbose_name='linkedin')
    social_website = models.URLField(blank=True, null=True, verbose_name='website')
    slug = models.SlugField(max_length=255, blank=True, null=True, db_index=True, verbose_name='slug')

    class Meta:
        managed = True
        db_table = 'Supplier'

    def __str__(self):
        return self.name


class Machine(BaseModel):
    id = models.TextField(primary_key=True)
    supplier = models.TextField(blank=True, null=True, verbose_name='supplier id')
    name = models.CharField(max_length=255, verbose_name='name')
    machine_image = models.URLField(blank=True, null=True, verbose_name='image')
    description = models.TextField(blank=True, null=True, verbose_name='description')
    machine_link = models.URLField(blank=True, null=True, verbose_name='link')
    slug = models.SlugField(max_length=255, blank=True, null=True, db_index=True, verbose_name='slug')

    class Meta:
        managed = True
        db_table = 'Machine'

    def __str__(self):
        return self.name


class Product(BaseModel):
    id = models.TextField(primary_key=True)
    machine = models.TextField(blank=True, null=True, verbose_name='machine id')
    name = models.CharField(max_length=255, verbose_name='name')
    product_image = models.URLField(blank=True, null=True, verbose_name='image')
    description = models.TextField(blank=True, null=True, verbose_name='description')
    product_link = models.URLField(blank=True, null=True, verbose_name='link')
    slug = models.SlugField(max_length=255, blank=True, null=True, db_index=True, verbose_name='slug')

    class Meta:
        managed = True
        db_table = 'Product'

    def __str__(self):
        return self.name


class Spareparts(BaseModel):
    id = models.TextField(primary_key=True)
    supplier = models.TextField(blank=True, null=True, verbose_name='supplier id')
    name = models.CharField(max_length=255, verbose_name='name')
    parts_image = models.URLField(blank=True, null=True, verbose_name='image')
    spare_parts_link = models.URLField(blank=True, null=True, verbose_name='link')
    slug = models.SlugField(max_length=255, blank=True, null=True, db_index=True, verbose_name='slug')

    class Meta:
        managed = True
        db_table = 'SpareParts'

    def __str__(self):
        return self.name


class Lineitems(BaseModel):
    id = models.TextField(primary_key=True)
    name = models.CharField(max_length=255, verbose_name='name')
    price = models.DecimalField(max_digits=10, decimal_places=2, verbose_name='price')
    image = models.URLField(blank=True, null=True, verbose_name='image')

    class Meta:
        managed = True
        db_table = 'LineItems'

    def __str__(self):
        return self.name


class Video(BaseModel):
    id = models.TextField(primary_key=True)
    supplier = models.TextField(blank=True, null=True, verbose_name='supplier id')
    web_url = models.URLField(blank=True, null=True, verbose_name='url')
    title = models.CharField(max_length=255, blank=True, null=True, verbose_name='title')
    description = models.TextField(blank=True, null=True, verbose_name='description')
    video_id = models.CharField(max_length=100, blank=True, null=True, verbose_name='video id')
    thumbnail_url = models.URLField(blank=True, null=True, verbose_name='thumbnail')

    class Meta:
        managed = True
        db_table = 'Video'

    def __str__(self):
        return self.title or self.video_id or 'Untitled Video'
