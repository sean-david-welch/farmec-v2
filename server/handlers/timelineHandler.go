package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type TimelineHandler struct {
	service services.TimelineService
}

func NewTimelineHandler(service services.TimelineService) *TimelineHandler {
	return &TimelineHandler{service: service}
}

func (handler *TimelineHandler) GetTimelines(context *gin.Context) {
	request := context.Request.Context()
	timelines, err := handler.service.GetTimelines(request)
	if err != nil {
		log.Printf("error getting timelines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting timeline"})
		return
	}

	context.JSON(http.StatusOK, timelines)
}

func (handler *TimelineHandler) CreateTimeline(context *gin.Context) {
	request := context.Request.Context()
	var timeline types.Timeline

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbTimeline := lib.DeserializeTimeline(timeline)
	if err := handler.service.CreateTimeline(request, &dbTimeline); err != nil {
		log.Printf("error while creating timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, timeline)
}

func (handler *TimelineHandler) UpdateTimeline(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")
	var timeline types.Timeline

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbTimeline := lib.DeserializeTimeline(timeline)
	if err := handler.service.UpdateTimeline(request, id, &dbTimeline); err != nil {
		log.Printf("error while updating timeline: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, timeline)
}

func (handler *TimelineHandler) DeleteTimeline(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteTimeline(request, id); err != nil {
		log.Printf("Error deleting timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "timeline deleted successfully", "id": id})
}
