package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func VideoRoutes(router *gin.Engine, handler *handlers.VideoHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	videoGroup := router.Group("/api/videos")

	videoGroup.GET("/:id", handler.GetVideos)

	protected := videoGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateVideo)
		protected.PUT("/:id", handler.UpdateVideo)
		protected.DELETE("/:id", handler.DeleteVideo)
	}
}
