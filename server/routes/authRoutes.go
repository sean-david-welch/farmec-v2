package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitAuth(router *gin.Engine, firebase *lib.Firebase, adminMiddleware *middleware.AdminMiddleware) {
	service := services.NewAuthService(firebase)
	handler := handlers.NewAuthHandler(service)

	AuthRoutes(router, handler, adminMiddleware)
}

func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler, adminMiddleware *middleware.AdminMiddleware) {
	authGroup := router.Group("/api/auth")

	authGroup.GET("/logout", handler.Logout)
	authGroup.GET("/login", handler.Login)

	protected := authGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.GET("/users", handler.GetUsers)
		protected.POST("/users", handler.Register)
		protected.PUT("/users/:uid", handler.UpdateUser)
		protected.DELETE("/users/:uid", handler.DeleteUser)
	}
}
