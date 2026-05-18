import io
import zipfile
from contextlib import contextmanager
from unittest.mock import MagicMock, patch

from django.contrib.auth.models import User
from django.test import TestCase
from django.urls import reverse
from model_bakery import baker

from support.models import Warrantyclaim, WarrantyImage, Partsrequired


@contextmanager
def mock_image_open(content: bytes = b'fake image data'):
    mock_file = MagicMock()
    mock_file.__enter__ = lambda s: s
    mock_file.__exit__ = MagicMock(return_value=False)
    mock_file.read.return_value = content
    with patch('django.db.models.fields.files.FieldFile.open', return_value=mock_file):
        yield


class WarrantyclaimAdminActionsTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_superuser(username='admin', password='adminpass')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_download_pdf_detail__returns_pdf(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        url: str = reverse('admin:support_warrantyclaim_download_pdf_detail', args=[claim.pk])
        response = self.client.get(url)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response['Content-Type'], 'application/pdf')
        self.assertTrue(response.content.startswith(b'%PDF'))

    def test_download_pdf_detail__with_parts_and_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(Partsrequired, warranty=claim, _quantity=2)
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=4)
        url: str = reverse('admin:support_warrantyclaim_download_pdf_detail', args=[claim.pk])
        response = self.client.get(url)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response['Content-Type'], 'application/pdf')

    def test_download_images_detail__returns_zip(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=2)
        url: str = reverse('admin:support_warrantyclaim_download_images_detail', args=[claim.pk])
        with mock_image_open():
            response = self.client.get(url)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response['Content-Type'], 'application/zip')

    def test_download_images_detail__zip_contains_images(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim, owner_name='Jane Doe', machine_model='SIP 350')
        baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=2)
        url: str = reverse('admin:support_warrantyclaim_download_images_detail', args=[claim.pk])
        with mock_image_open():
            response = self.client.get(url)
        zf = zipfile.ZipFile(io.BytesIO(response.content))
        names: list[str] = zf.namelist()
        self.assertEqual(len(names), 2)
        self.assertTrue(all('jane_doe_sip_350' in n for n in names))

    def test_download_images_detail__no_images_returns_empty_zip(self):
        claim: Warrantyclaim = baker.make(Warrantyclaim)
        url: str = reverse('admin:support_warrantyclaim_download_images_detail', args=[claim.pk])
        with mock_image_open():
            response = self.client.get(url)
        self.assertEqual(response.status_code, 200)
        zf = zipfile.ZipFile(io.BytesIO(response.content))
        self.assertEqual(len(zf.namelist()), 0)
