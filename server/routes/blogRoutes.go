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

func InitBlogs(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) {
	repo := repository.NewBlogRepo(database)
	service := services.NewBlogService(repo, s3Client, "Blogs")
	handler := handlers.NewBlogHandler(service, authMiddleware, supplierCache)

	BlogRoutes(router, handler, authMiddleware)
}

func BlogRoutes(router *gin.Engine, handler *handlers.BlogHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/blogs", authMiddleware.ViewMiddleware(), handler.BlogsView)

	blogGroup := router.Group("/api/blogs")
	blogGroup.GET("", handler.GetBlogs)
	blogGroup.GET("/:id", handler.GetBlogByID)

	protected := blogGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateBlog)
		protected.PUT("/:id", handler.UpdateBlog)
		protected.DELETE("/:id", handler.DeleteBlog)
	}
}
