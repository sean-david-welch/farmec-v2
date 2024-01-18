package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineController struct {
	service services.TimelineService
}

func NewTimelineController(service services.TimelineService) *TimelineController {
	return &TimelineController{service: service}
}

func(controller *TimelineController) GetTimelines(context *gin.Context) {
	timelines, err := controller.service.GetTimelines(); if err != nil {
		log.Printf("error getting timelines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting timeline"})
		return
	}

	context.JSON(http.StatusOK, timelines)
} 

func(controller *TimelineController) CreateTimeline(context *gin.Context) {
	var timeline types.Timeline

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := controller.service.CreateTimeline(&timeline); err != nil {
		log.Printf("error while creating timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, timeline)
}

func(controller *TimelineController) UpdateTimeline(context *gin.Context) {
	id := context.Param("id")
	var timeline types.Timeline

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := controller.service.UpdateTimeline(id, &timeline); err != nil {
		log.Printf("error while updating timeline: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, timeline)
}

func(controller *TimelineController) DeleteTimeline(context *gin.Context) {
	id := context.Param("id")

	if err := controller.service.DeleteTimeline(id); err != nil {
		log.Printf("Error deleting timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting timeline", "details": err.Error()})
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "timeline deleted successfully", "id": id})
}