package controllers

import (
	"net/http"

	"githib.com/sean-david-welch/Farmec-Astro/api/services"
	"github.com/gin-gonic/gin"
)

type SuppliersController struct {
	supplierService *services.SupplierService
}

func NewSuppliersContoller(supplierService *services.SupplierService) *SuppliersController {
	return &SuppliersController{supplierService: supplierService}
}

func (controller *SuppliersController) GetSuppliers(context *gin.Context) {
	suppliers, err := controller.supplierService.GetSuppliers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	context.JSON(http.StatusOK, suppliers)
}