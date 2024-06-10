package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/store"
)

func InitCheckout(router *gin.Engine, database *sql.DB, secrets *lib.Secrets) {
	itemStore := store.NewLineItemStore(database)
	service := services.NewCheckoutService(secrets, itemStore)
	handler := handlers.NewCheckoutHandler(service)

	CheckoutRoutes(router, handler)
}

func CheckoutRoutes(router *gin.Engine, handler *handlers.CheckoutHandler) {
	checkoutGroup := router.Group("/api/checkout")

	checkoutGroup.POST("/create-checkout-session/:id", handler.CreateCheckoutSession)
	checkoutGroup.GET("/session-status", handler.RetrieveCheckoutSession)
}
