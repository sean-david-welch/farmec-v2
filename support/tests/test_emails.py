from django.core import mail
from django.test import TestCase
from model_bakery import baker

from farmec.utils import EmailClient
from support.models import Warrantyclaim, WarrantyImage, Partsrequired, Machineregistration


class WarrantyEmailTest(TestCase):
    def test_warranty_email__sends(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        EmailClient().send_warranty_notification(
            claim=claim,
            parts=Partsrequired.objects.none(),
            images=WarrantyImage.objects.none(),
        )
        self.assertEqual(len(mail.outbox), 1)

    def test_warranty_email__subject_contains_owner_and_model(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim, owner_name='Jane Doe', machine_model='SIP 350')
        EmailClient().send_warranty_notification(
            claim=claim,
            parts=Partsrequired.objects.none(),
            images=WarrantyImage.objects.none(),
        )
        self.assertIn('Jane Doe', mail.outbox[0].subject)
        self.assertIn('SIP 350', mail.outbox[0].subject)

    def test_warranty_email__renders_with_parts(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(Partsrequired, warranty=claim, _quantity=2)
        EmailClient().send_warranty_notification(
            claim=claim,
            parts=Partsrequired.objects.filter(warranty=claim),
            images=WarrantyImage.objects.none(),
        )
        self.assertEqual(len(mail.outbox), 1)

    def test_warranty_email__renders_with_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=4)
        EmailClient().send_warranty_notification(
            claim=claim,
            parts=Partsrequired.objects.none(),
            images=WarrantyImage.objects.filter(warranty=claim),
        )
        self.assertEqual(len(mail.outbox), 1)
        html: str = mail.outbox[0].alternatives[0][0]
        self.assertIn('farmec_images/Warranty/test.jpg', html)


class MachineregistrationEmailTest(TestCase):
    def test_registration_email__sends(self):
        reg: Machineregistration = baker.make(Machineregistration)
        EmailClient().send_registration_notification(reg=reg)
        self.assertEqual(len(mail.outbox), 1)

    def test_registration_email__subject_contains_owner_and_model(self):
        reg: Machineregistration = baker.make(Machineregistration, owner_name='Jane Doe', machine_model='SIP 350')
        EmailClient().send_registration_notification(reg=reg)
        self.assertIn('Jane Doe', mail.outbox[0].subject)
        self.assertIn('SIP 350', mail.outbox[0].subject)

    def test_registration_email__html_contains_owner_and_model(self):
        reg: Machineregistration = baker.make(Machineregistration, owner_name='Jane Doe', machine_model='SIP 350')
        EmailClient().send_registration_notification(reg=reg)
        html: str = mail.outbox[0].alternatives[0][0]
        self.assertIn('Jane Doe', html)
        self.assertIn('SIP 350', html)

    def test_registration_email__html_contains_serial_number(self):
        reg: Machineregistration = baker.make(Machineregistration, serial_number='SN123456')
        EmailClient().send_registration_notification(reg=reg)
        html: str = mail.outbox[0].alternatives[0][0]
        self.assertIn('SN123456', html)
