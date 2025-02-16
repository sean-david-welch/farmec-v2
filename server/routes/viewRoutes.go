package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func ViewRoutes(router *gin.Engine, carouselHandler *handlers.AboutHandler, authMiddleware *middleware.AuthMiddlewareImpl, supplierCahce *middleware.SupplierCache) {
	router.GET("/", authMiddleware.ViewMiddleware(), func(c *gin.Context) {})
}
