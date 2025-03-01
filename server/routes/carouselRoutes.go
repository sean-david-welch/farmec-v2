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
	repo := repository.NewCarouselRepo(database)
	service := services.NewCarouselService(repo, s3Client, "Carousels")
	handler := handlers.NewCarouselHandler(service)

	CarouselRoutes(router, handler, adminMiddleare)
}

func CarouselRoutes(router *gin.Engine, handler *handlers.CarouselHandler, adminMiddleware *middleware.AdminMiddleware) {
	carouselGroup := router.Group("/api/carousels")

	carouselGroup.GET("", handler.GetCarousels)

	protected := carouselGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", handler.CreateCarousel)
		protected.PUT("/:id", handler.UpdateCarousel)
		protected.DELETE("/:id", handler.DeleteCarousel)
	}
}
