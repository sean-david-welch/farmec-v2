import uuid

from django.contrib import messages
from django.contrib.auth.mixins import LoginRequiredMixin
from django.http import HttpResponse
from django.shortcuts import redirect
from django.views.generic.edit import FormView

from farmec.utils import EmailClient
from support.forms import WarrantyclaimForm, MachineregistrationForm
from support.models import Partsrequired


class WarrantyclaimFormView(LoginRequiredMixin, FormView):
    template_name = 'support/warranty_form.html'
    form_class = WarrantyclaimForm

    def form_valid(self, form: WarrantyclaimForm) -> HttpResponse:
        claim = form.save(commit=False)
        claim.id = str(uuid.uuid4())
        claim.save()

        part_count: int = int(self.request.POST.get('part_count', 0))
        for i in range(part_count):
            part_number = self.request.POST.get(f'part_number_{i}', '').strip()
            quantity_needed = self.request.POST.get(f'quantity_needed_{i}', '').strip()
            invoice_number = self.request.POST.get(f'invoice_number_{i}', '').strip()
            description = self.request.POST.get(f'part_description_{i}', '').strip()
            if part_number or quantity_needed:
                Partsrequired.objects.create(
                    id=str(uuid.uuid4()),
                    warranty=claim,
                    part_number=part_number or None,
                    quantity_needed=int(quantity_needed) if quantity_needed else 1,
                    invoice_number=invoice_number or None,
                    description=description or None,
                )

        parts = Partsrequired.objects.filter(warranty=claim)
        EmailClient().send_warranty_notification(claim=claim, parts=parts)
        messages.success(self.request, 'Warranty claim submitted successfully.')
        return redirect('catalog:spareparts')


class MachineregistrationFormView(LoginRequiredMixin, FormView):
    template_name = 'support/machineregistration_form.html'
    form_class = MachineregistrationForm

    def form_valid(self, form: MachineregistrationForm) -> HttpResponse:
        registration = form.save(commit=False)
        registration.id = str(uuid.uuid4())
        registration.save()
        EmailClient().send_registration_notification(reg=registration)
        messages.success(self.request, 'Machine registration submitted successfully.')
        return redirect('catalog:spareparts')
