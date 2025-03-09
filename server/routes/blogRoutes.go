package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func BlogRoutes(router *gin.Engine, handler *handlers.BlogHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	router.GET("/blogs", authMiddleware.ViewMiddleware(), handler.BlogsView)
	router.GET("/blogs/:id", authMiddleware.ViewMiddleware(), handler.BlogDetailView)

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
