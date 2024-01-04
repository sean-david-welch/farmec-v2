package routes

import (
	"github.com/gin-gonic/gin"

	"githib.com/sean-david-welch/Farmec-Astro/api/controllers"
)

func SupplierRoutes(router *gin.Engine, supplierController *controllers.SuppliersController) {
	supplierGroup := router.Group("/suppliers")
	// supplierGroup.Use(AuthMiddleware())

	{
		supplierGroup.GET("/", supplierController.GetSuppliers)
		supplierGroup.GET("/:id", supplierController.GetSupplierByID) 

		// for sub resources 
		// supplierGroup.GET("/:id/products", supplierController.GetSupplierProducts)
	}
}