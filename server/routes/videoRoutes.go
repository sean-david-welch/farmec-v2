package routes

import (
	"context"
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func InitVideos(router *gin.Engine, database *sql.DB, secrets *lib.Secrets, adminMiddleware *middleware.AdminMiddleware) {
	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(secrets.YoutubeApiKey))
	if err != nil {
		log.Fatal("error calling YouTube API: ", err)
	}

	videoStore := repository.NewVideoStore(database)
	videoService := services.NewVideoService(videoStore, youtubeService)
	videoHandler := handlers.NewVideoHandler(videoService)

	VideoRoutes(router, videoHandler, adminMiddleware)
}

func VideoRoutes(router *gin.Engine, videoHandler *handlers.VideoHandler, adminMiddleware *middleware.AdminMiddleware) {
	videoGroup := router.Group("/api/videos")

	videoGroup.GET("/:id", videoHandler.GetVideos)

	protected := videoGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", videoHandler.CreateVideo)
		protected.PUT("/:id", videoHandler.UpdateVideo)
		protected.DELETE("/:id", videoHandler.DeleteVideo)
	}
}
