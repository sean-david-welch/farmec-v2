package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/store"
)

func InitTimelines(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	timelineRepository := store.NewTimelineRepository(database)
	service := services.NewTimelineService(timelineRepository)
	handler := handlers.NewTimelineHandler(service)

	TimelineRoutes(router, handler, adminMiddleware)
}

func TimelineRoutes(router *gin.Engine, handler *handlers.TimelineHandler, adminMiddleware *middleware.AdminMiddleware) {
	timelineGroup := router.Group("/api/timeline")

	timelineGroup.GET("", handler.GetTimelines)

	protected := timelineGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", handler.CreateTimeline)
		protected.PUT("/:id", handler.UpdateTimeline)
		protected.DELETE("/:id", handler.DeleteTimeline)
	}
}
