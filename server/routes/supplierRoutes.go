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

func InitSuppliers(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	supplierStore := stores.NewSupplierStore(database)
	supplierService := services.NewSupplierService(supplierStore, s3Client, "Suppliers")
	supplierHandler := handlers.NewSupplierContoller(supplierService)

	SupplierRoutes(router, supplierHandler, adminMiddleware)
}

func SupplierRoutes(router *gin.Engine, supplierHandler *handlers.SupplierHandler, adminMiddleware *middleware.AdminMiddleware) {
	supplierGroup := router.Group("/api/suppliers")

	supplierGroup.GET("", supplierHandler.GetSuppliers)
	supplierGroup.GET("/:id", supplierHandler.GetSupplierByID)

	protected := supplierGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", supplierHandler.CreateSupplier)
		protected.PUT("/:id", supplierHandler.UpdateSupplier)
		protected.DELETE("/:id", supplierHandler.DeleteSupplier)
	}
}
