import uuid

from django.http import HttpRequest, HttpResponse
from django.views.generic import ListView, DetailView

from farmec.mixin import HTMXViewMixin
from catalog.models import (
    Supplier,
    SupplierQuerySet,
    Machine,
    MachineQuerySet,
    Product,
    Video, SparepartsQuerySet, Spareparts,
)
from support.forms import WarrantyclaimForm
from support.models import Partsrequired


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


class SparePartsIndexView(HTMXViewMixin, ListView):
    model: type[Spareparts] = Spareparts
    template_name: str = 'support/spareparts.html'
    context_object_name: str = 'spareparts'
    queryset: SupplierQuerySet = Supplier.objects.publish().order_by('-created')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['form'] = WarrantyclaimForm()
        return context

    def handle_htmx(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        component = (request.GET if request.method == 'GET' else request.POST).get('component')
        if component == 'warranty':
            return self.handle_warranty(request)
        return super().handle_htmx(request, *args, **kwargs)

    def handle_warranty(self, request: HttpRequest) -> HttpResponse:
        if request.method == 'GET':
            return self.render_htmx_response(
                'support/warranty_form_dialog.html', include_base_context=False, extra_context={'form': WarrantyclaimForm()},
            )
        form = WarrantyclaimForm(request.POST)
        if not form.is_valid():
            return self.render_htmx_response(
                'support/warranty_form_dialog.html',
                include_base_context=False,
                extra_context={'form': form},
            )

        claim = form.save(commit=False)
        claim.id = str(uuid.uuid4())
        claim.save()

        part_count = int(request.POST.get('part_count', 0))
        for i in range(part_count):
            part_number = request.POST.get(f'part_number_{i}', '').strip()
            quantity_needed = request.POST.get(f'quantity_needed_{i}', '').strip()
            invoice_number = request.POST.get(f'invoice_number_{i}', '').strip()
            description = request.POST.get(f'part_description_{i}', '').strip()
            if part_number or quantity_needed:
                Partsrequired.objects.create(
                    id=str(uuid.uuid4()),
                    warranty=claim,
                    part_number=part_number or None,
                    quantity_needed=int(quantity_needed) if quantity_needed else 1,
                    invoice_number=invoice_number or None,
                    description=description or None,
                )

        return self.render_htmx_response(
            'support/warranty_form_dialog.html',
            include_base_context=False,
            extra_context={'form': WarrantyclaimForm()},
            message='Warranty claim submitted successfully.',
        )


class SparePartsListView(ListView):
    model: type[Spareparts] = Spareparts
    template_name: str = 'support/spareparts_detail.html'
    context_object_name: str = 'spareparts'
    slug_field: str = 'supplier__slug'
    slug_url_kwarg: str = 'slug'
    queryset: SparepartsQuerySet = Spareparts.objects.publish()

    def get_queryset(self) -> SparepartsQuerySet:
        return super().get_queryset().filter(supplier__slug=self.kwargs.get(self.slug_url_kwarg))
