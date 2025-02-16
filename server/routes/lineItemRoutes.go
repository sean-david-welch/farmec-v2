package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitLineItems(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, authMiddleware *middleware.AuthMiddlewareImpl) {
	repo := repository.NewLineItemRepo(database)
	service := services.NewLineItemService(repo, s3Client, "Lineitems")
	handler := handlers.NewLineItemHandler(service)

	LineItemRoutes(router, handler, authMiddleware)
}

func LineItemRoutes(router *gin.Engine, handler *handlers.LineItemHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	lineItemGroup := router.Group("/api/lineitems")

	lineItemGroup.GET("", handler.GetLineItems)
	lineItemGroup.GET("/:id", handler.GetLineItemById)

	protecteed := lineItemGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protecteed.POST("", handler.CreateLineItem)
		protecteed.PUT("/:id", handler.UpdateLineItem)
		protecteed.DELETE("/:id", handler.DeleteLineItem)
	}
}
