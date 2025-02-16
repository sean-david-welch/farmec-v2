package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitBlogs(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) {
	repo := repository.NewBlogRepo(database)
	service := services.NewBlogService(repo, s3Client, "Blogs")
	handler := handlers.NewBlogHandler(service, adminMiddleware, supplierCache)

	BlogRoutes(router, handler, adminMiddleware)
}

func BlogRoutes(router *gin.Engine, handler *handlers.BlogHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/blogs", adminMiddleware.ViewMiddleware(), handler.BlogsView)

	blogGroup := router.Group("/api/blogs")
	blogGroup.GET("", handler.GetBlogs)
	blogGroup.GET("/:id", handler.GetBlogByID)

	protected := blogGroup.Group("").Use(adminMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateBlog)
		protected.PUT("/:id", handler.UpdateBlog)
		protected.DELETE("/:id", handler.DeleteBlog)
	}
}
