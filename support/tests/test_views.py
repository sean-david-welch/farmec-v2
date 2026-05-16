from typing import Any

from django.core import mail
from django.core.files.uploadedfile import SimpleUploadedFile
from django.template.response import TemplateResponse
from django.test import TestCase
from django.urls import reverse_lazy
from django.contrib.auth.models import User

from support.models import Warrantyclaim, WarrantyImage, Partsrequired, Machineregistration


VALID_WARRANTY_DATA = {
    'dealer': 'Test Dealer',
    'dealer_contact': 'John',
    'owner_name': 'Jane Doe',
    'machine_model': 'SIP 350',
    'serial_number': 'SN123456',
    'install_date': '2025-01-01',
    'failure_date': '2025-06-01',
    'repair_date': '2025-06-05',
    'failure_details': 'Engine failed to start.',
    'repair_details': 'Replaced starter motor.',
    'labour_hours': '2.5',
    'completed_by': 'Tech Joe',
    'part_count': '0',
}

VALID_REGISTRATION_DATA = {
    'dealer_name': 'Test Dealer',
    'dealer_address': '1 Main St',
    'owner_name': 'Jane Doe',
    'owner_address': '2 Farm Rd',
    'machine_model': 'SIP 350',
    'serial_number': 'SN123456',
    'install_date': '2025-01-01',
    'invoice_number': 'INV-001',
    'date': '2025-01-01',
    'completed_by': 'Tech Joe',
    'complete_supply': True,
    'pdi_complete': True,
    'pto_correct': True,
    'machine_test_run': True,
    'safety_induction': True,
    'operator_handbook': True,
}


class WarrantyclaimFormViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.url = reverse_lazy('support:warranty_claim')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_warranty_claim__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_warranty_claim__anonymous_redirects(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 302)
        self.assertIn('/login/', response['Location'])

    def test_warranty_claim__valid_submission(self):
        response = self.client.post(self.url, data=VALID_WARRANTY_DATA, follow=True)
        with self.subTest('redirects to spareparts'):
            self.assertRedirects(response, reverse_lazy('catalog:spareparts'))
        with self.subTest('claim saved to db'):
            self.assertEqual(Warrantyclaim.objects.count(), 1)
            claim = Warrantyclaim.objects.get()
            self.assertEqual(claim.owner_name, 'Jane Doe')
            self.assertEqual(claim.serial_number, 'SN123456')
        with self.subTest('notification email sent'):
            self.assertEqual(len(mail.outbox), 1)
            self.assertIn('Jane Doe', mail.outbox[0].subject)

    def test_warranty_claim__invalid_missing_required(self):
        response = self.client.post(self.url, data={})
        with self.subTest('returns 200'):
            self.assertEqual(response.status_code, 200)
        with self.subTest('no claim saved'):
            self.assertEqual(Warrantyclaim.objects.count(), 0)
        with self.subTest('form errors present'):
            self.assertTrue(response.context['form'].errors)

    def test_warranty_claim__invalid_bad_dates(self):
        data = {**VALID_WARRANTY_DATA, 'install_date': 'not-a-date'}
        response = self.client.post(self.url, data=data)
        self.assertEqual(response.status_code, 200)
        self.assertIn('install_date', response.context['form'].errors)

    def test_warranty_claim__with_parts(self):
        data = {
            **VALID_WARRANTY_DATA,
            'part_count': '2',
            'part_number_0': 'PN-001',
            'quantity_needed_0': '3',
            'invoice_number_0': 'INV-001',
            'part_description_0': 'Starter motor',
            'part_number_1': 'PN-002',
            'quantity_needed_1': '1',
            'invoice_number_1': 'INV-002',
            'part_description_1': 'Drive belt',
        }
        self.client.post(self.url, data=data)
        claim = Warrantyclaim.objects.get()
        parts = Partsrequired.objects.filter(warranty=claim)
        with self.subTest('two parts created'):
            self.assertEqual(parts.count(), 2)
        with self.subTest('part numbers correct'):
            self.assertIn('PN-001', parts.values_list('part_number', flat=True))
            self.assertIn('PN-002', parts.values_list('part_number', flat=True))

    def _make_image(self, name: str) -> SimpleUploadedFile:
        return SimpleUploadedFile(name, b'\xff\xd8\xff' + b'\x00' * 10, content_type='image/jpeg')

    def test_warranty_claim__with_images(self):
        images: list[SimpleUploadedFile] = [self._make_image(f'repair{i}.jpg') for i in range(4)]
        data: dict[str, Any] = {**VALID_WARRANTY_DATA, 'warranty_images': images}
        self.client.post(self.url, data=data)
        claim: Warrantyclaim = Warrantyclaim.objects.get()
        with self.subTest('four images saved'):
            self.assertEqual(WarrantyImage.objects.filter(warranty=claim).count(), 4)
        with self.subTest('images linked to claim'):
            self.assertTrue(WarrantyImage.objects.filter(warranty=claim).exists())

    def test_warranty_claim__fewer_than_four_images_rejected(self):
        images: list[SimpleUploadedFile] = [self._make_image(f'repair{i}.jpg') for i in range(2)]
        data: dict[str, Any] = {**VALID_WARRANTY_DATA, 'warranty_images': images}
        response = self.client.post(self.url, data=data)
        with self.subTest('returns 200'):
            self.assertEqual(response.status_code, 200)
        with self.subTest('no claim saved'):
            self.assertEqual(Warrantyclaim.objects.count(), 0)
        with self.subTest('form error on warranty_images'):
            self.assertIn('warranty_images', response.context['form'].errors)

    def test_warranty_claim__no_images_is_valid(self):
        response: TemplateResponse = self.client.post(self.url, data=VALID_WARRANTY_DATA, follow=True)
        self.assertRedirects(response, reverse_lazy('catalog:spareparts'))
        self.assertEqual(WarrantyImage.objects.count(), 0)

    def test_warranty_claim__parts_with_empty_rows_skipped(self):
        data = {
            **VALID_WARRANTY_DATA,
            'part_count': '2',
            'part_number_0': '',
            'quantity_needed_0': '',
            'part_number_1': 'PN-001',
            'quantity_needed_1': '1',
        }
        self.client.post(self.url, data=data)
        claim = Warrantyclaim.objects.get()
        self.assertEqual(Partsrequired.objects.filter(warranty=claim).count(), 1)


class MachineregistrationFormViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.url = reverse_lazy('support:machine_registration')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_machine_registration__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_machine_registration__anonymous_redirects(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 302)
        self.assertIn('/login/', response['Location'])

    def test_machine_registration__valid_submission(self):
        response = self.client.post(self.url, data=VALID_REGISTRATION_DATA, follow=True)
        with self.subTest('redirects to spareparts'):
            self.assertRedirects(response, reverse_lazy('catalog:spareparts'))
        with self.subTest('registration saved to db'):
            self.assertEqual(Machineregistration.objects.count(), 1)
            reg = Machineregistration.objects.get()
            self.assertEqual(reg.owner_name, 'Jane Doe')
            self.assertEqual(reg.serial_number, 'SN123456')
        with self.subTest('notification email sent'):
            self.assertEqual(len(mail.outbox), 1)
            self.assertIn('Jane Doe', mail.outbox[0].subject)

    def test_machine_registration__invalid_missing_required(self):
        response = self.client.post(self.url, data={})
        with self.subTest('returns 200'):
            self.assertEqual(response.status_code, 200)
        with self.subTest('no registration saved'):
            self.assertEqual(Machineregistration.objects.count(), 0)
        with self.subTest('form errors present'):
            self.assertTrue(response.context['form'].errors)

    def test_machine_registration__invalid_bad_date(self):
        data = {**VALID_REGISTRATION_DATA, 'install_date': 'not-a-date'}
        response = self.client.post(self.url, data=data)
        self.assertEqual(response.status_code, 200)
        self.assertIn('install_date', response.context['form'].errors)

    def test_machine_registration__boolean_fields_default_false(self):
        data = {**VALID_REGISTRATION_DATA}
        for field in ['complete_supply', 'pdi_complete', 'pto_correct', 'machine_test_run', 'safety_induction', 'operator_handbook']:
            data.pop(field, None)
        self.client.post(self.url, data=data)
        reg = Machineregistration.objects.get()
        with self.subTest('complete_supply defaults false'):
            self.assertFalse(reg.complete_supply)
        with self.subTest('pdi_complete defaults false'):
            self.assertFalse(reg.pdi_complete)
