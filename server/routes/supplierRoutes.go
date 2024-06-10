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

func InitSuppliers(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	supplierRepository := repository.NewSupplierRepository(database)
	supplierService := services.NewSupplierService(supplierRepository, s3Client, "Suppliers")
	supplierController := controllers.NewSupplierContoller(supplierService)

	SupplierRoutes(router, supplierController, adminMiddleware)
}

func SupplierRoutes(router *gin.Engine, supplierController *controllers.SupplierController, adminMiddleware *middleware.AdminMiddleware) {
	supplierGroup := router.Group("/api/suppliers")

	supplierGroup.GET("", supplierController.GetSuppliers)
	supplierGroup.GET("/:id", supplierController.GetSupplierByID)

	protected := supplierGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", supplierController.CreateSupplier)
		protected.PUT("/:id", supplierController.UpdateSupplier)
		protected.DELETE("/:id", supplierController.DeleteSupplier)
	}
}
