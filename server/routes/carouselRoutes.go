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

func InitCarousel(router *gin.Engine, database *sql.DB, s3Client utils.S3Client, adminMiddleare *middleware.AdminMiddleware) {
	carouselRepository := repository.NewCarouselRepository(database)
	carouselService := services.NewCarouselService(carouselRepository, s3Client, "Carousels")
	carouselController := controllers.NewCarouselController(carouselService)

	CarouselRoutes(router, carouselController, adminMiddleare)
}

func CarouselRoutes(router *gin.Engine, carouselController *controllers.CarouselController, adminMiddleware *middleware.AdminMiddleware) {
	carouselGroup := router.Group("/api/carousels")

	carouselGroup.GET("", carouselController.GetCarousels)

	protected := carouselGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", carouselController.CreateCarousel)
		protected.PUT("/:id", carouselController.UpdateCarousel)
		protected.DELETE("/:id", carouselController.DeleteCarousel)
	}
}
