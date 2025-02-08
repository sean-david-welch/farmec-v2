package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitPrivacy(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AuthMiddleware) {
	repo := repository.NewPrivacyRepo(database)
	service := services.NewPrivacyService(repo)
	handler := handlers.NewPrivacyHandler(service)

	PrivacyRoutes(router, handler, adminMiddleware)
}

func PrivacyRoutes(router *gin.Engine, handler *handlers.PrivacyHandler, adminMiddleware *middleware.AuthMiddleware) {
	privacyGroup := router.Group("/api/privacy")

	privacyGroup.GET("", handler.GetPrivacys)

	protected := privacyGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.POST("", handler.CreatePrivacy)
		protected.PUT("/:id", handler.UpdatePrivacy)
		protected.DELETE("/:id", handler.DeletePrivacy)
	}
}
