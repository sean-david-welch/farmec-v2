package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/Farmec-Astro/api/models"
	"github.com/sean-david-welch/Farmec-Astro/api/services"
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

func (controller *SuppliersController) CreateSupplier(context *gin.Context) {
	var supplier models.Supplier
	
	if err := context.ShouldBindJSON(&supplier); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := controller.supplierService.CreateSupplier(&supplier)
	if err != nil {
		log.Printf("Error creating supplier: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating supplier"})
		return
	}

	context.JSON(http.StatusCreated, supplier)
}