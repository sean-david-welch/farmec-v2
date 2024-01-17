package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitializeCheckout(router *gin.Engine, database *sql.DB, secrets *config.Secrets) {
	repository := repository.NewLineItemRepository(database)
	service := services.NewCheckoutService(secrets, repository)
	controller := controllers.NewCheckoutController(service)

	CheckoutRoutes(router, controller)
}

func CheckoutRoutes(router *gin.Engine, controller *controllers.CheckoutController) {
	checkoutGroup := router.Group("/api/checkout")

	checkoutGroup.POST("/create-checkout-session", controller.CreateCheckoutSession)
	checkoutGroup.POST("/session-status", controller.RetrieveCheckoutSession)
}