package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitExhibitions(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) {
	repo := repository.NewExhibitionRepo(database)
	service := services.NewExhibitionService(repo)
	handler := handlers.NewExhibitionHandler(service, authMiddleware, supplierCache)

	ExhibitionRoutes(router, handler, authMiddleware)
}

func ExhibitionRoutes(router *gin.Engine, handler *handlers.ExhibitionHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/exhibitions", authMiddleware.ViewMiddleware(), handler.ExhibitionsView)

	exhibitionGroup := router.Group("/api/exhibitions")
	exhibitionGroup.GET("", handler.GetExhibitions)

	protected := exhibitionGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateExhibition)
		protected.PUT("/:id", handler.UpdateExhibition)
		protected.DELETE("/:id", handler.DeleteExhibition)
	}
}
