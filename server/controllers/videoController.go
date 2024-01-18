package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type VideoController struct {
	videoService services.VideoService
}

func NewVideoController(videoService services.VideoService) *VideoController{
	return &VideoController{videoService: videoService}
}

func (controller *VideoController) GetVideos(context *gin.Context) {
	id := context.Param("id")

	videos, err := controller.videoService.GetVideos(id); if err != nil {
		log.Printf("Error getting suppliers: %v", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
        return
	}

	context.JSON(http.StatusOK, videos)
} 

func (controller *VideoController) CreateVideo(context *gin.Context) {
	var video types.Video

	if err := context.ShouldBindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "Details": err.Error()})
		return
	}

	if err := controller.videoService.CreateVideo(&video); err != nil {
		log.Printf("Error creating video: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating video", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, video)
}

func (controller *VideoController) UpdateVideo(context *gin.Context) {
	id := context.Param("id")

	var video types.Video

	if err := context.ShouldBindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "Details": err.Error()})
		return
	}

	if err := controller.videoService.UpdateVideo(id, &video); err != nil {
		log.Printf("Error updating video: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating video", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, video)
}

func (controller *VideoController) DeleteVideo(context *gin.Context) {
	id := context.Param("id")

	if err := controller.videoService.DeleteVideo(id); err != nil {
		log.Printf("Error deleting video: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Erropr occurred while deleting supplier", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully", "id": id})
}