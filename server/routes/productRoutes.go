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

func InitProduct(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	productRepository := repository.NewProductRepository(database)
	productService := services.NewProductService(productRepository, s3Client, "Products")
	productController := handlers.NewProductController(productService)

	ProductRoutes(router, productController, adminMiddleware)
}

func ProductRoutes(router *gin.Engine, productController *handlers.ProductController, adminMiddleware *middleware.AdminMiddleware) {
	productGroup := router.Group("/api/products")

	productGroup.GET("/:id", productController.GetProducts)

	protected := productGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", productController.CreateProduct)
		protected.PUT("/:id", productController.UpdateProduct)
		protected.DELETE("/:id", productController.DeleteProduct)
	}
}
