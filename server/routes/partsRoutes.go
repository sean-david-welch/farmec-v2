package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitParts(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AuthMiddlewareImpl) {
	repo := repository.NewPartsRepo(database)
	service := services.NewPartsService(repo, s3Client, "Spareparts")
	handler := handlers.NewPartsHandler(service)

	PartsRoutes(router, handler, adminMiddleware)
}

func PartsRoutes(router *gin.Engine, handler *handlers.PartsHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	partsGroup := router.Group("/api/spareparts")

	partsGroup.GET("/:id", handler.GetParts)

	protected := partsGroup.Group("").Use(adminMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateParts)
		protected.PUT("/:id", handler.UpdateParts)
		protected.DELETE("/:id", handler.DeletePart)
	}
}
