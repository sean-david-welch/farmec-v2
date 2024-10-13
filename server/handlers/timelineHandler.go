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
	ctx := context.Request.Context()
	timelines, err := handler.service.GetTimelines(ctx)
	if err != nil {
		log.Printf("error getting timelines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting timeline"})
		return
	}

	context.JSON(http.StatusOK, timelines)
}

func (handler *TimelineHandler) CreateTimeline(context *gin.Context) {
	ctx := context.Request.Context()
	var timeline types.Timeline
	dbTimeline := lib.DeserializeTimeline(timeline)

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := handler.service.CreateTimeline(ctx, &dbTimeline); err != nil {
		log.Printf("error while creating timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, timeline)
}

func (handler *TimelineHandler) UpdateTimeline(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")
	var timeline types.Timeline
	dbTimeline := lib.DeserializeTimeline(timeline)

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := handler.service.UpdateTimeline(ctx, id, &dbTimeline); err != nil {
		log.Printf("error while updating timeline: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, timeline)
}

func (handler *TimelineHandler) DeleteTimeline(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteTimeline(ctx, id); err != nil {
		log.Printf("Error deleting timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "timeline deleted successfully", "id": id})
}
