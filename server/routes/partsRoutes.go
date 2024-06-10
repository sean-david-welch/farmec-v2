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

func InitParts(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	partsRepository := store.NewPartsRepository(database)
	partsService := services.NewPartsService(partsRepository, s3Client, "Spareparts")
	partsHandler := handlers.NewPartsHandler(partsService)

	PartsRoutes(router, partsHandler, adminMiddleware)
}

func PartsRoutes(router *gin.Engine, partsHandler *handlers.PartsHandler, adminMiddleware *middleware.AdminMiddleware) {
	partsGroup := router.Group("/api/spareparts")

	partsGroup.GET("/:id", partsHandler.GetParts)

	protected := partsGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", partsHandler.CreateParts)
		protected.PUT("/:id", partsHandler.UpdateParts)
		protected.DELETE("/:id", partsHandler.DeletePart)
	}
}
