package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitAuth(router *gin.Engine, firebase *lib.Firebase, adminMiddleware *middleware.AdminMiddleware) {
	service := services.NewAuthService(firebase)
	controller := controllers.NewAuthController(service)

	AuthRoutes(router, controller, adminMiddleware)
}

func AuthRoutes(router *gin.Engine, controller *controllers.AuthController, adminMiddleware *middleware.AdminMiddleware) {
	authGroup := router.Group("/api/auth")

	authGroup.GET("/logout", controller.Logout)
	authGroup.GET("/login", controller.Login)

	protected := authGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.GET("/users", controller.GetUsers)
		protected.POST("/users", controller.Register)
		protected.PUT("/users/:uid", controller.UpdateUser)
		protected.DELETE("/users/:uid", controller.DeleteUser)
	}
}
