package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitExhibitions(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewExhibitionRepository(database)
	service := services.NewExhibitionService(repository)
	controller := controllers.NewExhibitionController(service)

	ExhibitionRoutes(router, controller, adminMiddleware)
}

func ExhibitionRoutes(router *gin.Engine, controller *controllers.ExhibitionController, adminMiddleware *middleware.AdminMiddleware) {
	exhibitionGroup := router.Group("/api/exhibitions")

	exhibitionGroup.GET("", controller.GetExhibitions)

	protected := exhibitionGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", controller.CreateExhibition)
		protected.PUT("/:id", controller.UpdateExhibition)
		protected.DELETE("/:id", controller.DeleteExhibition)
	}
}