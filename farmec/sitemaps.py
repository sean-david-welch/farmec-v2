from django.contrib.sitemaps import Sitemap
from django.urls import reverse

from catalog.models import Supplier, Machine, Spareparts
from content.models import Blog, Exhibition


class SupplierSitemap(Sitemap):
    changefreq = 'monthly'
    priority = 0.9

    def items(self):
        return Supplier.objects.publish().filter(slug__isnull=False)

    def lastmod(self, obj):
        return obj.modified

    def location(self, obj):
        return f'/suppliers/{obj.slug}/'


class MachineSitemap(Sitemap):
    changefreq = 'monthly'
    priority = 0.8

    def items(self):
        return Machine.objects.publish().filter(slug__isnull=False)

    def lastmod(self, obj):
        return obj.modified

    def location(self, obj):
        return f'/machines/{obj.slug}/'


class SparePartsSitemap(Sitemap):
    changefreq = 'monthly'
    priority = 0.8

    def items(self):
        return Spareparts.objects.publish().filter(slug__isnull=False)

    def lastmod(self, obj):
        return obj.modified

    def location(self, obj):
        return f'/spareparts/{obj.slug}/'


class BlogSitemap(Sitemap):
    changefreq = 'weekly'
    priority = 0.7

    def items(self):
        return Blog.objects.publish()

    def lastmod(self, obj):
        return obj.modified

    def location(self, obj):
        return f'/blog/{obj.id}/'


class ExhibitionSitemap(Sitemap):
    changefreq = 'monthly'
    priority = 0.6

    def items(self):
        return Exhibition.objects.publish()

    def lastmod(self, obj):
        return obj.modified

    def location(self, obj):
        return f'/exhibitions/'


STATIC_URLS = {
    'content:home': 1.0,
    'catalog:supplier_list': 0.9,
    'catalog:machine_list': 0.8,
    'catalog:spareparts': 0.8,
    'content:blog_list': 0.7,
    'content:exhibition_list': 0.6,
    'team:employee_list': 0.6,
    'legal:privacy_list': 0.4,
}


class StaticSitemap(Sitemap):
    changefreq = 'monthly'

    def items(self):
        return list(STATIC_URLS.keys())

    def priority(self, item):
        return STATIC_URLS[item]

    def location(self, item):
        return reverse(item)


sitemaps = {
    'suppliers': SupplierSitemap,
    'machines': MachineSitemap,
    'spareparts': SparePartsSitemap,
    'blog': BlogSitemap,
    'exhibitions': ExhibitionSitemap,
    'static': StaticSitemap,
}
