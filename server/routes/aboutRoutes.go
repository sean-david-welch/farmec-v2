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

func InitAbout(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) {
	// Initialize Employee resources
	employeeRepo := repository.NewEmployeeRepo(database)
	employeeService := services.NewEmployeeService(employeeRepo, s3Client, "Employees")
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	// Initialize Timeline resources
	timelineRepo := repository.NewTimelineRepo(database)
	timelineService := services.NewTimelineService(timelineRepo)
	timelineHandler := handlers.NewTimelineHandler(timelineService)

	// init routes
	aboutHandler := handlers.NewAboutHandler(employeeService, timelineService, authMiddleware, supplierCache)
	AboutRoutes(router, aboutHandler, employeeHandler, timelineHandler, authMiddleware)
}

func AboutRoutes(router *gin.Engine, aboutHandler *handlers.AboutHandler, employeeHandler *handlers.EmployeeHandler, timelineHandler *handlers.TimelineHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/about", authMiddleware.ViewMiddleware(), aboutHandler.AboutView)
	// Employee routes
	employeeGroup := router.Group("/api/employees")
	employeeGroup.GET("", employeeHandler.GetEmployees)

	protectedEmployeeRoutes := employeeGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protectedEmployeeRoutes.POST("", employeeHandler.CreateEmployee)
		protectedEmployeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		protectedEmployeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}

	// Timeline routes
	timelineGroup := router.Group("/api/timeline")
	timelineGroup.GET("", timelineHandler.GetTimelines)

	protectedTimelineRoutes := timelineGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protectedTimelineRoutes.POST("", timelineHandler.CreateTimeline)
		protectedTimelineRoutes.PUT("/:id", timelineHandler.UpdateTimeline)
		protectedTimelineRoutes.DELETE("/:id", timelineHandler.DeleteTimeline)
	}
}
