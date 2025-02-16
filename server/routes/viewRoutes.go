package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func ViewRoutes(router *gin.Engine, handler *handlers.ViewHandler, authMiddleware *middleware.AuthMiddlewareImpl, supplierCahce *middleware.SupplierCache) {
	router.GET("/", authMiddleware.ViewMiddleware(), func(c *gin.Context) {})
	router.GET("/carousel-admin", authMiddleware.ViewMiddleware(), handler.CarouselAdminView)
}
