package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeBlogs(router *gin.Engine, database *sql.DB, s3Client *utils.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewBlogRepository(database)
	service := services.NewBlogService(repository, s3Client, "Blogs")
	controller := controllers.NewBlogController(service)

	BlogRoutes(router, controller, adminMiddleware)
}

func BlogRoutes(router *gin.Engine, controller *controllers.BlogController, adminMiddleware *middleware.AdminMiddleware) {
	blogGroup := router.Group("/api/blogs")

	blogGroup.GET("", controller.GetBlogs)
	blogGroup.GET("/:id", controller.GetBlogByID)

	protected := blogGroup.Group("").Use(adminMiddleware.Middleware()); {
		protected.POST("", controller.CreateBlog)
		protected.PUT("/:id", controller.UpdateBlog)
		protected.DELETE("/:id", controller.DeleteBlog)
	}
}