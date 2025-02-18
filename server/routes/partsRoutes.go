package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func PartsRoutes(router *gin.Engine, handler *handlers.PartsHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/spareparts", authMiddleware.ViewMiddleware(), handler.PartsListView)
	router.GET("/spareparts/:id", authMiddleware.ViewMiddleware(), handler.PartsDetailView)

	partsGroup := router.Group("/api/spareparts")
	partsGroup.GET("/:id", handler.GetParts)
	protected := partsGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateParts)
		protected.PUT("/:id", handler.UpdateParts)
		protected.DELETE("/:id", handler.DeletePart)
	}
}
