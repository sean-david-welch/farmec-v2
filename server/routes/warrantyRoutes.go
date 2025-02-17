package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

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
