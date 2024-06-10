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

func InitParts(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	partsRepository := repository.NewPartsRepository(database)
	partsService := services.NewPartsService(partsRepository, s3Client, "Spareparts")
	partsController := handlers.NewPartsController(partsService)

	PartsRoutes(router, partsController, adminMiddleware)
}

func PartsRoutes(router *gin.Engine, partsController *handlers.PartsController, adminMiddleware *middleware.AdminMiddleware) {
	partsGroup := router.Group("/api/spareparts")

	partsGroup.GET("/:id", partsController.GetParts)

	protected := partsGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", partsController.CreateParts)
		protected.PUT("/:id", partsController.UpdateParts)
		protected.DELETE("/:id", partsController.DeletePart)
	}
}
