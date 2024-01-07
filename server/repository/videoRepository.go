package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

type VideoRepositoy struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepositoy {
	return &VideoRepositoy{db: db}
}

func (repository *VideoRepositoy) GetVideos(id string) ([]models.Video, error) {
	var videos []models.Video

	query := `SELECT * FROM "Video" WHERE "supplierId" = $1`
	rows, err := repository.db.Query(query, id); if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var video models.Video

		err := rows.Scan(&video.ID, &video.SupplierID, &video.WebURL, &video.Title, &video.Description, &video.VideoID, &video.ThumbnailURL)
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

func (repository *VideoRepositoy) CreateVideo(video *models.Video) error {

	video.ID = uuid.NewString()
	video.Created = time.Now()

    query := `INSERT INTO "Video" ("id", "supplierId", "web_url", "title", "description", "video_id", "thumbnail_url", "created") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := repository.db.Exec(query, video.ID, video.SupplierID, video.WebURL, video.Title, video.Description, video.VideoID, video.ThumbnailURL, video.Created)

	if err != nil {
		return fmt.Errorf("error creating video: %w", err)
	}

	return nil
}

func (repository *VideoRepositoy) UpdateVideo(id string, video *models.Video) error {
    query := `UPDATE "Video" SET "supplierId" = $1, "web_url" = $2, "title" = $3, "description" = $4, "video_id" = $5, "thumbnail_url" = $6 WHERE "id" = $7`

    _, err := repository.db.Exec(query, video.SupplierID, video.WebURL, video.Title, video.Description, video.VideoID, video.ThumbnailURL, id)

	if err != nil {
		return fmt.Errorf("error updating video: %w", err)
	}

	return nil
}

func (repository *VideoRepositoy) DeleteVideo(id string) error {
	query := `DELETE FROM "Video" WHERE "id" = $1`

	_, err := repository.db.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting video: %w", err)
	}
	
	return nil
}