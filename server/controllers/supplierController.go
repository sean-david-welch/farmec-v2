package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/services"
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

func (controller *SuppliersController) CreateSupplier(context *gin.Context) {
    var supplier models.Supplier
    
    if err := context.ShouldBindJSON(&supplier); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
        return
    }

    err := controller.supplierService.CreateSupplier(&supplier)
    if err != nil {
        log.Printf("Error creating supplier: %v", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating supplier", "details": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, supplier)
}

func (controller *SuppliersController) GetSupplierByID(context *gin.Context) {
	id := context.Param("id")
	supplier, err := controller.supplierService.GetSupplierById(id)

	if err != nil {
        log.Printf("Error getting suppliers: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting suppliers"})
		return
	}

	context.JSON(http.StatusOK, supplier)
}

func (controller *SuppliersController) UpdateSupplier(context *gin.Context) {
	id := context.Param("id")

	var supplier models.Supplier

	if err := context.ShouldBindJSON(&supplier); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return 
	}

	err := controller.supplierService.UpdateSupplier(id, &supplier)
	if err != nil {
		log.Printf("Error updating supplier: %v", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating supplier", "details": err.Error()})
        return
	}

	context.JSON(http.StatusAccepted, supplier)
}

func (controller *SuppliersController) DeleteSupplier(context *gin.Context) {
	id := context.Param("id")

	err := controller.supplierService.DeleteSupplier(id)
	if err != nil {
		log.Printf("Error deleting supplier: %v", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting supplier", "details": err.Error()})
        return
	}

    context.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully", "id": id})
}