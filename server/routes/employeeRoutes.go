package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/store"
)

func InitializeEmployee(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	employeeRepository := store.NewEmployeeRepository(database)
	service := services.NewEmployeeService(employeeRepository, s3Client, "Employees")
	handler := handlers.NewEmployeeHandler(service)

	EmployeeRoutes(router, handler, adminMiddleware)
}

func EmployeeRoutes(router *gin.Engine, handler *handlers.EmployeeHandler, middleware *middleware.AdminMiddleware) {
	employeeGroup := router.Group("/api/employees")

	employeeGroup.GET("", handler.GetEmployees)

	protected := employeeGroup.Group("").Use(middleware.Middleware())
	{
		protected.POST("", handler.CreateEmployee)
		protected.PUT("/:id", handler.UpdateEmployee)
		protected.DELETE("/:id", handler.DeleteEmployee)
	}
}
