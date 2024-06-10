package services

import (
	"fmt"
	"strings"

	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"google.golang.org/api/youtube/v3"
)

type VideoService interface {
	TransformData(video *types.Video) (*types.Video, error)
	GetVideos(id string) ([]types.Video, error)
	CreateVideo(video *types.Video) error
	UpdateVideo(id string, video *types.Video) error
	DeleteVideo(id string) error
}

type VideoServiceImpl struct {
	store          store.VideoStore
	youtubeService *youtube.Service
}

func NewVideoService(store store.VideoStore, youtubeService *youtube.Service) *VideoServiceImpl {
	return &VideoServiceImpl{
		store:          store,
		youtubeService: youtubeService,
	}
}

func (service *VideoServiceImpl) TransformData(video *types.Video) (*types.Video, error) {
	splits := strings.Split(video.WebURL, "v=")
	if len(splits) < 2 {
		return nil, fmt.Errorf("invalid web_url format")
	}

	videoIdSplits := strings.Split(splits[1], "&")
	if len(videoIdSplits) < 1 {
		return nil, fmt.Errorf("invalid web_url format")
	}

	videoId := videoIdSplits[0]

	call := service.youtubeService.Videos.List([]string{"id", "snippet"}).Id(videoId).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("error calling YouTube API: %w", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("video not found on YouTube")
	}

	item := response.Items[0]
	videoData := &types.Video{
		ID:           video.ID,
		SupplierID:   video.SupplierID,
		WebURL:       video.WebURL,
		Title:        &item.Snippet.Title,
		Description:  &item.Snippet.Description,
		VideoID:      &item.Id,
		ThumbnailURL: &item.Snippet.Thumbnails.Medium.Url,
		Created:      video.Created,
	}

	return videoData, nil
}

func (service *VideoServiceImpl) GetVideos(id string) ([]types.Video, error) {
	return service.store.GetVideos(id)
}

func (service *VideoServiceImpl) CreateVideo(video *types.Video) error {
	videoData, err := service.TransformData(video)
	if err != nil {
		return err
	}

	err = service.store.CreateVideo(videoData)
	if err != nil {
		return err
	}

	return nil
}

func (service *VideoServiceImpl) UpdateVideo(id string, video *types.Video) error {
	videoData, err := service.TransformData(video)
	if err != nil {
		return err
	}

	err = service.store.UpdateVideo(id, videoData)
	if err != nil {
		return err
	}

	return nil
}

func (service *VideoServiceImpl) DeleteVideo(id string) error {
	err := service.store.DeleteVideo(id)
	if err != nil {
		return err
	}

	return nil
}
