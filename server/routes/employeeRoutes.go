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

func InitializeEmployee(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	employeeRepository := repository.NewEmployeeRepository(database)
	service := services.NewEmployeeService(employeeRepository, s3Client, "Employees")
	controller := handlers.NewEmployeeController(service)

	EmployeeRoutes(router, controller, adminMiddleware)
}

func EmployeeRoutes(router *gin.Engine, controller *handlers.EmployeeController, middleware *middleware.AdminMiddleware) {
	employeeGroup := router.Group("/api/employees")

	employeeGroup.GET("", controller.GetEmployees)

	protected := employeeGroup.Group("").Use(middleware.Middleware())
	{
		protected.POST("", controller.CreateEmployee)
		protected.PUT("/:id", controller.UpdateEmployee)
		protected.DELETE("/:id", controller.DeleteEmployee)
	}
}
