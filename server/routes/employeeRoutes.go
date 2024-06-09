package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeEmployee(router *gin.Engine, database *sql.DB, s3Client utils.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	employeeRepository := repository.NewEmployeeRepository(database)
	service := services.NewEmployeeService(employeeRepository, s3Client, "Employees")
	controller := controllers.NewEmployeeController(service)

	EmployeeRoutes(router, controller, adminMiddleware)
}

func EmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController, middleware *middleware.AdminMiddleware) {
	employeeGroup := router.Group("/api/employees")

	employeeGroup.GET("", controller.GetEmployees)

	protected := employeeGroup.Group("").Use(middleware.Middleware())
	{
		protected.POST("", controller.CreateEmployee)
		protected.PUT("/:id", controller.UpdateEmployee)
		protected.DELETE("/:id", controller.DeleteEmployee)
	}
}
