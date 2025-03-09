package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"github.com/sean-david-welch/farmec-v2/server/views/pages/details"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type WarrantyHandler struct {
	service        services.WarrantyService
	authMiddleware *middleware.AuthMiddlewareImpl
	supplierCache  *middleware.SupplierCache
}

func NewWarrantyHandler(service services.WarrantyService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *WarrantyHandler {
	return &WarrantyHandler{service: service, authMiddleware: authMiddleware, supplierCache: supplierCache}
}

func (handler *WarrantyHandler) WarrantyListView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	warranties, err := handler.service.GetWarranties(request)
	if err != nil {
		log.Printf("Error getting warranties: %v", err)
	}

	isError := err != nil
	component := pages.Warranties(isAdmin, isError, warranties, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering warranties: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering warranties"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

func (handler *WarrantyHandler) WarrantyDetailView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	id := context.Param("id")
	warranties, parts, err := handler.service.GetWarrantyById(request, id)
	if err != nil {
		log.Printf("Error getting warranties: %v", err)
	}

	isError := err != nil
	component := details.WarrantyDetail(isAdmin, isError, *warranties, parts, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering warranties: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering warranties"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

func (handler *WarrantyHandler) GetWarranties(context *gin.Context) {
	request := context.Request.Context()
	warranties, err := handler.service.GetWarranties(request)
	if err != nil {
		log.Printf("error occurred while getting warranties: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting warranties"})
		return
	}

	context.JSON(http.StatusOK, warranties)
}

func (handler *WarrantyHandler) GetWarrantyById(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	warranty, parts, err := handler.service.GetWarrantyById(request, id)
	if err != nil {
		log.Printf("error occurred while getting warranty and adjoining parts: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting warrantiy and adjoining parts"})
		return
	}

	response := gin.H{
		"warranty": warranty,
		"parts":    parts,
	}

	context.JSON(http.StatusOK, response)
}

func (handler *WarrantyHandler) CreateWarranty(context *gin.Context) {
	request := context.Request.Context()

	var warrantyParts types.WarrantyParts
	if err := context.ShouldBindJSON(&warrantyParts); err != nil {
		log.Printf("error occurred - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred bad request"})
		return
	}

	var parts []db.PartsRequired
	for _, part := range warrantyParts.Parts {
		parts = append(parts, lib.DeserializePartsRequired(part))
	}

	warranty := lib.DeserializeWarrantyClaim(*warrantyParts.Warranty)
	if err := handler.service.CreateWarranty(request, &warranty, parts); err != nil {
		log.Printf("error occurred while creating warranty: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating warranty claim"})
		return
	}

	context.JSON(http.StatusCreated, warrantyParts)
}

func (handler *WarrantyHandler) UpdateWarranty(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var warrantyParts types.WarrantyParts
	if err := context.ShouldBindJSON(&warrantyParts); err != nil {
		log.Printf("error occurred - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred bad request"})
		return
	}

	var parts []db.PartsRequired
	for _, part := range warrantyParts.Parts {
		parts = append(parts, lib.DeserializePartsRequired(part))
	}

	warranty := lib.DeserializeWarrantyClaim(*warrantyParts.Warranty)
	if err := handler.service.UpdateWarranty(request, id, &warranty, parts); err != nil {
		log.Printf("error occurred while updating warranty: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating warranty claim"})
		return
	}

	context.JSON(http.StatusAccepted, warrantyParts)
}

func (handler *WarrantyHandler) DeleteWarranty(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteWarranty(request, id); err != nil {
		log.Printf("error occurred while deleting warranty: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while deleting warranty"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully deleted warranty claim", "id": id})
}
