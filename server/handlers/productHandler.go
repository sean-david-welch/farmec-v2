package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (handler *ProductHandler) GetProducts(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	products, err := handler.productService.GetProducts(ctx, id)
	if err != nil {
		log.Printf("Error getting machines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while getting products"})
		return
	}

	context.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) CreateProduct(context *gin.Context) {
	ctx := context.Request.Context()
	var product types.Product
	dbProduct := lib.DeserializeProduct(product)

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if product.MachineID == "" || product.MachineID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	result, err := handler.productService.CreateProduct(ctx, &dbProduct)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating product", "details": err.Error()})
		return
	}

	response := gin.H{
		"product":      product,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *ProductHandler) UpdateProduct(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	var product types.Product
	dbProduct := lib.DeserializeProduct(product)

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "details": err.Error()})
		return
	}

	if product.MachineID == "" || product.MachineID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	result, err := handler.productService.UpdateProduct(ctx, id, &dbProduct)
	if err != nil {
		log.Printf("Error updating product: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating machine", "details": err.Error()})
		return
	}

	response := gin.H{
		"product":      product,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *ProductHandler) DeleteProduct(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.productService.DeleteProduct(ctx, id); err != nil {
		log.Printf("Error deleting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "id": id})
}
