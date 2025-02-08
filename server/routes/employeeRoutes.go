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

func InitializeEmployee(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AuthMiddleware) {
	repo := repository.NewEmployeeRepo(database)
	service := services.NewEmployeeService(repo, s3Client, "Employees")
	handler := handlers.NewEmployeeHandler(service)

	EmployeeRoutes(router, handler, adminMiddleware)
}

func EmployeeRoutes(router *gin.Engine, handler *handlers.EmployeeHandler, middleware *middleware.AuthMiddleware) {
	employeeGroup := router.Group("/api/employees")

	employeeGroup.GET("", handler.GetEmployees)

	protected := employeeGroup.Group("").Use(middleware.RouteMiddleware())
	{
		protected.POST("", handler.CreateEmployee)
		protected.PUT("/:id", handler.UpdateEmployee)
		protected.DELETE("/:id", handler.DeleteEmployee)
	}
}
