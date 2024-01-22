package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type CarouselController struct {
	carouselService services.CarouselService
}

func NewCarouselController(carouselService services.CarouselService) *CarouselController {
	return &CarouselController{carouselService: carouselService}
}

func (controller *CarouselController) GetCarousels(context *gin.Context) {
	carousels, err := controller.carouselService.GetCarousels()
	if err != nil {
		log.Printf("error getting carousels: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting carousel images"})
		return
	}

	context.JSON(http.StatusOK, carousels)
}

func (controller *CarouselController) CreateCarousel(context *gin.Context) {
	var carousel types.Carousel

	if err := context.ShouldBindJSON(&carousel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := controller.carouselService.CreateCarousel(&carousel)
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

func (controller *CarouselController) UpdateCarousel(context *gin.Context) {
	id := context.Param("id")

	var carousel types.Carousel

	if err := context.ShouldBindJSON(&carousel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := controller.carouselService.UpdateCarousel(id, &carousel)
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

func (controller *CarouselController) DeleteCarousel(context *gin.Context) {
	id := context.Param("id")

	if err := controller.carouselService.DeleteCarousel(id); err != nil {
		log.Printf("Error deleting carousel: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "carousel deleted successfully", "id": id})
}
