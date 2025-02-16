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

func InitWarranty(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddlewareImpl, emailClient *lib.EmailClientImpl) {
	repo := repository.NewWarrantyRepo(database)
	service := services.NewWarrantyService(repo, emailClient)
	handler := handlers.NewWarrantyHandler(service)

	WarrantyRoutes(router, handler, authMiddleware)
}

func WarrantyRoutes(router *gin.Engine, handler *handlers.WarrantyHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	warrantyGroup := router.Group("/api/warranty")

	warrantyGroup.GET("", handler.GetWarranties)
	warrantyGroup.GET("/:id", handler.GetWarrantyById)
	warrantyGroup.POST("", handler.CreateWarranty)

	protected := warrantyGroup.Group("").Use(authMiddleware.AuthRouteMiddleware())
	{
		protected.PUT("/:id", handler.UpdateWarranty)
		protected.DELETE("/:id", handler.DeleteWarranty)
	}
}
