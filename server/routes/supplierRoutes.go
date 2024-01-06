package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeSuppliers(router *gin.Engine, db *sql.DB, s3Client *utils.S3Client) {
    supplierRepository := repository.NewSupplierRepository(db)
    supplierService := services.NewSupplierService(supplierRepository, s3Client, "suppliers")
    supplierController := controllers.NewSupplierContoller(supplierService)

    SupplierRoutes(router, supplierController)
}

func SupplierRoutes(router *gin.Engine, supplierController *controllers.SupplierController) {
	supplierGroup := router.Group("/api/suppliers")
	// supplierGroup.Use(AuthMiddleware())

	{
		supplierGroup.GET("", supplierController.GetSuppliers)
		supplierGroup.POST("", supplierController.CreateSupplier)
		supplierGroup.GET("/:id", supplierController.GetSupplierByID) 
		supplierGroup.PUT("/:id", supplierController.UpdateSupplier) 
		supplierGroup.DELETE("/:id", supplierController.DeleteSupplier) 

		// for sub resources 
		// supplierGroup.GET("/:id/products", supplierController.GetSupplierProducts)
	}
}