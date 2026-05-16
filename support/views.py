import uuid

from django.contrib import messages
from django.contrib.auth.mixins import LoginRequiredMixin
from django.db.models import QuerySet
from django.http import HttpResponse
from django.shortcuts import redirect
from django.views.generic.edit import FormView

from farmec.utils import EmailClient
from support.forms import WarrantyclaimForm, MachineregistrationForm
from support.models import Warrantyclaim, WarrantyImage, Machineregistration, Partsrequired


class WarrantyclaimFormView(LoginRequiredMixin, FormView):
    template_name: str = 'support/warranty_form.html'
    form_class: type[WarrantyclaimForm] = WarrantyclaimForm

    def form_invalid(self, form: WarrantyclaimForm) -> HttpResponse:
        field_labels: list[str] = [str(form.fields[f].label) or f.replace('_', ' ').title() for f in form.errors if f != '__all__']
        fields_str: str = ', '.join(field_labels)
        messages.error(self.request, f'Please correct the errors in: {fields_str}.')
        return super().form_invalid(form)

    def form_valid(self, form: WarrantyclaimForm) -> HttpResponse:
        claim: Warrantyclaim = form.save(commit=False)
        claim.id = str(uuid.uuid4())
        claim.save()
        part_count: int = int(self.request.POST.get('part_count', 0))
        for i in range(part_count):
            part_number: str = self.request.POST.get(f'part_number_{i}', '').strip()
            quantity_needed: str = self.request.POST.get(f'quantity_needed_{i}', '').strip()
            invoice_number: str = self.request.POST.get(f'invoice_number_{i}', '').strip()
            description: str = self.request.POST.get(f'part_description_{i}', '').strip()
            if part_number or quantity_needed:
                Partsrequired.objects.create(
                    id=str(uuid.uuid4()),
                    warranty=claim,
                    part_number=part_number or None,
                    quantity_needed=int(quantity_needed) if quantity_needed else 1,
                    invoice_number=invoice_number or None,
                    description=description or None,
                )
        for image_file in self.request.FILES.getlist('warranty_images'):
            WarrantyImage.objects.create(id=str(uuid.uuid4()), warranty=claim, image=image_file)
        parts: QuerySet[Partsrequired] = Partsrequired.objects.filter(warranty=claim)
        images: QuerySet[WarrantyImage] = WarrantyImage.objects.filter(warranty=claim)
        EmailClient().send_warranty_notification(claim=claim, parts=parts, images=images)
        messages.success(self.request, 'Warranty claim submitted successfully.')
        return redirect('catalog:spareparts')


class MachineregistrationFormView(LoginRequiredMixin, FormView):
    template_name: str = 'support/machineregistration_form.html'
    form_class: type[MachineregistrationForm] = MachineregistrationForm

    def form_invalid(self, form: MachineregistrationForm) -> HttpResponse:
        messages.error(self.request, 'Please correct the errors below.')
        return super().form_invalid(form)

    def form_valid(self, form: MachineregistrationForm) -> HttpResponse:
        registration: Machineregistration = form.save(commit=False)
        registration.id = str(uuid.uuid4())
        registration.save()
        EmailClient().send_registration_notification(reg=registration)
        messages.success(self.request, 'Machine registration submitted successfully.')
        return redirect('catalog:spareparts')
