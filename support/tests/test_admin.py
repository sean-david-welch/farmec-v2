import io
import zipfile

from django.contrib.auth.models import User
from django.test import TestCase
from model_bakery import baker

from support.models import Warrantyclaim, WarrantyImage, Partsrequired


class WarrantyclaimAdminActionsTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_superuser(username='admin', password='adminpass')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def action_url(self, pk: str, action: str) -> str:
        return f'/admin/support/warrantyclaim/{pk}/actions/{action}/'

    def test_download_pdf_detail__returns_pdf(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        response = self.client.get(self.action_url(claim.pk, 'download_pdf_detail'))
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response['Content-Type'], 'application/pdf')
        self.assertTrue(response.content.startswith(b'%PDF'))

    def test_download_pdf_detail__with_parts_and_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(Partsrequired, warranty=claim, _quantity=2)
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=4)
        response = self.client.get(self.action_url(claim.pk, 'download_pdf_detail'))
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response['Content-Type'], 'application/pdf')

    def test_download_images_detail__returns_zip(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=2)
        response = self.client.get(self.action_url(claim.pk, 'download_images_detail'))
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response['Content-Type'], 'application/zip')

    def test_download_images_detail__zip_contains_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim, owner_name='Jane Doe', machine_model='SIP 350')
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=2)
        response = self.client.get(self.action_url(claim.pk, 'download_images_detail'))
        zf = zipfile.ZipFile(io.BytesIO(response.content))
        names: list[str] = zf.namelist()
        self.assertEqual(len(names), 2)
        self.assertTrue(all('jane_doe_sip_350' in n for n in names))

    def test_download_images_detail__no_images_returns_empty_zip(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        response = self.client.get(self.action_url(claim.pk, 'download_images_detail'))
        self.assertEqual(response.status_code, 200)
        zf = zipfile.ZipFile(io.BytesIO(response.content))
        self.assertEqual(len(zf.namelist()), 0)
