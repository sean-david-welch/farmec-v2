from django.db import models
from django.utils.translation import gettext_lazy as _
from base_model import BaseModel, BaseQuerySet


class BlogQuerySet(BaseQuerySet):
    pass


class Blog(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    title = models.CharField(max_length=500, verbose_name=_('title'), help_text=_('Blog post headline'))
    date = models.DateField(blank=True, null=True, verbose_name=_('date'), help_text=_('Publication date'))
    main_image = models.ImageField(upload_to='farmec_images/Blogs/', blank=True, null=True, verbose_name=_('main image'), help_text=_('Featured/header image'))
    subheading = models.CharField(max_length=500, blank=True, null=True, verbose_name=_('subheading'), help_text=_('Optional subtitle or summary'))
    body = models.TextField(blank=True, null=True, verbose_name=_('body'), help_text=_('Blog post content'))
    slug = models.SlugField(max_length=500, blank=True, null=True, db_index=True, verbose_name=_('slug'), help_text=_('URL-friendly identifier'))

    objects = BlogQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Blog'

    def __str__(self):
        return self.title


class CarouselQuerySet(BaseQuerySet):
    pass


class Carousel(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    name = models.CharField(max_length=255, verbose_name=_('name'), help_text=_('Carousel slide name or identifier'))
    image = models.ImageField(upload_to='farmec_images/Carousels/', blank=True, null=True, verbose_name=_('image'), help_text=_('Carousel slide image'))

    objects = CarouselQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Carousel'

    def __str__(self):
        return self.name


class ExhibitionQuerySet(BaseQuerySet):
    pass


class Exhibition(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    title = models.CharField(max_length=255, verbose_name=_('title'), help_text=_('Exhibition or event name'))
    date = models.DateField(blank=True, null=True, verbose_name=_('date'), help_text=_('Event date or start date'))
    location = models.CharField(max_length=255, blank=True, null=True, verbose_name=_('location'), help_text=_('Venue or location name'))
    info = models.TextField(blank=True, null=True, verbose_name=_('information'), help_text=_('Event details and description'))

    objects = ExhibitionQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Exhibition'

    def __str__(self):
        return self.title


class TimelineQuerySet(BaseQuerySet):
    pass


class Timeline(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    title = models.CharField(max_length=255, verbose_name=_('title'), help_text=_('Timeline event title'))
    date = models.DateField(blank=True, null=True, verbose_name=_('date'), help_text=_('Event date'))
    body = models.TextField(blank=True, null=True, verbose_name=_('body'), help_text=_('Event description and details'))

    objects = TimelineQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Timeline'

    def __str__(self):
        return self.title
