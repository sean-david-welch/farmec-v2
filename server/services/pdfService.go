package services

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/signintech/gopdf"
)

type PdfService interface {
	RenderRegistrationPdf(registration *types.MachineRegistration) ([]byte, error)
	RenderWarrantyClaimPdf(warranty *types.WarranrtyParts) ([]byte, error)
	InitPdf(model string) (*gopdf.GoPdf, error)
	WrapText(text string) []string
	RenderStruct(pdf *gopdf.GoPdf, data interface{}, startY float64, startX float64) error
	RenderSlice(pdf *gopdf.GoPdf, data interface{}, startY float64, startX float64) error
}

type PdfServiceImpl struct {
}

func NewPdfService() *PdfServiceImpl {
	return &PdfServiceImpl{}
}

func (service *PdfServiceImpl) RenderRegistrationPdf(registration *types.MachineRegistration) ([]byte, error) {
	pdf, err := service.InitPdf("Registration Form " + "-- " + registration.OwnerName)

	if err != nil {
		return nil, err
	}

	registrationCopy := types.MachineRegistrationPDF{
		DealerName:       registration.DealerName,
		DealerAddress:    registration.DealerAddress,
		OwnerName:        registration.OwnerName,
		OwnerAddress:     registration.OwnerAddress,
		MachineModel:     registration.MachineModel,
		SerialNumber:     registration.SerialNumber,
		InstallDate:      registration.InstallDate,
		InvoiceNumber:    registration.InvoiceNumber,
		CompleteSupply:   registration.CompleteSupply,
		PdiComplete:      registration.PdiComplete,
		PtoCorrect:       registration.PtoCorrect,
		MachineTestRun:   registration.MachineTestRun,
		SafetyInduction:  registration.SafetyInduction,
		OperatorHandbook: registration.OperatorHandbook,
		Date:             registration.Date,
		CompletedBy:      registration.CompletedBy,
	}

	err = service.RenderStruct(pdf, &registrationCopy, 50, 50)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = pdf.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (service *PdfServiceImpl) RenderWarrantyClaimPdf(warranty *types.WarranrtyParts) ([]byte, error) {
	pdf, err := service.InitPdf("Warranty Claim " + "-- " + warranty.Warranty.OwnerName)
	if err != nil {
		return nil, err
	}

	if warranty.Warranty != nil {
		warrantyCopy := types.WarrantyClaimPDF{
			Dealer:         warranty.Warranty.Dealer,
			DealerContact:  warranty.Warranty.DealerContact,
			OwnerName:      warranty.Warranty.OwnerName,
			MachineModel:   warranty.Warranty.MachineModel,
			SerialNumber:   warranty.Warranty.SerialNumber,
			InstallDate:    warranty.Warranty.InstallDate,
			FailureDate:    warranty.Warranty.FailureDate,
			RepairDate:     warranty.Warranty.RepairDate,
			FailureDetails: warranty.Warranty.FailureDetails,
			RepairDetails:  warranty.Warranty.RepairDetails,
			LabourHours:    warranty.Warranty.LabourHours,
			CompletedBy:    warranty.Warranty.CompletedBy,
		}

		if err = service.RenderStruct(pdf, warrantyCopy, 50, 50); err != nil {
			return nil, err
		}
	}

	if len(warranty.Parts) > 0 {
		var partsCopy []types.PartsRequiredPDF
		for _, part := range warranty.Parts {
			partsCopy = append(partsCopy, types.PartsRequiredPDF{
				PartNumber:     part.PartNumber,
				QuantityNeeded: part.QuantityNeeded,
				InvoiceNumber:  part.InvoiceNumber,
				Description:    part.Description,
			})
		}

		if err = service.RenderSlice(pdf, partsCopy, 285, 50); err != nil {
			return nil, err
		}
	}

	var buf bytes.Buffer
	_, err = pdf.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (service *PdfServiceImpl) InitPdf(model string) (*gopdf.GoPdf, error) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	if err := pdf.AddTTFFont("OpenSans", "./fonts/opensans.ttf"); err != nil {
		log.Printf("error adding font %v", err)
		return nil, err
	}

	if err := pdf.SetFont("OpenSans", "", 16); err != nil {
		return nil, err
	}

	pdf.SetXY(50, 30)

	title := model
	pdf.AddHeader(func() {
		pdf.Text(title)
	})

	if err := pdf.Text(title); err != nil {
		return nil, err
	}

	pdf.SetY(pdf.GetY() + 110)
	if err := pdf.SetFont("OpenSans", "", 12); err != nil {
		log.Printf("error with font %v", err)
		return nil, err
	}

	return pdf, nil
}

func (service *PdfServiceImpl) WrapText(text string) []string {
	maxCharsPerLine := 100

	var wrapped []string
	for len(text) > 0 {
		if len(text) <= maxCharsPerLine {
			wrapped = append(wrapped, text)
			break
		}

		spaceIndex := strings.LastIndex(text[:maxCharsPerLine], " ")
		if spaceIndex == -1 {
			spaceIndex = maxCharsPerLine
		}
		wrapped = append(wrapped, text[:spaceIndex])
		text = text[spaceIndex+1:]
	}
	return wrapped
}

func (service *PdfServiceImpl) RenderStruct(pdf *gopdf.GoPdf, data interface{}, startY float64, startX float64) error {

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOfS.Field(i).Name

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				pdf.Text(fieldName + ": <nil>")
			} else {
				field = field.Elem()
			}
		}

		fieldValue := fmt.Sprintf("%v", field.Interface())

		text := fieldName + ": " + fieldValue
		wrappedText := service.WrapText(text)

		for _, line := range wrappedText {
			pdf.SetY(startY)
			pdf.SetX(startX)
			if err := pdf.Text(line); err != nil {
				log.Printf("error when rendering struct field to text: %s", err)
				return err
			}
			startY += 20
		}
	}

	return nil
}

func (service *PdfServiceImpl) RenderSlice(pdf *gopdf.GoPdf, data interface{}, startY float64, startX float64) error {
	startY += 40

	if err := pdf.SetFont("OpenSans", "", 10); err != nil {
		log.Printf("error with font %v", err)
		return err
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if elem.Kind() == reflect.Pointer {
				elem = elem.Elem()
			}

			if err := service.RenderStruct(pdf, elem.Interface(), startY, startX); err != nil {
				log.Printf("error when rendering slice element to text: %s", err)
				return err
			}

			startY += 85
		}
	}

	return nil
}
