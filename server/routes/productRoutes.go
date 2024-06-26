package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/stores"
)

func InitProduct(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	productStore := stores.NewProductStore(database)
	productService := services.NewProductService(productStore, s3Client, "Products")
	productHandler := handlers.NewProductHandler(productService)

	ProductRoutes(router, productHandler, adminMiddleware)
}

func ProductRoutes(router *gin.Engine, productHandler *handlers.ProductHandler, adminMiddleware *middleware.AdminMiddleware) {
	productGroup := router.Group("/api/products")

	productGroup.GET("/:id", productHandler.GetProducts)

	protected := productGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", productHandler.CreateProduct)
		protected.PUT("/:id", productHandler.UpdateProduct)
		protected.DELETE("/:id", productHandler.DeleteProduct)
	}
}
