package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"strings"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"google.golang.org/api/youtube/v3"
)

type VideoService interface {
	TransformData(video *db.Video) (*db.Video, error)
	GetVideos(ctx context.Context, id string) ([]db.Video, error)
	CreateVideo(ctx context.Context, video *db.Video) error
	UpdateVideo(ctx context.Context, id string, video *db.Video) error
	DeleteVideo(ctx context.Context, id string) error
}

type VideoServiceImpl struct {
	store          stores.VideoStore
	youtubeService *youtube.Service
}

func NewVideoService(store stores.VideoStore, youtubeService *youtube.Service) *VideoServiceImpl {
	return &VideoServiceImpl{
		store:          store,
		youtubeService: youtubeService,
	}
}

func (service *VideoServiceImpl) TransformData(video *db.Video) (*db.Video, error) {
	splits := strings.Split(video.WebUrl.String, "v=")
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
	videoData := &db.Video{
		ID:         video.ID,
		SupplierID: video.SupplierID,
		WebUrl:     video.WebUrl,
		Title: sql.NullString{
			String: item.Snippet.Title,
			Valid:  true,
		},
		Description: sql.NullString{
			String: item.Snippet.Description,
			Valid:  true,
		},
		VideoID: sql.NullString{
			String: item.Id,
			Valid:  true,
		},
		ThumbnailUrl: sql.NullString{
			String: item.Snippet.Thumbnails.Medium.Url,
			Valid:  true,
		},
		Created: video.Created,
	}

	return videoData, nil
}

func (service *VideoServiceImpl) GetVideos(id string) ([]db.Video, error) {
	return service.store.GetVideos(id)
}

func (service *VideoServiceImpl) CreateVideo(video *db.Video) error {
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

func (service *VideoServiceImpl) UpdateVideo(id string, video *db.Video) error {
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
