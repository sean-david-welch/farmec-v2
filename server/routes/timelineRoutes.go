package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitTimelines(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewTimelineRepository(database)
	service := services.NewTimelineService(repository)
	controller := controllers.NewTimelineController(service)

	TimelineRoutes(router, controller, adminMiddleware)
}

func TimelineRoutes(router *gin.Engine, controller *controllers.TimelineController, adminMiddleware *middleware.AdminMiddleware) {
	timelineGroup := router.Group("/api/timeline")

	timelineGroup.GET("", controller.GetTimelines)

	protected := timelineGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", controller.CreateTimeline)
		protected.PUT("/:id", controller.UpdateTimeline)
		protected.DELETE("/:id", controller.DeleteTimeline)
	}
}
