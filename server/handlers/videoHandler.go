package handlers

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type VideoHandler struct {
	service services.VideoService
}

func NewVideoHandler(service services.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}

func (handler *VideoHandler) GetVideos(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	videos, err := handler.service.GetVideos(ctx, id)
	if err != nil {
		log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}

	context.JSON(http.StatusOK, videos)
}

func (handler *VideoHandler) CreateVideo(context *gin.Context) {
	ctx := context.Request.Context()
	var video types.Video
	fmt.Printf("video: %v", video)

	if err := context.ShouldBindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "Details": err.Error()})
		return
	}

	dbVideo := lib.DeserializeVideo(video)
	fmt.Printf("videodb: %v", dbVideo)
	if err := handler.service.CreateVideo(ctx, &dbVideo); err != nil {
		log.Printf("Error creating video: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating video", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, video)
}

func (handler *VideoHandler) UpdateVideo(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	var video types.Video
	if err := context.ShouldBindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "Details": err.Error()})
		return
	}

	dbVideo := lib.DeserializeVideo(video)
	if err := handler.service.UpdateVideo(ctx, id, &dbVideo); err != nil {
		log.Printf("Error updating video: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating video", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, video)
}

func (handler *VideoHandler) DeleteVideo(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteVideo(ctx, id); err != nil {
		log.Printf("Error deleting video: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Erropr occurred while deleting supplier", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully", "id": id})
}
