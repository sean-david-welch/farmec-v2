import uuid

import pytest
from model_bakery import baker

from support.models import Warrantyclaim, WarrantyImage, Partsrequired
from support.pdf import PDFDownloadAction


WARRANTY_PDF = PDFDownloadAction(
    template='support/pdf/warranty_claim.html',
    context_fn=lambda claim: {'claim': claim, 'parts': claim.partsrequired_set.all(), 'images': claim.images.all()},
    filename_fn=lambda claim: f'warranty_{claim.owner_name}_{claim.machine_model}',
    zip_filename='warranty_claims.zip',
)


@pytest.mark.django_db
def test_warranty_pdf_renders_without_images():
    claim: Warrantyclaim = baker.make(Warrantyclaim)
    result: bytes = WARRANTY_PDF.render_pdf(claim)
    assert isinstance(result, bytes)
    assert result.startswith(b'%PDF')


@pytest.mark.django_db
def test_warranty_pdf_renders_with_images():
    claim: Warrantyclaim = baker.make(Warrantyclaim)
    baker.make(WarrantyImage, warranty=claim, image='farmec_images/Warranty/test.jpg', _quantity=4)
    result: bytes = WARRANTY_PDF.render_pdf(claim)
    assert isinstance(result, bytes)
    assert result.startswith(b'%PDF')


@pytest.mark.django_db
def test_warranty_pdf_renders_with_parts():
    claim: Warrantyclaim = baker.make(Warrantyclaim)
    baker.make(Partsrequired, warranty=claim, _quantity=2)
    result: bytes = WARRANTY_PDF.render_pdf(claim)
    assert isinstance(result, bytes)
    assert result.startswith(b'%PDF')


@pytest.mark.django_db
def test_warranty_pdf_build_filename():
    claim: Warrantyclaim = baker.make(Warrantyclaim, owner_name='Jane Doe', machine_model='SIP 350')
    assert WARRANTY_PDF.build_filename(claim) == 'warranty_jane_doe_sip_350.pdf'
