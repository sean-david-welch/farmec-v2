package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitParts(router *gin.Engine, database *sql.DB, s3Client utils.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	partsRepository := repository.NewPartsRepository(database)
	partsService := services.NewPartsService(partsRepository, s3Client, "Spareparts")
	partsController := controllers.NewPartsController(partsService)

	PartsRoutes(router, partsController, adminMiddleware)
}

func PartsRoutes(router *gin.Engine, partsController *controllers.PartsController, adminMiddleware *middleware.AdminMiddleware) {
	partsGroup := router.Group("/api/spareparts")

	partsGroup.GET("/:id", partsController.GetParts)

	protected := partsGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", partsController.CreateParts)
		protected.PUT("/:id", partsController.UpdateParts)
		protected.DELETE("/:id", partsController.DeletePart)
	}
}
