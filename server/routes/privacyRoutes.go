package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitPrivacy(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	privacyRepository := repository.NewPrivacyRepository(database)
	service := services.NewPrivacyService(privacyRepository)
	controller := handlers.NewPrivacyController(service)

	PrivacyRoutes(router, controller, adminMiddleware)
}

func PrivacyRoutes(router *gin.Engine, controller *handlers.PrivacyController, adminMiddleware *middleware.AdminMiddleware) {
	privacyGroup := router.Group("/api/privacy")

	privacyGroup.GET("", controller.GetPrivacys)

	protected := privacyGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", controller.CreatePrivacy)
		protected.PUT("/:id", controller.UpdatePrivacy)
		protected.DELETE("/:id", controller.DeletePrivacy)
	}
}
