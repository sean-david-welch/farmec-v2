package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type WarrantyHandler struct {
	service services.WarrantyService
}

func NewWarrantyHandler(service services.WarrantyService) *WarrantyHandler {
	return &WarrantyHandler{service: service}
}

func (handler *WarrantyHandler) GetWarranties(context *gin.Context) {
	warranties, err := handler.service.GetWarranties()
	if err != nil {
		log.Printf("error occurred while getting warranties: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting warranties"})
		return
	}

	context.JSON(http.StatusOK, warranties)
}

func (handler *WarrantyHandler) GetWarrantyById(context *gin.Context) {
	id := context.Param("id")

	warranty, parts, err := handler.service.GetWarrantyById(id)
	if err != nil {
		log.Printf("error occurred while getting warrantiy and adjoining parts: %v", err)
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
	var warrantyParts types.WarranrtyParts

	if err := context.ShouldBindJSON(&warrantyParts); err != nil {
		log.Printf("error occurred - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred bad request"})
		return
	}

	if err := handler.service.CreateWarranty(warrantyParts.Warranty, warrantyParts.Parts); err != nil {
		log.Printf("error occurred while creating warranty: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating warranty claim"})
		return
	}

	context.JSON(http.StatusCreated, warrantyParts)
}

func (handler *WarrantyHandler) UpdateWarranty(context *gin.Context) {
	id := context.Param("id")

	var warranty *types.WarrantyClaim
	var parts []types.PartsRequired

	body := gin.H{
		"warranty": warranty,
		"parts":    parts,
	}

	if err := context.ShouldBindJSON(body); err != nil {
		log.Printf("error occurred - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred bad request"})
		return
	}

	if err := handler.service.UpdateWarranty(id, warranty, parts); err != nil {
		log.Printf("error occurred while updating warranty: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating warranty claim"})
		return
	}

	context.JSON(http.StatusAccepted, body)
}

func (handler *WarrantyHandler) DeleteWarranty(context *gin.Context) {
	id := context.Param("id")

	if err := handler.service.DeleteWarranty(id); err != nil {
		log.Printf("error occurred while deleting warranty: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while deleting warranty"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully deleted warranty claim", "id": id})
}
