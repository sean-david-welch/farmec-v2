package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitilizeEmployee(router *gin.Engine, db *sql.DB, s3Client *utils.S3Client, firebase *lib.Firebase) {
	repository := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repository, s3Client, "employees")
	cotroller := controllers.NewEmployeeController(service)

	admimMiddleware := middleware.NewAdminMiddleware(firebase)


	EmployeeRoutes(router, cotroller, admimMiddleware)
}

func EmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController, middleware *middleware.AdminMiddleware) {
	employeeGroup := router.Group("/api/employees")

	employeeGroup.GET("", controller.GetEmployees)

	protected := employeeGroup.Group("")
	protected.Use(middleware.Middleware()); {
		protected.POST("", controller.CreateEmployee)
		protected.PUT("/:id", controller.UpdateEmployee)
		protected.DELETE("/:id", controller.DeleteEmployee)
	}
}