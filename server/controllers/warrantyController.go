package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type WarrantyController struct {
	service *services.WarrantyService
}

func NewWarrantyController(service *services.WarrantyService) *WarrantyController {
	return &WarrantyController{service: service}
}

func(controller *WarrantyController) GetWarranties(context *gin.Context) {
	warranties, err := controller.service.GetWarranties(); if err != nil {
		log.Printf("error occurred while getting warranties: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting warranties"})
		return
	}

	context.JSON(http.StatusOK, warranties)
}

func(controller *WarrantyController) GetWarrantyById(context *gin.Context) {
	id := context.Param("id")

	warranty, parts, err := controller.service.GetWarrantyById(id); if err != nil {
		log.Printf("error occurred while getting warrantiy and adjoining parts: %v", err)	
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting warrantiy and adjoining parts"})
		return
	}

	response := gin.H{
		"warranty": warranty,
		"parts": parts,
	}

	context.JSON(http.StatusOK, response)
}

func(controller *WarrantyController) CreateWarranty(context *gin.Context) {
	var warranty *types.WarrantyClaim
	var parts []types.PartsRequired

	body := gin.H{
		"warranty": warranty,
		"parts": parts,
	}

	if err := context.ShouldBindJSON(body); err != nil {
		log.Printf("error occurred - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred bad request"})
		return
	}

	if err := controller.service.CreateWarranty(warranty, parts); err != nil {
		log.Printf("error occurred while creating warranty: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating warranty claim"})
		return
	}

	context.JSON(http.StatusCreated, body)	
}

func(controller *WarrantyController) UpdateWarranty(context *gin.Context) {
	id := context.Param("id")

	var warranty *types.WarrantyClaim
	var parts []types.PartsRequired

	body := gin.H{
		"warranty": warranty,
		"parts": parts,
	}

	if err := context.ShouldBindJSON(body); err != nil {
		log.Printf("error occurred - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred bad request"})
		return
	}

	if err := controller.service.UpdateWarranty(id, warranty, parts); err != nil {
		log.Printf("error occurred while updating warranty: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating warranty claim"})
		return
	}

	context.JSON(http.StatusAccepted, body)	
}

func(controller *WarrantyController) DeleteWarranty(context *gin.Context) {
	id := context.Param("id")

	if err := controller.service.DeleteWarranty(id); err != nil {
		log.Printf("error occurred while deleting warranty: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while deleting warranty"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully deleted warranty claim", "id": id})
}