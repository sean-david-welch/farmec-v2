package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (controller *ProductController) GetProducts(context *gin.Context) {
	id := context.Param("id")

	products, err := controller.productService.GetProducts(id)
	if err != nil {
		log.Printf("Error getting machines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while getting products"})
		return
	}

	context.JSON(http.StatusOK, products)
}

func (controller *ProductController) CreateProduct(context *gin.Context) {
	var product types.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	result, err := controller.productService.CreateProduct(&product)
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

func (controller *ProductController) UpdateProduct(context *gin.Context) {
	id := context.Param("id")

	var product types.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body", "details": err.Error()})
		return
	}

	result, err := controller.productService.UpdateProduct(id, &product)
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

func (controller *ProductController) DeleteProduct(context *gin.Context) {
	id := context.Param("id")

	if err := controller.productService.DeleteProduct(id); err != nil {
		log.Printf("Error deleting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "id": id})
}
