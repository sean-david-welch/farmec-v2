package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type VideoStore interface {
	GetVideos(ctx context.Context, id string) ([]db.Video, error)
	CreateVideo(ctx context.Context, video *db.Video) error
	UpdateVideo(ctx context.Context, id string, video *db.Video) error
	DeleteVideo(ctx context.Context, id string) error
}

type VideoStoreImpl struct {
	queries *db.Queries
}

func NewVideoStore(sql *sql.DB) *VideoStoreImpl {
	queries := db.New(sql)
	return &VideoStoreImpl{queries: queries}
}

func (store *VideoStoreImpl) GetVideos(ctx context.Context, id string) ([]db.Video, error) {
	videos, err := store.queries.SelectVideos(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting videos from the database: %w", err)
	}
	var result []db.Video
	for _, video := range videos {
		result = append(result, db.Video{
			ID:           video.ID,
			SupplierID:   video.SupplierID,
			WebUrl:       video.WebUrl,
			Title:        video.Title,
			Description:  video.Description,
			VideoID:      video.VideoID,
			ThumbnailUrl: video.ThumbnailUrl,
			Created:      video.Created,
		})
	}
	return result, nil
}

func (store *VideoStoreImpl) CreateVideo(ctx context.Context, video *db.Video) error {
	video.ID = uuid.NewString()
	video.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}
	params := db.CreateVideoParams{
		ID:           video.ID,
		SupplierID:   video.SupplierID,
		WebUrl:       video.WebUrl,
		Title:        video.Title,
		Description:  video.Description,
		VideoID:      video.VideoID,
		ThumbnailUrl: video.ThumbnailUrl,
		Created:      video.Created,
	}
	if err := store.queries.CreateVideo(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating a video in the db: %w", err)
	}
	return nil
}

func (store *VideoStoreImpl) UpdateVideo(ctx context.Context, id string, video *db.Video) error {
	params := db.UpdateVideoParams{
		SupplierID:   video.SupplierID,
		WebUrl:       video.ThumbnailUrl,
		Title:        video.ThumbnailUrl,
		Description:  video.ThumbnailUrl,
		VideoID:      video.ThumbnailUrl,
		ThumbnailUrl: video.ThumbnailUrl,
		ID:           id,
	}
	if err := store.queries.UpdateVideo(ctx, params); err != nil {
		return fmt.Errorf("error occurred while updating a video: %w", err)
	}
	return nil
}

func (store *VideoStoreImpl) DeleteVideo(ctx context.Context, id string) error {
	if err := store.queries.DeleteMachine(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting video: %w", err)
	}
	return nil
}
