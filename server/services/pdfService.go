package services

import (
	"bytes"
	"fmt"

	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/signintech/gopdf"
)

type PdfService interface {
	RenderRegistrationPdf(registration *types.MachineRegistration) ([]byte, error)

	RenderWarrantyClaimPdf(warranty *types.WarranrtyParts) ([]byte, error)
}

type PdfServiceImpl struct {
}

func NewPdfService() *PdfServiceImpl {
	return &PdfServiceImpl{}
}

func (service *PdfServiceImpl) RenderRegistrationPdf(registration *types.MachineRegistration) ([]byte, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	if err := pdf.SetFont("Arial", "", 12); err != nil {
		return nil, err
	}

	content := fmt.Sprintf("Dealer Name: %s\nOwner Name: %s\nMachine Model: %s", registration.DealerName, registration.OwnerName, registration.MachineModel)
	pdf.SetX(10)
	pdf.SetY(20)
	if err := pdf.Text(content); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err := pdf.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (service *PdfServiceImpl) RenderWarrantyClaimPdf(warranty *types.WarranrtyParts) ([]byte, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	if err := pdf.SetFont("Arial", "", 12); err != nil {
		return nil, err
	}

	content := fmt.Sprintf("Dealer Name: %s\nDealer Contact: %s\nMachine Model: %s", warranty.Warranty.Dealer, *warranty.Warranty.DealerContact, *warranty.Warranty.OwnerName)
	pdf.SetX(10)
	pdf.SetY(20)
	if err := pdf.Text(content); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err := pdf.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
