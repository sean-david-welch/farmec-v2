package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeCarousel(router *gin.Engine, db *sql.DB, s3Client *utils.S3Client, firebase *lib.Firebase) {
	carouselRepository := repository.NewCarouselRepository(db)
	carouselService := services.NewCarouselService(carouselRepository, s3Client, "carousels")
	carouselController := controllers.NewCarouselController(carouselService)

	adminMiddleware := middleware.NewAdminMiddleware(firebase)

	CarouselRoutes(router, carouselController, adminMiddleware)
}

func CarouselRoutes(router *gin.Engine, carouselController *controllers.CarouselController, adminMiddleware *middleware.AdminMiddleware) {
	carouselGroup := router.Group("/api/carousels")

	carouselGroup.GET("", carouselController.GetCarousels)

	protected := carouselGroup.Group("")
	protected.Use(adminMiddleware.Middleware()); {
		protected.POST("", carouselController.CreateCarousel)
		protected.PUT("/:id", carouselController.UpdateCarousel)
		protected.DELETE("/:id", carouselController.DeleteCarousel)
	}
}