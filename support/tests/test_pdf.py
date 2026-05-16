from django.test import TestCase
from model_bakery import baker

from support.models import Warrantyclaim, WarrantyImage, Partsrequired, Machineregistration
from support.pdf import PDFDownloadAction


WARRANTY_PDF = PDFDownloadAction(
    template='support/pdf/warranty_claim.html',
    context_fn=lambda claim: {'claim': claim, 'parts': claim.partsrequired_set.all(), 'images': claim.images.all()},
    filename_fn=lambda claim: f'warranty_{claim.owner_name}_{claim.machine_model}',
    zip_filename='warranty_claims.zip',
)


class WarrantyclaimPDFTest(TestCase):
    def test_warranty_pdf__renders_without_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        result: bytes = WARRANTY_PDF.render_pdf(claim)
        self.assertIsInstance(result, bytes)
        self.assertTrue(result.startswith(b'%PDF'))

    def test_warranty_pdf__renders_with_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=4)
        result: bytes = WARRANTY_PDF.render_pdf(claim)
        self.assertIsInstance(result, bytes)
        self.assertTrue(result.startswith(b'%PDF'))

    def test_warranty_pdf__renders_with_parts(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(Partsrequired, warranty=claim, _quantity=2)
        result: bytes = WARRANTY_PDF.render_pdf(claim)
        self.assertIsInstance(result, bytes)
        self.assertTrue(result.startswith(b'%PDF'))

    def test_warranty_pdf__build_filename(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim, owner_name='Jane Doe', machine_model='SIP 350')
        self.assertEqual(WARRANTY_PDF.build_filename(claim), 'warranty_jane_doe_sip_350.pdf')


REGISTRATION_PDF = PDFDownloadAction(
    template='support/pdf/machine_registration.html',
    context_fn=lambda reg: {'reg': reg},
    filename_fn=lambda reg: f'registration_{reg.owner_name}_{reg.machine_model}',
    zip_filename='machine_registrations.zip',
)


class MachineregistrationPDFTest(TestCase):
    def test_registration_pdf__renders(self):
        reg: Machineregistration = baker.make(Machineregistration)
        result: bytes = REGISTRATION_PDF.render_pdf(reg)
        self.assertIsInstance(result, bytes)
        self.assertTrue(result.startswith(b'%PDF'))

    def test_registration_pdf__build_filename(self):
        reg: Machineregistration = baker.make(Machineregistration, owner_name='Jane Doe', machine_model='SIP 350')
        self.assertEqual(REGISTRATION_PDF.build_filename(reg), 'registration_jane_doe_sip_350.pdf')
