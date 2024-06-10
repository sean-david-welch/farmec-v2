package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PartsHandler struct {
	partsService services.PartsService
}

func NewPartsHandler(partsService services.PartsService) *PartsHandler {
	return &PartsHandler{partsService: partsService}
}

func (handler *PartsHandler) GetParts(context *gin.Context) {
	id := context.Param("id")

	parts, err := handler.partsService.GetParts(id)
	if err != nil {
		log.Printf("Error getting parts: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting parts"})
		return
	}

	context.JSON(http.StatusOK, parts)
}

func (handler *PartsHandler) CreateParts(context *gin.Context) {
	var part types.Sparepart

	if err := context.ShouldBindJSON(&part); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Ivalid request body", "details": err.Error()})
		return
	}

	if part.SupplierID == "" || part.SupplierID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	result, err := handler.partsService.CreatePart(&part)
	if err != nil {
		log.Printf("Error creating part: %v", err)
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error occurred while creating part", "details": err.Error()})
		return
	}

	response := gin.H{
		"part":         part,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *PartsHandler) UpdateParts(context *gin.Context) {
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

	result, err := handler.partsService.UpdatePart(id, &part)
	if err != nil {
		log.Printf("Error updating part: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating machine", "details": err.Error()})
		return
	}

	response := gin.H{
		"part":         part,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *PartsHandler) DeletePart(context *gin.Context) {
	id := context.Param("id")

	if err := handler.partsService.DeletePart(id); err != nil {
		log.Printf("Erroir deleting part: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurrrd while deleting part", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "part deleted successfully", "id": id})
}
