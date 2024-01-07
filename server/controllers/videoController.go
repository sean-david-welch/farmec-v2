package controllers

import "github.com/sean-david-welch/farmec-v2/server/services"

type VideoController struct {
	videoService *services.VideoService
}

func NewVideoController(videoService *services.VideoService) *VideoController{
	return &VideoController{videoService: videoService}
}

