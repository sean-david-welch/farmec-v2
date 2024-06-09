package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitCheckout(router *gin.Engine, database *sql.DB, secrets *config.Secrets) {
	itemRepository := repository.NewLineItemRepository(database)
	service := services.NewCheckoutService(secrets, itemRepository)
	controller := controllers.NewCheckoutController(service)

	CheckoutRoutes(router, controller)
}

func CheckoutRoutes(router *gin.Engine, controller *controllers.CheckoutController) {
	checkoutGroup := router.Group("/api/checkout")

	checkoutGroup.POST("/create-checkout-session/:id", controller.CreateCheckoutSession)
	checkoutGroup.GET("/session-status", controller.RetrieveCheckoutSession)
}
