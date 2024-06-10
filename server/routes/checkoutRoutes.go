package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitCheckout(router *gin.Engine, database *sql.DB, secrets *lib.Secrets) {
	itemRepository := repository.NewLineItemRepository(database)
	service := services.NewCheckoutService(secrets, itemRepository)
	controller := handlers.NewCheckoutController(service)

	CheckoutRoutes(router, controller)
}

func CheckoutRoutes(router *gin.Engine, controller *handlers.CheckoutHandler) {
	checkoutGroup := router.Group("/api/checkout")

	checkoutGroup.POST("/create-checkout-session/:id", controller.CreateCheckoutSession)
	checkoutGroup.GET("/session-status", controller.RetrieveCheckoutSession)
}
