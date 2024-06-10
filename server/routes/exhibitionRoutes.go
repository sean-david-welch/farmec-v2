package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/store"
)

func InitExhibitions(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	exhibitionRepository := store.NewExhibitionRepository(database)
	service := services.NewExhibitionService(exhibitionRepository)
	handler := handlers.NewExhibitionHandler(service)

	ExhibitionRoutes(router, handler, adminMiddleware)
}

func ExhibitionRoutes(router *gin.Engine, handler *handlers.ExhibitionHandler, adminMiddleware *middleware.AdminMiddleware) {
	exhibitionGroup := router.Group("/api/exhibitions")

	exhibitionGroup.GET("", handler.GetExhibitions)

	protected := exhibitionGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", handler.CreateExhibition)
		protected.PUT("/:id", handler.UpdateExhibition)
		protected.DELETE("/:id", handler.DeleteExhibition)
	}
}
