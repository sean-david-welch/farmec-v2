from django.db import models
from django.utils.translation import gettext_lazy as _
from base_model import BaseModel, BaseQuerySet


class PrivacyQuerySet(BaseQuerySet):
    pass


class Privacy(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    title = models.CharField(max_length=255, verbose_name=_('title'), help_text=_('Page title'))
    body = models.TextField(blank=True, null=True, verbose_name=_('body'), help_text=_('Legal privacy policy text'))

    objects = PrivacyQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Privacy'
        verbose_name = _('privacy policy')
        verbose_name_plural = _('privacy policy')

    def __str__(self):
        return self.title


class TermsQuerySet(BaseQuerySet):
    pass


class Terms(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    title = models.CharField(max_length=255, verbose_name=_('title'), help_text=_('Page title'))
    body = models.TextField(blank=True, null=True, verbose_name=_('body'), help_text=_('Legal terms and conditions text'))

    objects = TermsQuerySet.as_manager()

    class Meta:
        managed = True
        db_table = 'Terms'
        verbose_name = _('terms & conditions')
        verbose_name_plural = _('terms & conditions')

    def __str__(self):
        return self.title
