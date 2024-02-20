package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PdfController struct {
	service services.PdfService
}

func NewPdfController(service services.PdfService) *PdfController {
	return &PdfController{service: service}
}

func (controller *PdfController) RenderRegistrationPdf(context *gin.Context) {
	var registration types.MachineRegistration

	if err := context.ShouldBindJSON(&registration); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	pdfBytes, err := controller.service.RenderRegistrationPdf(&registration)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render PDF", "details": err.Error()})
		return
	}

	fileName := fmt.Sprintf("%s -- %s - registration.pdf", strings.ReplaceAll(registration.DealerName, " ", "_"), strings.ReplaceAll(registration.OwnerName, " ", "_"))

	context.Header("Content-Type", "application/pdf")
	context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	context.Header("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))

	context.Data(http.StatusOK, "application/pdf", pdfBytes)
}

func (controller *PdfController) RenderWarrantyClaimPdf(context *gin.Context) {
	var warranty types.WarranrtyParts

	if err := context.ShouldBindJSON(&warranty); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "pdf rendered successfully"})
}
