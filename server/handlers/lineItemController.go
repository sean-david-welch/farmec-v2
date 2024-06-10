package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemController struct {
	service services.LineItemService
}

func NewLineItemController(service services.LineItemService) *LineItemController {
	return &LineItemController{service: service}
}

func (controller *LineItemController) GetLineItems(context *gin.Context) {
	lineItems, err := controller.service.GetLineItems()
	if err != nil {
		log.Printf("error occurred while getting lineItems: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting LineItem"})
		return
	}

	context.JSON(http.StatusOK, lineItems)
}

func (controller *LineItemController) CreateLineItem(context *gin.Context) {
	var lineItem types.LineItem

	if err := context.ShouldBindJSON(&lineItem); err != nil {
		log.Printf("error when creating lineItem: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating lineItem -  bad request"})
		return
	}

	result, err := controller.service.CreateLineItem(&lineItem)
	if err != nil {
		log.Printf("error when creating lineItem: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when creating lineItem"})
	}

	response := gin.H{
		"lineItem":     lineItem,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (controller *LineItemController) GetLineItemById(context *gin.Context) {
	id := context.Param("id")

	lineItem, err := controller.service.GetLineItemById(id)
	if err != nil {
		log.Printf("error occurred while getting lineItem: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting lineItem"})
		return
	}

	context.JSON(http.StatusOK, lineItem)
}

func (controller *LineItemController) UpdateLineItem(context *gin.Context) {
	id := context.Param("id")
	var lineItem types.LineItem

	if err := context.ShouldBindJSON(&lineItem); err != nil {
		log.Printf("error when updating lineItem: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating lineItem -  bad request"})
		return
	}

	result, err := controller.service.UpdateLineItem(id, &lineItem)
	if err != nil {
		log.Printf("error when updating lineItem: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when updating lineItem"})
	}

	response := gin.H{
		"lineItem":     lineItem,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (controller *LineItemController) DeleteLineItem(context *gin.Context) {
	id := context.Param("id")

	if err := controller.service.DeleteLineItem(id); err != nil {
		log.Printf("error occurred while deleting lineItem: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting lineItem"})
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "lineItem deleted successfully", "id": id})
}
