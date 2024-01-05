package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/sean-david-welch/farmec-v2/server/controllers"
)

func SupplierRoutes(router *gin.Engine, supplierController *controllers.SuppliersController) {
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