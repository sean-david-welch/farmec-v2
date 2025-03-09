package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func RegistrationRoutes(router *gin.Engine, handler *handlers.RegistrationHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/registrations", authMiddleware.ViewMiddleware(), handler.RegistrationsView)
	router.GET("/registrations/:id", authMiddleware.ViewMiddleware(), handler.RegistrationsDetailView)
	registrationGroup := router.Group("/api/registrations")
	{
		registrationGroup.GET("", handler.GetRegistrations)
		registrationGroup.GET("/:id", handler.GetRegistrationById)
		registrationGroup.POST("", handler.CreateRegistration)
	}
	protected := registrationGroup.Group("").Use(authMiddleware.AuthRouteMiddleware())
	{
		protected.PUT("/:id", handler.UpdateRegistration)
		protected.DELETE("/:id", handler.DeleteRegistration)
	}
}
