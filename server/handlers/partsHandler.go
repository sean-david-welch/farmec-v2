package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type PartsHandler struct {
	service        services.PartsService
	authMiddleware *middleware.AuthMiddlewareImpl
	supplierCache  *middleware.SupplierCache
}

func NewPartsHandler(service services.PartsService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *PartsHandler {
	return &PartsHandler{service: service, authMiddleware: authMiddleware, supplierCache: supplierCache}
}

func (handler *PartsHandler) PartsListView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.auth

}

func (handler *PartsHandler) PartsDetailView(context *gin.Context) {}

func (handler *PartsHandler) GetParts(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	parts, err := handler.service.GetParts(request, id)
	if err != nil {
		log.Printf("Error getting parts: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting parts"})
		return
	}

	context.JSON(http.StatusOK, parts)
}

func (handler *PartsHandler) CreateParts(context *gin.Context) {
	request := context.Request.Context()
	var part types.Sparepart

	if err := context.ShouldBindJSON(&part); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Ivalid request body", "details": err.Error()})
		return
	}
	if part.SupplierID == "" || part.SupplierID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	dbPart := lib.DeserializeSparePart(part)
	result, err := handler.service.CreatePart(request, &dbPart)
	if err != nil {
		log.Printf("Error creating part: %v", err)
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error occurred while creating part", "details": err.Error()})
		return
	}

	response := gin.H{
		"part":             part,
		"presignedUrl":     result.PresignedImageUrl,
		"imageUrl":         result.ImageUrl,
		"presignedLinkUrl": result.PresignedLinkUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *PartsHandler) UpdateParts(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var part types.Sparepart
	if err := context.ShouldBindJSON(&part); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "details": err.Error()})
		return
	}
	if part.SupplierID == "" || part.SupplierID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	dbPart := lib.DeserializeSparePart(part)
	result, err := handler.service.UpdatePart(request, id, &dbPart)
	if err != nil {
		log.Printf("Error updating part: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating machine", "details": err.Error()})
		return
	}

	response := gin.H{
		"part":             part,
		"presignedUrl":     result.PresignedImageUrl,
		"imageUrl":         result.ImageUrl,
		"presignedLinkUrl": result.PresignedLinkUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *PartsHandler) DeletePart(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeletePart(request, id); err != nil {
		log.Printf("Erroir deleting part: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurrrd while deleting part", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "part deleted successfully", "id": id})
}
