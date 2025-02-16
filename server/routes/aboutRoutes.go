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
	// Initialize Timeline resources
	timelineRepo := repository.NewTimelineRepo(database)
	timelineService := services.NewTimelineService(timelineRepo)

	// init routes
	aboutHandler := handlers.NewAboutHandler(employeeService, timelineService, authMiddleware, supplierCache)
	AboutRoutes(router, aboutHandler, authMiddleware)
}

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
