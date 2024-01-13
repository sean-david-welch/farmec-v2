package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type BlogController struct {
	service *services.BlogService
}

func NewBlogController(service *services.BlogService) *BlogController {
	return &BlogController{service: service}
}

func(controller *BlogController) GetBlogs(context *gin.Context) {
	blogs, err := controller.service.GetBlogs(); if err != nil {
		log.Printf("error getting blogs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting blogs"})
		return
	}

	context.JSON(http.StatusOK, blogs)
}

func(controller *BlogController) GetBlogByID(context *gin.Context) {
	id := context.Param("id")
	blog, err := controller.service.GetBlogsByID(id); if err != nil {
		log.Printf("error getting blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting blog"})
		return
	}

	context.JSON(http.StatusOK, blog)
}

func(controller *BlogController) CreateBlog(context *gin.Context) {
	var blog models.Blog

	if err := context.ShouldBindJSON(&blog); err != nil {
		log.Printf("error while creating blog: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating blog"})
		return
	}

	result, err := controller.service.CreateBlog(&blog); if err != nil {
		log.Printf("error while creating blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating blog"})
		return
	}

	response := gin.H{"blog": blog, "presginedUrl": result.PresginedUrl, "imageUrl": result.ImageUrl}

	context.JSON(http.StatusCreated, response)
}

func(controller *BlogController) UpdateBlog(context *gin.Context) {
	id := context.Param("id")
	
	var blog models.Blog

	if err := context.ShouldBindJSON(&blog); err != nil {
		log.Printf("error while updating blog: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating blog"})
		return
	}

	result, err := controller.service.UpdateBlog(id, &blog); if err != nil {
		log.Printf("error while updating blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating blog"})
		return
	}

	response := gin.H{"blog": blog, "presginedUrl": result.PresginedUrl, "imageUrl": result.ImageUrl}

	context.JSON(http.StatusAccepted, response)
}

func(controller *BlogController) DeleteBlog(context *gin.Context) {
	id := context.Param("id")

	if err := controller.service.DeleteBlog(id); err != nil {
		log.Printf("error while deleting blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting blog"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "blog successfully deleted", "id": id})
}
