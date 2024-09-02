package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type BlogHandler struct {
	service services.BlogService
}

func NewBlogHandler(service services.BlogService) *BlogHandler {
	return &BlogHandler{service: service}
}

func (handler *BlogHandler) GetBlogs(context *gin.Context) {
	ctx := context.Request.Context()
	blogs, err := handler.service.GetBlogs(ctx)
	if err != nil {
		log.Printf("error getting blogs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting blogs"})
		return
	}

	context.JSON(http.StatusOK, blogs)
}

func (handler *BlogHandler) GetBlogByID(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")
	blog, err := handler.service.GetBlogsByID(ctx, id)
	if err != nil {
		log.Printf("error getting blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting blog"})
		return
	}

	context.JSON(http.StatusOK, blog)
}

func (handler *BlogHandler) CreateBlog(context *gin.Context) {
	ctx := context.Request.Context()
	var blog db.Blog

	if err := context.ShouldBindJSON(&blog); err != nil {
		log.Printf("error while creating blog: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating blog"})
		return
	}

	result, err := handler.service.CreateBlog(ctx, &blog)
	if err != nil {
		log.Printf("error while creating blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating blog"})
		return
	}

	response := gin.H{"blog": blog, "presignedUrl": result.PresignedUrl, "imageUrl": result.ImageUrl}

	context.JSON(http.StatusCreated, response)
}

func (handler *BlogHandler) UpdateBlog(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	var blog db.Blog

	if err := context.ShouldBindJSON(&blog); err != nil {
		log.Printf("error while updating blog: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating blog"})
		return
	}

	result, err := handler.service.UpdateBlog(ctx, id, &blog)
	if err != nil {
		log.Printf("error while updating blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating blog"})
		return
	}

	response := gin.H{"blog": blog, "presignedUrl": result.PresignedUrl, "imageUrl": result.ImageUrl}

	context.JSON(http.StatusAccepted, response)
}

func (handler *BlogHandler) DeleteBlog(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteBlog(ctx, id); err != nil {
		log.Printf("error while deleting blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting blog"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "blog successfully deleted", "id": id})
}
