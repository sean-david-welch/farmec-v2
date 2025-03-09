package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PdfHandler struct {
	service services.PdfService
}

func NewPdfHandler(service services.PdfService) *PdfHandler {
	return &PdfHandler{service: service}
}

func (handler *PdfHandler) RenderRegistrationPdf(context *gin.Context) {
	var registration types.MachineRegistration

	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error with request body: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	pdfBytes, err := handler.service.RenderRegistrationPdf(&registration)
	if err != nil {
		log.Printf("error when rendering pdf: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render PDF", "details": err.Error()})
		return
	}

	fileName := fmt.Sprintf("%s-%s.registration.pdf",
		strings.ReplaceAll(registration.DealerName, " ", ""),
		strings.ReplaceAll(registration.OwnerName, " ", ""))

	contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", fileName)

	context.Header("Content-Type", "application/pdf")
	context.Header("Content-Disposition", contentDisposition)
	context.Header("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))

	context.Data(http.StatusOK, "application/pdf", pdfBytes)
	context.JSON(http.StatusOK, gin.H{"filename": fileName})
}

func (handler *PdfHandler) RenderWarrantyClaimPdf(context *gin.Context) {
	var warranty types.WarrantyParts

	if err := context.ShouldBindJSON(&warranty); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	pdfBytes, err := handler.service.RenderWarrantyClaimPdf(&warranty)
	if err != nil {
		log.Printf("error when rendering pdf: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render PDF", "details": err.Error()})
		return
	}

	fileName := fmt.Sprintf("%s-%s.warranty.pdf",
		strings.ReplaceAll(warranty.Warranty.Dealer, " ", ""),
		func(ownerName *string) string {
			if ownerName != nil {
				return strings.ReplaceAll(*ownerName, " ", "")
			}
			return "unknown_owner"
		}(warranty.Warranty.OwnerName),
	)

	contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", fileName)

	context.Header("Content-Type", "application/pdf")
	context.Header("Content-Disposition", contentDisposition)
	context.Header("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))

	context.Data(http.StatusOK, "application/pdf", pdfBytes)
	context.JSON(http.StatusOK, gin.H{"filename": fileName})
}
