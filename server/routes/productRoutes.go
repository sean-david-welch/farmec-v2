package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func ProductRoutes(router *gin.Engine, handler *handlers.ProductHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	productGroup := router.Group("/api/products")

	productGroup.GET("/:id", handler.GetProducts)

	protected := productGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateProduct)
		protected.PUT("/:id", handler.UpdateProduct)
		protected.DELETE("/:id", handler.DeleteProduct)
	}
}
