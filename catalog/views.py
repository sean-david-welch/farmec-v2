from django.views.generic import ListView, DetailView

from .models import Supplier, SupplierQuerySet


class SupplierListView(ListView):
    model: type[Supplier] = Supplier
    template_name: str = 'catalog/supplier_list.html'
    context_object_name: str = 'suppliers'
    queryset: SupplierQuerySet = Supplier.objects.publish().order_by('-created')


class SupplierDetailView(DetailView):
    model: type[Supplier] = Supplier
    template_name: str = 'catalog/supplier_detail.html'
    context_object_name: str = 'supplier'
    slug_field: str = 'slug'
    slug_url_kwarg: str = 'slug'
    queryset: SupplierQuerySet = Supplier.objects.publish()
