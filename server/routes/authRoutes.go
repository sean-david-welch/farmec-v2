package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	authGroup := router.Group("/api/auth")

	authGroup.GET("/logout", handler.Logout)
	authGroup.GET("/login", handler.Login)

	protected := authGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.GET("/users", handler.GetUsers)
		protected.POST("/users", handler.Register)
		protected.PUT("/users/:uid", handler.UpdateUser)
		protected.DELETE("/users/:uid", handler.DeleteUser)
	}
}
