package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func LineItemRoutes(router *gin.Engine, handler *handlers.LineItemHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	lineItemGroup := router.Group("/api/lineitems")

	lineItemGroup.GET("", handler.GetLineItems)
	lineItemGroup.GET("/:id", handler.GetLineItemById)

	protecteed := lineItemGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protecteed.POST("", handler.CreateLineItem)
		protecteed.PUT("/:id", handler.UpdateLineItem)
		protecteed.DELETE("/:id", handler.DeleteLineItem)
	}
}
