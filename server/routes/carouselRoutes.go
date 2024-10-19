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
	carouselRepo := repository.NewCarouselRepo(database)
	carouselService := services.NewCarouselService(carouselRepo, s3Client, "Carousels")
	carouselHandler := handlers.NewCarouselHandler(carouselService)

	CarouselRoutes(router, carouselHandler, adminMiddleare)
}

func CarouselRoutes(router *gin.Engine, carouselHandler *handlers.CarouselHandler, adminMiddleware *middleware.AdminMiddleware) {
	carouselGroup := router.Group("/api/carousels")

	carouselGroup.GET("", carouselHandler.GetCarousels)

	protected := carouselGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", carouselHandler.CreateCarousel)
		protected.PUT("/:id", carouselHandler.UpdateCarousel)
		protected.DELETE("/:id", carouselHandler.DeleteCarousel)
	}
}
