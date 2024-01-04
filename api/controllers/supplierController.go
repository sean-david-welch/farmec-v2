package controllers

import (
	"log"
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
        log.Printf("Error getting suppliers: %v", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
        return
	}
	context.JSON(http.StatusOK, suppliers)
}

func (controller *SuppliersController) GetSupplierByID(context *gin.Context) {
	id := context.Param("id")
	supplier, err := controller.supplierService.GetSupplierByID(id)

	if err != nil {
        log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}

	context.JSON(http.StatusOK, supplier)
}