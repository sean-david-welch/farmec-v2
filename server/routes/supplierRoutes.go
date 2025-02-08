package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func SupplierRoutes(router *gin.Engine, handler *handlers.SupplierHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	supplierGroup := router.Group("/api/suppliers")

	supplierGroup.GET("", handler.GetSuppliers)
	supplierGroup.GET("/:id", handler.GetSupplierByID)

	protected := supplierGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.POST("", handler.CreateSupplier)
		protected.PUT("/:id", handler.UpdateSupplier)
		protected.DELETE("/:id", handler.DeleteSupplier)
	}
}
