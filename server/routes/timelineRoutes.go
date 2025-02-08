package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitTimelines(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AuthMiddleware) {
	repo := repository.NewTimelineRepo(database)
	service := services.NewTimelineService(repo)
	handler := handlers.NewTimelineHandler(service)

	TimelineRoutes(router, handler, adminMiddleware)
}

func TimelineRoutes(router *gin.Engine, handler *handlers.TimelineHandler, adminMiddleware *middleware.AuthMiddleware) {
	timelineGroup := router.Group("/api/timeline")

	timelineGroup.GET("", handler.GetTimelines)

	protected := timelineGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.POST("", handler.CreateTimeline)
		protected.PUT("/:id", handler.UpdateTimeline)
		protected.DELETE("/:id", handler.DeleteTimeline)
	}
}
