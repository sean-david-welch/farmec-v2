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

func InitializeEmployee(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, authMiddleware *middleware.AuthMiddlewareImpl) {
	repo := repository.NewEmployeeRepo(database)
	service := services.NewEmployeeService(repo, s3Client, "Employees")
	handler := handlers.NewEmployeeHandler(service)

	EmployeeRoutes(router, handler, authMiddleware)
}

func EmployeeRoutes(router *gin.Engine, handler *handlers.EmployeeHandler, middleware *middleware.AuthMiddlewareImpl) {
	employeeGroup := router.Group("/api/employees")

	employeeGroup.GET("", handler.GetEmployees)

	protected := employeeGroup.Group("").Use(middleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateEmployee)
		protected.PUT("/:id", handler.UpdateEmployee)
		protected.DELETE("/:id", handler.DeleteEmployee)
	}
}

func InitTimelines(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddlewareImpl) {
	repo := repository.NewTimelineRepo(database)
	service := services.NewTimelineService(repo)
	handler := handlers.NewTimelineHandler(service)

	TimelineRoutes(router, handler, authMiddleware)
}

func TimelineRoutes(router *gin.Engine, handler *handlers.TimelineHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	timelineGroup := router.Group("/api/timeline")

	timelineGroup.GET("", handler.GetTimelines)

	protected := timelineGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateTimeline)
		protected.PUT("/:id", handler.UpdateTimeline)
		protected.DELETE("/:id", handler.DeleteTimeline)
	}
}

func InitAbout(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddlewareImpl) {}

func AboutRoutes(router *gin.Engine, handler *handlers.AboutHandler, middleware *middleware.AuthMiddlewareImpl) {
}
