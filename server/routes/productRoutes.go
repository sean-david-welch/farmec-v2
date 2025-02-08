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

func InitProduct(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AuthMiddlewareImpl) {
	repo := repository.NewProductRepo(database)
	service := services.NewProductService(repo, s3Client, "Products")
	handler := handlers.NewProductHandler(service)

	ProductRoutes(router, handler, adminMiddleware)
}

func ProductRoutes(router *gin.Engine, handler *handlers.ProductHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	productGroup := router.Group("/api/products")

	productGroup.GET("/:id", handler.GetProducts)

	protected := productGroup.Group("").Use(adminMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateProduct)
		protected.PUT("/:id", handler.UpdateProduct)
		protected.DELETE("/:id", handler.DeleteProduct)
	}
}
