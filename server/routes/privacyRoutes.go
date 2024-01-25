package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitPrivacy(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewPrivacyRepository(database)
	service := services.NewPrivacyService(repository)
	controller := controllers.NewPrivacyController(service)

	PrivacyRoutes(router, controller, adminMiddleware)
}

func PrivacyRoutes(router *gin.Engine, controller *controllers.PrivacyController, adminMiddleware *middleware.AdminMiddleware) {
	privacyGroup := router.Group("/api/privacy")

	privacyGroup.GET("", controller.GetPrivacys)

	protected := privacyGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", controller.CreatePrivacy)
		protected.PUT("/:id", controller.UpdatePrivacy)
		protected.DELETE("/:id", controller.DeletePrivacy)
	}
}
