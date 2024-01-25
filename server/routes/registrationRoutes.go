package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitRegistrations(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddleware) {
	repository := repository.NewRegistrationRepository(database)
	service := services.NewRegistrationService(repository)
	controller := controllers.NewRegistrationController(service)

	RegistrationRoutes(router, controller, authMiddleware)
}

func RegistrationRoutes(router *gin.Engine, controller *controllers.RegistrationController, authMiddleware *middleware.AuthMiddleware) {
	registrationGroup := router.Group("/api/registrations")

	registrationGroup.GET("", controller.GetRegistrations)
	registrationGroup.GET("/:id", controller.GetRegistrationById)

	protecteed := registrationGroup.Group("").Use(authMiddleware.Middleware())
	{
		protecteed.POST("", controller.CreateRegistration)
		protecteed.PUT("/:id", controller.UpdateRegistration)
		protecteed.DELETE("/:id", controller.DeleteRegistration)
	}
}
