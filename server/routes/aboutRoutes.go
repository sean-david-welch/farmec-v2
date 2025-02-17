package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func AboutRoutes(router *gin.Engine, aboutHandler *handlers.AboutHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/about", authMiddleware.ViewMiddleware(), aboutHandler.AboutView)
	// Employee routes
	employeeGroup := router.Group("/api/employees")
	employeeGroup.GET("", aboutHandler.GetEmployees)
	protectedEmployeeRoutes := employeeGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protectedEmployeeRoutes.POST("", aboutHandler.CreateEmployee)
		protectedEmployeeRoutes.PUT("/:id", aboutHandler.UpdateEmployee)
		protectedEmployeeRoutes.DELETE("/:id", aboutHandler.DeleteEmployee)
	}

	// Timeline routes
	timelineGroup := router.Group("/api/timeline")
	timelineGroup.GET("", aboutHandler.GetTimelines)
	protectedTimelineRoutes := timelineGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protectedTimelineRoutes.POST("", aboutHandler.CreateTimeline)
		protectedTimelineRoutes.PUT("/:id", aboutHandler.UpdateTimeline)
		protectedTimelineRoutes.DELETE("/:id", aboutHandler.DeleteTimeline)
	}
}
