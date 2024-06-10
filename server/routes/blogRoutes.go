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

func InitBlogs(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	blogRepository := repository.NewBlogRepository(database)
	service := services.NewBlogService(blogRepository, s3Client, "Blogs")
	controller := handlers.NewBlogController(service)

	BlogRoutes(router, controller, adminMiddleware)
}

func BlogRoutes(router *gin.Engine, controller *handlers.BlogController, adminMiddleware *middleware.AdminMiddleware) {
	blogGroup := router.Group("/api/blogs")

	blogGroup.GET("", controller.GetBlogs)
	blogGroup.GET("/:id", controller.GetBlogByID)

	protected := blogGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", controller.CreateBlog)
		protected.PUT("/:id", controller.UpdateBlog)
		protected.DELETE("/:id", controller.DeleteBlog)
	}
}
