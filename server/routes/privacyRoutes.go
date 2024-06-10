package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/store"
)

func InitPrivacy(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	privacyRepository := store.NewPrivacyRepository(database)
	service := services.NewPrivacyService(privacyRepository)
	handler := handlers.NewPrivacyHandler(service)

	PrivacyRoutes(router, handler, adminMiddleware)
}

func PrivacyRoutes(router *gin.Engine, handler *handlers.PrivacyHandler, adminMiddleware *middleware.AdminMiddleware) {
	privacyGroup := router.Group("/api/privacy")

	privacyGroup.GET("", handler.GetPrivacys)

	protected := privacyGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", handler.CreatePrivacy)
		protected.PUT("/:id", handler.UpdatePrivacy)
		protected.DELETE("/:id", handler.DeletePrivacy)
	}
}
