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

func InitializeProduct(router *gin.Engine, db *sql.DB, s3Client *utils.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	productRepository := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepository, s3Client, "products")
	productController := controllers.NewProductController(productService)

	ProductRoutes(router, productController, adminMiddleware)
}

func ProductRoutes(router *gin.Engine, productController *controllers.ProductController, adminMiddleware *middleware.AdminMiddleware) {
	productGroup := router.Group("/api/products")

	productGroup.GET("/:id", productController.GetProducts)

	protected := productGroup.Group("").Use(adminMiddleware.Middleware()); {
		protected.POST("", productController.CreateProduct)
		protected.PUT("/:id", productController.UpdateProduct)
		protected.DELETE("/:id", productController.DeleteProduct)
	}
}