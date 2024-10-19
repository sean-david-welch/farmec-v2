package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"google.golang.org/api/youtube/v3"
	"regexp"
)

type VideoService interface {
	TransformData(video *db.Video) (*db.Video, error)
	GetVideos(ctx context.Context, id string) ([]types.Video, error)
	CreateVideo(ctx context.Context, video *db.Video) error
	UpdateVideo(ctx context.Context, id string, video *db.Video) error
	DeleteVideo(ctx context.Context, id string) error
}

type VideoServiceImpl struct {
	repo    repository.VideoRepo
	service *youtube.Service
}

func NewVideoService(repo repository.VideoRepo, service *youtube.Service) *VideoServiceImpl {
	return &VideoServiceImpl{
		repo:    repo,
		service: service,
	}
}

func extractVideoID(urlStr string) (string, error) {
	re := regexp.MustCompile(`(?:v=|/)([0-9A-Za-z_-]{11}).*`)
	matches := re.FindStringSubmatch(urlStr)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("invalid YouTube URL")
}

func (service *VideoServiceImpl) TransformData(ctx context.Context, video *db.Video) (*db.Video, error) {
	if !video.WebUrl.Valid {
		return nil, fmt.Errorf("web_url is invalid")
	}

	videoId, err := extractVideoID(video.WebUrl.String)
	if err != nil {
		return nil, err
	}

	call := service.service.Videos.List([]string{"id", "snippet"}).Id(videoId)
	response, err := call.Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error calling YouTube API: %w", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("video not found on YouTube")
	}

	item := response.Items[0]
	if item.Snippet == nil {
		return nil, fmt.Errorf("video snippet is missing in YouTube API response")
	}
	if item.Snippet.Thumbnails == nil || item.Snippet.Thumbnails.Medium == nil {
		return nil, fmt.Errorf("video thumbnail is missing in YouTube API response")
	}

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

func (service *VideoServiceImpl) GetVideos(ctx context.Context, id string) ([]types.Video, error) {
	videos, err := service.repo.GetVideos(ctx, id)
	if err != nil {
		return nil, err
	}
	var result []types.Video
	for _, video := range videos {
		result = append(result, lib.SerializeVideo(video))
	}
	return result, nil
}

func (service *VideoServiceImpl) CreateVideo(ctx context.Context, video *db.Video) error {
	videoData, err := service.TransformData(ctx, video)
	if err != nil {
		return err
	}
	return service.repo.CreateVideo(ctx, videoData)
}

func (service *VideoServiceImpl) UpdateVideo(ctx context.Context, id string, video *db.Video) error {
	videoData, err := service.TransformData(ctx, video)
	if err != nil {
		return err
	}

	err = service.repo.UpdateVideo(ctx, id, videoData)
	if err != nil {
		return err
	}

	return nil
}

func (service *VideoServiceImpl) DeleteVideo(ctx context.Context, id string) error {
	err := service.repo.DeleteVideo(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
