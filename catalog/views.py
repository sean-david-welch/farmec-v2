from django.views.generic import ListView, DetailView

from .models import (
    Supplier,
    SupplierQuerySet,
    Machine,
    MachineQuerySet,
    Product,
    ProductQuerySet,
    Video,
)


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

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['machines'] = Machine.objects.filter(supplier=self.object).publish()
        context['videos'] = Video.objects.filter(supplier=self.object).publish()
        return context


class MachineListView(ListView):
    model: type[Machine] = Machine
    template_name: str = 'catalog/machine_list.html'
    context_object_name: str = 'machines'
    queryset: MachineQuerySet = Machine.objects.publish().order_by('-created')


class MachineDetailView(DetailView):
    model: type[Machine] = Machine
    template_name: str = 'catalog/machine_detail.html'
    context_object_name: str = 'machine'
    slug_field: str = 'slug'
    slug_url_kwarg: str = 'slug'
    queryset: MachineQuerySet = Machine.objects.publish()

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['products'] = Product.objects.filter(machine=self.object).publish()
        return context


class ProductListView(ListView):
    model: type[Product] = Product
    template_name: str = 'catalog/product_list.html'
    context_object_name: str = 'products'
    queryset: ProductQuerySet = Product.objects.publish().order_by('-created')


class ProductDetailView(DetailView):
    model: type[Product] = Product
    template_name: str = 'catalog/product_detail.html'
    context_object_name: str = 'product'
    slug_field: str = 'slug'
    slug_url_kwarg: str = 'slug'
    queryset: ProductQuerySet = Product.objects.publish()
