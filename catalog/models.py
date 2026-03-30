from django.db import models
from django.utils.translation import gettext_lazy as _
from farmec.base_model import BaseModel, BaseQuerySet


class SupplierQuerySet(BaseQuerySet):
    pass


class Supplier(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    name = models.CharField(max_length=255, verbose_name=_('name'), help_text=_('Supplier business name'))
    logo_image = models.ImageField(upload_to='farmec_images/Suppliers/', blank=True, null=True, verbose_name=_('logo image'), help_text=_('Supplier logo'))
    marketing_image = models.ImageField(upload_to='farmec_images/Suppliers/', blank=True, null=True, verbose_name=_('marketing image'), help_text=_('Marketing/promotional image'))
    description = models.TextField(blank=True, null=True, verbose_name=_('description'), help_text=_('Supplier description and information'))
    social_facebook = models.URLField(blank=True, null=True, verbose_name=_('facebook'), help_text=_('Facebook page URL'))
    social_twitter = models.URLField(blank=True, null=True, verbose_name=_('twitter'), help_text=_('Twitter profile URL'))
    social_instagram = models.URLField(blank=True, null=True, verbose_name=_('instagram'), help_text=_('Instagram profile URL'))
    social_youtube = models.URLField(blank=True, null=True, verbose_name=_('youtube'), help_text=_('YouTube channel URL'))
    social_linkedin = models.URLField(blank=True, null=True, verbose_name=_('linkedin'), help_text=_('LinkedIn company URL'))
    social_website = models.URLField(blank=True, null=True, verbose_name=_('website'), help_text=_('Company website URL'))
    slug = models.SlugField(max_length=255, blank=True, null=True, verbose_name=_('slug'), help_text=_('URL-friendly identifier for this supplier'))

    objects = SupplierQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Supplier'

    def __str__(self):
        return self.name


class MachineQuerySet(BaseQuerySet):
    pass


class Machine(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, blank=True, null=True, verbose_name=_('supplier'), help_text=_('Supplier that manufactures this machine'))
    name = models.CharField(max_length=255, verbose_name=_('name'), help_text=_('Machine model name or title'))
    machine_image = models.ImageField(upload_to='farmec_images/Machines/', blank=True, null=True, verbose_name=_('image'), help_text=_('Machine product image'))
    description = models.TextField(blank=True, null=True, verbose_name=_('description'), help_text=_('Detailed machine specifications and features'))
    machine_link = models.URLField(blank=True, null=True, verbose_name=_('link'), help_text=_('URL to machine product page'))
    slug = models.SlugField(max_length=255, blank=True, null=True, verbose_name=_('slug'), help_text=_('URL-friendly identifier'))

    objects = MachineQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Machine'

    def __str__(self):
        return self.name


class ProductQuerySet(BaseQuerySet):
    pass


class Product(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    machine = models.ForeignKey('Machine', on_delete=models.CASCADE, blank=True, null=True, verbose_name=_('machine'), help_text=_('Machine this product is associated with'))
    name = models.CharField(max_length=255, verbose_name=_('name'), help_text=_('Product name or part name'))
    product_image = models.ImageField(upload_to='farmec_images/Products/', blank=True, null=True, verbose_name=_('image'), help_text=_('Product image'))
    description = models.TextField(blank=True, null=True, verbose_name=_('description'), help_text=_('Product description and details'))
    product_link = models.URLField(blank=True, null=True, verbose_name=_('link'), help_text=_('URL to product page or datasheet'))
    slug = models.SlugField(max_length=255, blank=True, null=True, verbose_name=_('slug'), help_text=_('URL-friendly identifier'))

    objects = ProductQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Product'

    def __str__(self):
        return self.name


class SparepartsQuerySet(BaseQuerySet):
    pass


class Spareparts(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, blank=True, null=True, verbose_name=_('supplier'), help_text=_('Supplier that provides this spare part'))
    name = models.CharField(max_length=255, verbose_name=_('name'), help_text=_('Spare part name or description'))
    parts_image = models.ImageField(upload_to='farmec_images/Spareparts/', blank=True, null=True, verbose_name=_('image'), help_text=_('Spare part image'))
    spare_parts_link = models.URLField(blank=True, null=True, verbose_name=_('link'), help_text=_('URL to spare part datasheet or ordering page'))
    slug = models.SlugField(max_length=255, blank=True, null=True, verbose_name=_('slug'), help_text=_('URL-friendly identifier'))

    objects = SparepartsQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'SpareParts'
        verbose_name = _('spare part')
        verbose_name_plural = _('spare parts')

    def __str__(self):
        return self.name



class VideoQuerySet(BaseQuerySet):
    pass


class Video(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, blank=True, null=True, verbose_name=_('supplier'), help_text=_('Supplier associated with this video'))
    web_url = models.URLField(blank=True, null=True, verbose_name=_('URL'), help_text=_('URL to video page or hosting platform'))
    title = models.CharField(max_length=255, blank=True, null=True, verbose_name=_('title'), help_text=_('Video title or name'))
    description = models.TextField(blank=True, null=True, verbose_name=_('description'), help_text=_('Video description and content info'))
    video_id = models.CharField(max_length=100, blank=True, null=True, verbose_name=_('video ID'), help_text=_('Video ID from hosting platform (e.g., YouTube ID)'))
    thumbnail_url = models.URLField(blank=True, null=True, verbose_name=_('thumbnail'), help_text=_('URL to video thumbnail image'))

    objects = VideoQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Video'

    def __str__(self):
        return self.title or self.video_id or 'Untitled Video'
