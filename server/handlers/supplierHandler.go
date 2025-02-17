package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type SupplierHandler struct {
	service        services.SupplierService
	authMiddleware *middleware.AuthMiddlewareImpl
	supplierCache  *middleware.SupplierCache
}

func NewSupplierHandler(service services.SupplierService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *SupplierHandler {
	return &SupplierHandler{service: service, authMiddleware: authMiddleware, supplierCache: supplierCache}
}

func (handler *SupplierHandler) SupplierView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	component := pages.Suppliers(isAdmin, false, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering suppliers"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

func (handler *SupplierHandler) GetSuppliers(context *gin.Context) {
	request := context.Request.Context()
	suppliers, err := handler.service.GetSuppliers(request)

	if err != nil {
		log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}
	context.JSON(http.StatusOK, suppliers)
}

func (handler *SupplierHandler) GetSupplierByID(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")
	supplier, err := handler.service.GetSupplierById(request, id)

	if err != nil {
		log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}

	context.JSON(http.StatusOK, supplier)
}

func (handler *SupplierHandler) CreateSupplier(context *gin.Context) {
	request := context.Request.Context()
	var supplier types.Supplier
	if err := context.ShouldBindJSON(&supplier); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	dbSupplier := lib.DeserializeSupplier(supplier)
	result, err := handler.service.CreateSupplier(request, dbSupplier)
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

func (handler *SupplierHandler) UpdateSupplier(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var supplier types.Supplier
	if err := context.ShouldBindJSON(&supplier); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	dbSupplier := lib.DeserializeSupplier(supplier)
	result, err := handler.service.UpdateSupplier(request, id, &dbSupplier)
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
	request := context.Request.Context()
	id := context.Param("id")

	err := handler.service.DeleteSupplier(request, id)

	if err != nil {
		log.Printf("Error deleting supplier: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting supplier", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully", "id": id})
}
