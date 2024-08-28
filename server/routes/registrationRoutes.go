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

func InitRegistrations(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddleware, smtp lib.SMTPClient) {
	store := store.NewRegistrationStore(database)
	service := services.NewRegistrationService(store, smtp)
	handler := handlers.NewRegistrationHandler(service)

	RegistrationRoutes(router, handler, authMiddleware)
}

func RegistrationRoutes(router *gin.Engine, handler *handlers.RegistrationHandler, authMiddleware *middleware.AuthMiddleware) {
	registrationGroup := router.Group("/api/registrations")

	registrationGroup.GET("", handler.GetRegistrations)
	registrationGroup.GET("/:id", handler.GetRegistrationById)
	registrationGroup.POST("", handler.CreateRegistration)

	protected := registrationGroup.Group("").Use(authMiddleware.Middleware())
	{
		protected.PUT("/:id", handler.UpdateRegistration)
		protected.DELETE("/:id", handler.DeleteRegistration)
	}
}
