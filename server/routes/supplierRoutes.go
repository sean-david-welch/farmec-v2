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

func InitSuppliers(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) services.SupplierService {
	repo := repository.NewSupplierRepo(database)
	service := services.NewSupplierService(repo, s3Client, "Suppliers")
	handler := handlers.NewSupplierContoller(service)

	SupplierRoutes(router, handler, adminMiddleware)
	return service
}

func SupplierRoutes(router *gin.Engine, handler *handlers.SupplierHandler, adminMiddleware *middleware.AdminMiddleware) {
	supplierGroup := router.Group("/api/suppliers")

	supplierGroup.GET("", handler.GetSuppliers)
	supplierGroup.GET("/:id", handler.GetSupplierByID)

	protected := supplierGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", handler.CreateSupplier)
		protected.PUT("/:id", handler.UpdateSupplier)
		protected.DELETE("/:id", handler.DeleteSupplier)
	}
}
