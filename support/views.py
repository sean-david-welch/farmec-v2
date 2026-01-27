from django.views.generic import ListView, DetailView

from .models import (
    Warrantyclaim,
    WarrantyclaimQuerySet,
    Partsrequired,
    PartsrequiredQuerySet,
    Machineregistration,
    MachineregistrationQuerySet,
)


class WarrantyclaimListView(ListView):
    model: type[Warrantyclaim] = Warrantyclaim
    template_name: str = 'support/warrantyclaim_list.html'
    context_object_name: str = 'warranty_claims'
    queryset: WarrantyclaimQuerySet = Warrantyclaim.objects.publish().order_by('-created')


class WarrantyclaimDetailView(DetailView):
    model: type[Warrantyclaim] = Warrantyclaim
    template_name: str = 'support/warrantyclaim_detail.html'
    context_object_name: str = 'warranty_claim'
    queryset: WarrantyclaimQuerySet = Warrantyclaim.objects.publish().prefetch_related('partsrequired_set')


class PartsrequiredListView(ListView):
    model: type[Partsrequired] = Partsrequired
    template_name: str = 'support/partsrequired_list.html'
    context_object_name: str = 'parts_required'
    queryset: PartsrequiredQuerySet = Partsrequired.objects.publish().order_by('-created')


class MachineregistrationListView(ListView):
    model: type[Machineregistration] = Machineregistration
    template_name: str = 'support/machineregistration_list.html'
    context_object_name: str = 'machine_registrations'
    queryset: MachineregistrationQuerySet = Machineregistration.objects.publish().order_by('-created')


class MachineregistrationDetailView(DetailView):
    model: type[Machineregistration] = Machineregistration
    template_name: str = 'support/machineregistration_detail.html'
    context_object_name: str = 'machine_registration'
    queryset: MachineregistrationQuerySet = Machineregistration.objects.publish()
