package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/stores"
)

func InitWarranty(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddleware) {
	warrantyStore := stores.NewWarrantyStore(database)
	service := services.NewWarrantyService(warrantyStore)
	handler := handlers.NewWarrantyHandler(service)

	WarrantyRoutes(router, handler, authMiddleware)
}

func WarrantyRoutes(router *gin.Engine, handler *handlers.WarrantyHandler, authMiddleware *middleware.AuthMiddleware) {
	warrantyGroup := router.Group("/api/warranty")

	warrantyGroup.GET("", handler.GetWarranties)
	warrantyGroup.GET("/:id", handler.GetWarrantyById)
	warrantyGroup.POST("", handler.CreateWarranty)

	protected := warrantyGroup.Group("").Use(authMiddleware.Middleware())
	{
		protected.PUT("/:id", handler.UpdateWarranty)
		protected.DELETE("/:id", handler.DeleteWarranty)
	}
}
