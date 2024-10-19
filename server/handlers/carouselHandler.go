package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type CarouselHandler struct {
	service services.CarouselService
}

func NewCarouselHandler(service services.CarouselService) *CarouselHandler {
	return &CarouselHandler{service: service}
}

func (handler *CarouselHandler) GetCarousels(context *gin.Context) {
	ctx := context.Request.Context()
	carousels, err := handler.service.GetCarousels(ctx)
	if err != nil {
		log.Printf("error getting carousels: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting carousel images"})
		return
	}

	context.JSON(http.StatusOK, carousels)
}

func (handler *CarouselHandler) CreateCarousel(context *gin.Context) {
	ctx := context.Request.Context()
	var carousel types.Carousel
	dbCarousel := lib.DeserializeCarousel(carousel)

	if err := context.ShouldBindJSON(&carousel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := handler.service.CreateCarousel(ctx, &dbCarousel)
	if err != nil {
		log.Printf("Error creating carousel: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating carousel", "details": err.Error()})
		return
	}

	response := gin.H{
		"carousel":     carousel,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *CarouselHandler) UpdateCarousel(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	var carousel types.Carousel
	dbCarousel := lib.DeserializeCarousel(carousel)

	if err := context.ShouldBindJSON(&carousel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := handler.service.UpdateCarousel(ctx, id, &dbCarousel)
	if err != nil {
		log.Printf("Error updating carousel: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating carousel", "details": err.Error()})
		return
	}

	response := gin.H{
		"carousel":     carousel,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *CarouselHandler) DeleteCarousel(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteCarousel(ctx, id); err != nil {
		log.Printf("Error deleting carousel: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "carousel deleted successfully", "id": id})
}
