package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitializeRegistrations(router *gin.Engine, db *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewRegistrationRepository(db)
	service := services.NewRegistrationService(repository)
	controller := controllers.NewRegistrationController(service)

	RegistrationRoutes(router, controller, adminMiddleware)
}

func RegistrationRoutes(router *gin.Engine, controller *controllers.RegistrationController, adminMiddleware *middleware.AdminMiddleware) {
	registrationGroup := router.Group("/api/registrations")

	registrationGroup.GET("", controller.GetRegistrations)
	registrationGroup.GET("/:id", controller.GetRegistrationById)

	protecteed := registrationGroup.Group("").Use(adminMiddleware.Middleware()); {
		protecteed.POST("", controller.CreateRegistration)
		protecteed.PUT("/:id", controller.UpdateRegistration)
		protecteed.DELETE("/:id", controller.DeleteRegistration)
	}
}