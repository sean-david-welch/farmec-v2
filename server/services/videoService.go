package services

import (
	"fmt"
	"strings"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"google.golang.org/api/youtube/v3"
)

type VideoService struct {
	repository *repository.VideoRepositoy
	youtubeService *youtube.Service
}

func NewVideoService(repository *repository.VideoRepositoy, youtubeService *youtube.Service) *VideoService {
	return &VideoService{
		repository: repository,
		youtubeService: youtubeService,
	}
}

func (service *VideoService) TransformData(video *types.Video) (*types.Video, error) {
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
        ID: video.ID,
        SupplierID: video.SupplierID,
        WebURL: video.WebURL,
        Title: &item.Snippet.Title,
        Description: &item.Snippet.Description,
        VideoID: &item.Id,
        ThumbnailURL: &item.Snippet.Thumbnails.Medium.Url,
        Created: video.Created,
    }


	return videoData, nil
}

func (service *VideoService) GetVideos(id string) ([]types.Video, error) {
	return service.repository.GetVideos(id)
}

func (service *VideoService) CreateVideo(video *types.Video) error {
	videoData, err := service.TransformData(video); if err != nil {
		return err
	}

	service.repository.CreateVideo(videoData); if err != nil {
		return err
	}

	return nil
}

func (service *VideoService) UpdateVideo(id string, video *types.Video) error {
	videoData, err := service.TransformData(video); if err != nil {
		return err
	}

	service.repository.UpdateVideo(id, videoData); if err != nil {
		return err
	}

	return nil
}

func (service *VideoService) DeleteVideo(id string) error {
	err := service.repository.DeleteVideo(id); if err != nil {
		return err
	}

	return nil
}
