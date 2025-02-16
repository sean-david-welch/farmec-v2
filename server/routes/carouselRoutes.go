package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func CarouselRoutes(router *gin.Engine, handler *handlers.CarouselHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	carouselGroup := router.Group("/api/carousels")
	carouselGroup.GET("", handler.GetCarousels)
	protected := carouselGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateCarousel)
		protected.PUT("/:id", handler.UpdateCarousel)
		protected.DELETE("/:id", handler.DeleteCarousel)
	}
}
