package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
)

func CheckoutRoutes(router *gin.Engine, handler *handlers.CheckoutHandler) {
	checkoutGroup := router.Group("/api/checkout")

	checkoutGroup.POST("/create-checkout-session/:id", handler.CreateCheckoutSession)
	checkoutGroup.GET("/session-status", handler.RetrieveCheckoutSession)
}
