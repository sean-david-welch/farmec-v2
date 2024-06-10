package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type SupplierHandler struct {
	supplierService services.SupplierService
}

func NewSupplierContoller(supplierService services.SupplierService) *SupplierHandler {
	return &SupplierHandler{supplierService: supplierService}
}

func (handler *SupplierHandler) GetSuppliers(context *gin.Context) {
	suppliers, err := handler.supplierService.GetSuppliers()

	if err != nil {
		log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}
	context.JSON(http.StatusOK, suppliers)
}

func (handler *SupplierHandler) CreateSupplier(context *gin.Context) {
	var supplier types.Supplier

	if err := context.ShouldBindJSON(&supplier); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	result, err := handler.supplierService.CreateSupplier(&supplier)
	if err != nil {
		log.Printf("Error creating supplier: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating supplier", "details": err.Error()})
		return
	}

	response := gin.H{
		"supplier":              supplier,
		"presignedLogoUrl":      result.PresignedLogoUrl,
		"logoUrl":               result.LogoUrl,
		"presignedMarketingUrl": result.PresignedMarketingUrl,
		"marketingUrl":          result.MarketingUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *SupplierHandler) GetSupplierByID(context *gin.Context) {
	id := context.Param("id")
	supplier, err := handler.supplierService.GetSupplierById(id)

	if err != nil {
		log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}

	context.JSON(http.StatusOK, supplier)
}

func (handler *SupplierHandler) UpdateSupplier(context *gin.Context) {
	id := context.Param("id")

	var supplier types.Supplier

	if err := context.ShouldBindJSON(&supplier); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	result, err := handler.supplierService.UpdateSupplier(id, &supplier)
	if err != nil {
		log.Printf("Error updating supplier: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating supplier", "details": err.Error()})
		return
	}

	response := gin.H{
		"supplier":              supplier,
		"presignedLogoUrl":      result.PresignedLogoUrl,
		"logoUrl":               result.LogoUrl,
		"presignedMarketingUrl": result.PresignedMarketingUrl,
		"marketingUrl":          result.MarketingUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *SupplierHandler) DeleteSupplier(context *gin.Context) {
	id := context.Param("id")

	err := handler.supplierService.DeleteSupplier(id)

	if err != nil {
		log.Printf("Error deleting supplier: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting supplier", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully", "id": id})
}
