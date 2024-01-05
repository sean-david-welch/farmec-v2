package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/sean-david-welch/farmec-v2/server/controllers"
)

func SupplierRoutes(router *gin.Engine, supplierController *controllers.SuppliersController) {
	supplierGroup := router.Group("/suppliers")
	// supplierGroup.Use(AuthMiddleware())

	{
		supplierGroup.GET("", supplierController.GetSuppliers)
		supplierGroup.POST("", supplierController.CreateSupplier)
		supplierGroup.GET("/:id", supplierController.GetSupplierByID) 



		// for sub resources 
		// supplierGroup.GET("/:id/products", supplierController.GetSupplierProducts)
	}
}