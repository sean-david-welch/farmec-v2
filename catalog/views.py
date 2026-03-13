from django.views.generic import ListView, DetailView

from .models import (
    Supplier,
    SupplierQuerySet,
    Machine,
    MachineQuerySet,
    Product,
    ProductQuerySet,
    Video, SparepartsQuerySet, Spareparts,
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


class SparePartsIndexView(ListView):
    model: type[Spareparts] = Spareparts
    template_name: str = 'support/spareparts.html'
    context_object_name: str = 'spareparts'
    queryset: SupplierQuerySet = Supplier.objects.publish().order_by('-created')


class SparePartsListView(ListView):
    model: type[Spareparts] = Spareparts
    template_name: str = 'support/spareparts_detail.html'
    context_object_name: str = 'spareparts'
    slug_field: str = 'supplier__slug'
    slug_url_kwarg: str = 'slug'
    queryset: SparepartsQuerySet = Spareparts.objects.publish()

    def get_queryset(self) -> SparepartsQuerySet:
        return super().get_queryset().filter(supplier__slug=self.kwargs.get(self.slug_url_kwarg))
