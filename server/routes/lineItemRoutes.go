package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitializeLineItems(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewLineItemRepository(database)
	service := services.NewLineItemService(repository)
	controller := controllers.NewLineItemController(service)

	LineItemRoutes(router, controller, adminMiddleware)
}

func LineItemRoutes(router *gin.Engine, controller *controllers.LineItemController, adminMiddleware *middleware.AdminMiddleware) {
	lineItemGroup := router.Group("/api/lineItems")

	lineItemGroup.GET("", controller.GetLineItems)
	lineItemGroup.GET("/:id", controller.GetLineItemById)

	protecteed := lineItemGroup.Group("").Use(adminMiddleware.Middleware()); {
		protecteed.POST("", controller.CreateLineItem)
		protecteed.PUT("/:id", controller.UpdateLineItem)
		protecteed.DELETE("/:id", controller.DeleteLineItem)
	}
}