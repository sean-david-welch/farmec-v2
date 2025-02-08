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

func InitVideos(router *gin.Engine, database *sql.DB, secrets *lib.Secrets, adminMiddleware *middleware.AuthMiddlewareImpl) {
	yt, err := youtube.NewService(context.Background(), option.WithAPIKey(secrets.YoutubeApiKey))
	if err != nil {
		log.Fatal("error calling YouTube API: ", err)
	}

	repo := repository.NewVideoRepo(database)
	service := services.NewVideoService(repo, yt)
	handler := handlers.NewVideoHandler(service)

	VideoRoutes(router, handler, adminMiddleware)
}

func VideoRoutes(router *gin.Engine, handler *handlers.VideoHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	videoGroup := router.Group("/api/videos")

	videoGroup.GET("/:id", handler.GetVideos)

	protected := videoGroup.Group("").Use(adminMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateVideo)
		protected.PUT("/:id", handler.UpdateVideo)
		protected.DELETE("/:id", handler.DeleteVideo)
	}
}
