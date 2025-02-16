package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type BlogHandler struct {
	service         services.BlogService
	adminMiddleware *middleware.AuthMiddlewareImpl
	supplierCache   *middleware.SupplierCache
}

func NewBlogHandler(service services.BlogService, adminMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *BlogHandler {
	return &BlogHandler{service: service, adminMiddleware: adminMiddleware, supplierCache: supplierCache}
}

func (handler *BlogHandler) BlogsView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.adminMiddleware.GetIsAdmin(context)
	suppliers := middleware.GetSuppliersFromContext(context)

	blogs, err := handler.service.GetBlogs(request)
	if err != nil {
		log.Printf("Error getting blogs: %v\n", err)
	}

	isError := err != nil
	component := pages.Blogs(isAdmin, isError, blogs, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering blogs: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering the page"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

func (handler *BlogHandler) GetBlogs(context *gin.Context) {
	request := context.Request.Context()
	blogs, err := handler.service.GetBlogs(request)
	if err != nil {
		log.Printf("error getting blogs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting blogs"})
		return
	}

	context.JSON(http.StatusOK, blogs)
}

func (handler *BlogHandler) GetBlogByID(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")
	blog, err := handler.service.GetBlogsByID(request, id)
	if err != nil {
		log.Printf("error getting blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting blog"})
		return
	}

	context.JSON(http.StatusOK, blog)
}

func (handler *BlogHandler) CreateBlog(context *gin.Context) {
	request := context.Request.Context()
	var blog types.Blog

	if err := context.ShouldBindJSON(&blog); err != nil {
		log.Printf("error while creating blog: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating blog"})
		return
	}

	dbBlog := lib.DeserializeBlog(blog)
	result, err := handler.service.CreateBlog(request, &dbBlog)
	if err != nil {
		log.Printf("error while creating blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating blog"})
		return
	}

	response := gin.H{"blog": blog, "presignedUrl": result.PresignedUrl, "imageUrl": result.ImageUrl}

	context.JSON(http.StatusCreated, response)
}

func (handler *BlogHandler) UpdateBlog(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var blog types.Blog
	if err := context.ShouldBindJSON(&blog); err != nil {
		log.Printf("error while updating blog: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating blog"})
		return
	}

	dbBlog := lib.DeserializeBlog(blog)
	result, err := handler.service.UpdateBlog(request, id, &dbBlog)
	if err != nil {
		log.Printf("error while updating blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating blog"})
		return
	}

	response := gin.H{"blog": blog, "presignedUrl": result.PresignedUrl, "imageUrl": result.ImageUrl}

	context.JSON(http.StatusAccepted, response)
}

func (handler *BlogHandler) DeleteBlog(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteBlog(request, id); err != nil {
		log.Printf("error while deleting blog: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting blog"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "blog successfully deleted", "id": id})
}
