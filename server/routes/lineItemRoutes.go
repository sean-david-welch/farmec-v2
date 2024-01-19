package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeLineItems(router *gin.Engine, database *sql.DB, s3Client *utils.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewLineItemRepository(database)
	service := services.NewLineItemService(repository, s3Client, "Lineitems")
	controller := controllers.NewLineItemController(service)

	LineItemRoutes(router, controller, adminMiddleware)
}

func LineItemRoutes(router *gin.Engine, controller *controllers.LineItemController, adminMiddleware *middleware.AdminMiddleware) {
	lineItemGroup := router.Group("/api/line-items")

	lineItemGroup.GET("", controller.GetLineItems)
	lineItemGroup.GET("/:id", controller.GetLineItemById)

	protecteed := lineItemGroup.Group("").Use(adminMiddleware.Middleware()); {
		protecteed.POST("", controller.CreateLineItem)
		protecteed.PUT("/:id", controller.UpdateLineItem)
		protecteed.DELETE("/:id", controller.DeleteLineItem)
	}
}