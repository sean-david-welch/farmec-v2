package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitializeResources(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, authMiddleware *middleware.AuthMiddlewareImpl) {
	// Initialize Employee resources
	employeeRepo := repository.NewEmployeeRepo(database)
	employeeService := services.NewEmployeeService(employeeRepo, s3Client, "Employees")
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	// Initialize Timeline resources
	timelineRepo := repository.NewTimelineRepo(database)
	timelineService := services.NewTimelineService(timelineRepo)
	timelineHandler := handlers.NewTimelineHandler(timelineService)

	// Setup routes for both resources
	SetupRoutes(router, employeeHandler, timelineHandler, authMiddleware)
}

func SetupRoutes(
	router *gin.Engine,
	employeeHandler *handlers.EmployeeHandler,
	timelineHandler *handlers.TimelineHandler,
	middleware *middleware.AuthMiddlewareImpl,
) {
	// Employee routes
	employeeGroup := router.Group("/api/employees")
	employeeGroup.GET("", employeeHandler.GetEmployees)

	protectedEmployeeRoutes := employeeGroup.Group("").Use(middleware.AdminRouteMiddleware())
	{
		protectedEmployeeRoutes.POST("", employeeHandler.CreateEmployee)
		protectedEmployeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		protectedEmployeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}

	// Timeline routes
	timelineGroup := router.Group("/api/timeline")
	timelineGroup.GET("", timelineHandler.GetTimelines)

	protectedTimelineRoutes := timelineGroup.Group("").Use(middleware.AdminRouteMiddleware())
	{
		protectedTimelineRoutes.POST("", timelineHandler.CreateTimeline)
		protectedTimelineRoutes.PUT("/:id", timelineHandler.UpdateTimeline)
		protectedTimelineRoutes.DELETE("/:id", timelineHandler.DeleteTimeline)
	}
}
