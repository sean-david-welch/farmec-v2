package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitLineItems(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	itemRepository := repository.NewLineItemRepository(database)
	service := services.NewLineItemService(itemRepository, s3Client, "Lineitems")
	controller := controllers.NewLineItemController(service)

	LineItemRoutes(router, controller, adminMiddleware)
}

func LineItemRoutes(router *gin.Engine, controller *controllers.LineItemController, adminMiddleware *middleware.AdminMiddleware) {
	lineItemGroup := router.Group("/api/lineitems")

	lineItemGroup.GET("", controller.GetLineItems)
	lineItemGroup.GET("/:id", controller.GetLineItemById)

	protecteed := lineItemGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protecteed.POST("", controller.CreateLineItem)
		protecteed.PUT("/:id", controller.UpdateLineItem)
		protecteed.DELETE("/:id", controller.DeleteLineItem)
	}
}
