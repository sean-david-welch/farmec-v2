package routes

import (
	"context"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func InitializeVideos(router *gin.Engine, db *sql.DB, secrets *config.Secrets, firebase *lib.Firebase) {
	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(secrets.YoutubeApiKey)); if err != nil {
		log.Fatal("error calling YouTube API: ", err)
	}

	videoRepository := repository.NewVideoRepository(db)
	videoService := services.NewVideoService(videoRepository, youtubeService)
	videoController := controllers.NewVideoController(videoService)

	adminMiddleware := middleware.NewAdminMiddleware(firebase)

	VideoRoutes(router, videoController, adminMiddleware)
}

func VideoRoutes(router *gin.Engine, videoController *controllers.VideoController, adminMiddleware *middleware.AdminMiddleware) {
	videoGroup := router.Group("/api/videos")

	videoGroup.GET("/:id", videoController.GetVideos)

	protected := videoGroup.Group("")
	protected.Use(adminMiddleware.Middleware()); {
		protected.POST("", videoController.CreateVideo)
		protected.PUT("/:id", videoController.UpdateVideo)
		protected.DELETE("/:id", videoController.DeleteVideo)
	}
}