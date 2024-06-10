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

func InitCarousel(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleare *middleware.AdminMiddleware) {
	carouselRepository := repository.NewCarouselRepository(database)
	carouselService := services.NewCarouselService(carouselRepository, s3Client, "Carousels")
	carouselController := handlers.NewCarouselController(carouselService)

	CarouselRoutes(router, carouselController, adminMiddleare)
}

func CarouselRoutes(router *gin.Engine, carouselController *handlers.CarouselHandler, adminMiddleware *middleware.AdminMiddleware) {
	carouselGroup := router.Group("/api/carousels")

	carouselGroup.GET("", carouselController.GetCarousels)

	protected := carouselGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", carouselController.CreateCarousel)
		protected.PUT("/:id", carouselController.UpdateCarousel)
		protected.DELETE("/:id", carouselController.DeleteCarousel)
	}
}
