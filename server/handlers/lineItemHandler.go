package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type LineItemHandler struct {
	service services.LineItemService
}

func NewLineItemHandler(service services.LineItemService) *LineItemHandler {
	return &LineItemHandler{service: service}
}

func (handler *LineItemHandler) GetLineItems(context *gin.Context) {
	ctx := context.Request.Context()
	lineItems, err := handler.service.GetLineItems(ctx)
	if err != nil {
		log.Printf("error occurred while getting lineItems: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting LineItem"})
		return
	}

	context.JSON(http.StatusOK, lineItems)
}

func (handler *LineItemHandler) CreateLineItem(context *gin.Context) {
	ctx := context.Request.Context()
	var lineItem db.LineItem

	if err := context.ShouldBindJSON(&lineItem); err != nil {
		log.Printf("error when creating lineItem: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating lineItem -  bad request"})
		return
	}

	result, err := handler.service.CreateLineItem(ctx, &lineItem)
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

func (handler *LineItemHandler) GetLineItemById(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	lineItem, err := handler.service.GetLineItemById(ctx, id)
	if err != nil {
		log.Printf("error occurred while getting lineItem: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting lineItem"})
		return
	}

	context.JSON(http.StatusOK, lineItem)
}

func (handler *LineItemHandler) UpdateLineItem(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")
	var lineItem db.LineItem

	if err := context.ShouldBindJSON(&lineItem); err != nil {
		log.Printf("error when updating lineItem: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating lineItem -  bad request"})
		return
	}

	result, err := handler.service.UpdateLineItem(ctx, id, &lineItem)
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

func (handler *LineItemHandler) DeleteLineItem(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteLineItem(ctx, id); err != nil {
		log.Printf("error occurred while deleting lineItem: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting lineItem"})
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "lineItem deleted successfully", "id": id})
}
