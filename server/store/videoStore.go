package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type VideoStore interface {
	GetVideos(id string) ([]types.Video, error)
	CreateVideo(video *types.Video) error
	UpdateVideo(id string, video *types.Video) error
	DeleteVideo(id string) error
}

type VideoStoreImpl struct {
	database *sql.DB
}

func NewVideoStore(database *sql.DB) *VideoStoreImpl {
	return &VideoStoreImpl{database: database}
}

func (store *VideoStoreImpl) GetVideos(id string) ([]types.Video, error) {
	var videos []types.Video

	query := `SELECT * FROM "Video" WHERE "supplier_id" = ?`
	rows, err := store.database.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var video types.Video

		err := rows.Scan(&video.ID, &video.SupplierID, &video.WebURL, &video.Title, &video.Description, &video.VideoID, &video.ThumbnailURL, &video.Created)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return videos, nil
}

func (store *VideoStoreImpl) CreateVideo(video *types.Video) error {

	video.ID = uuid.NewString()
	video.Created = time.Now().String()

	query := `INSERT INTO "Video" ("id", "supplier_id", "web_url", "title", "description", "video_id", "thumbnail_url", "created") VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := store.database.Exec(query, video.ID, video.SupplierID, video.WebURL, video.Title, video.Description, video.VideoID, video.ThumbnailURL, video.Created)

	if err != nil {
		return fmt.Errorf("error creating video: %w", err)
	}

	return nil
}

func (store *VideoStoreImpl) UpdateVideo(id string, video *types.Video) error {
	query := `UPDATE "Video" SET "supplier_id" = ?, "web_url" = ?, "title" = ?, "description" = ?, "video_id" = ?, "thumbnail_url" = ? WHERE "id" = ?`

	_, err := store.database.Exec(query, video.SupplierID, video.WebURL, video.Title, video.Description, video.VideoID, video.ThumbnailURL, id)

	if err != nil {
		return fmt.Errorf("error updating video: %w", err)
	}

	return nil
}

func (store *VideoStoreImpl) DeleteVideo(id string) error {
	query := `DELETE FROM "Video" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting video: %w", err)
	}

	return nil
}
