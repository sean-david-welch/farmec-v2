from datetime import datetime

from django.contrib.sitemaps import Sitemap
from django.db.models import QuerySet
from django.urls import reverse

from catalog.models import Supplier, Machine, Spareparts
from content.models import Blog


class SupplierSitemap(Sitemap):
    changefreq: str = 'monthly'
    priority: float = 0.9

    def items(self) -> QuerySet[Supplier]:
        return Supplier.objects.publish().filter(slug__isnull=False)

    def lastmod(self, obj: Supplier) -> datetime:
        return obj.modified

    def location(self, obj: Supplier) -> str:
        return f'/suppliers/{obj.slug}/'


class MachineSitemap(Sitemap):
    changefreq: str = 'monthly'
    priority: float = 0.8

    def items(self) -> QuerySet[Machine]:
        return Machine.objects.publish().filter(slug__isnull=False)

    def lastmod(self, obj: Machine) -> datetime:
        return obj.modified

    def location(self, obj: Machine) -> str:
        return f'/machines/{obj.slug}/'


class SparePartsSitemap(Sitemap):
    changefreq: str = 'monthly'
    priority: float = 0.8

    def items(self) -> QuerySet[Spareparts]:
        return Spareparts.objects.publish().filter(slug__isnull=False)

    def lastmod(self, obj: Spareparts) -> datetime:
        return obj.modified

    def location(self, obj: Spareparts) -> str:
        return f'/spareparts/{obj.slug}/'


class BlogSitemap(Sitemap):
    changefreq: str = 'weekly'
    priority: float = 0.7

    def items(self) -> QuerySet[Blog]:
        return Blog.objects.publish()

    def lastmod(self, obj: Blog) -> datetime:
        return obj.modified

    def location(self, obj: Blog) -> str:
        return f'/blog/{obj.id}/'


STATIC_URLS: dict[str, float] = {
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
    changefreq: str = 'monthly'

    def items(self) -> list[str]:
        return list(STATIC_URLS.keys())

    def priority(self, item: str) -> float:
        return STATIC_URLS[item]

    def location(self, item: str) -> str:
        return reverse(item)


sitemaps: dict[str, type[Sitemap]] = {
    'suppliers': SupplierSitemap,
    'machines': MachineSitemap,
    'spareparts': SparePartsSitemap,
    'blog': BlogSitemap,
    'static': StaticSitemap,
}
