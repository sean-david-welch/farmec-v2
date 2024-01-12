package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

// type Timeline struct {
//     ID      string `json:"id"`
//     Title   string `json:"title"`
//     Date    string `json:"date"`
//     Body    string `json:"body"`
//     Created string `json:"created"`
// }

type TimelineRepository struct {
	db *sql.DB
}

func NewTimelineRepository(db *sql.DB) *TimelineRepository {
	return &TimelineRepository{db: db}
}

func(repository *TimelineRepository) GetTimelines() ([]models.Timeline, error) {
	var timelines []models.Timeline

	query := `SELECT * FROM "Timeline"`
	rows, err := repository.db.Query(query); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var timeline models.Timeline

		if err := rows.Scan(&timeline.ID, &timeline.Title, &timeline.Date, &timeline.Body, &timeline.Created); err != nil {
			return nil, fmt.Errorf("error occurred while scanning rows: %w", err)
		}
		timelines = append(timelines, timeline)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
	}

	return timelines, nil
}

func(repository *TimelineRepository) CreateTimeline(timeline *models.Timeline) error {
	timeline.ID = uuid.NewString()
	timeline.Created = time.Now()

	query := `INSERT INTO "Timeline" (id, title, date, body, created) VALUES ($1, $2, $3, $4, $5)`

	_, err := repository.db.Exec(query, timeline.ID, timeline.Title, timeline.Date, timeline.Body, timeline.Created); if err != nil {
		return fmt.Errorf("error occurred while creating timeline: %w", err)
	}
	
	return nil
}

func(repository *TimelineRepository) UpdateTimeline(id string, timeline *models.Timeline) error {
	query := `UPDATE "Timeline" SET title = $1, data = %2, body = $3 WHERE "id" = $4`

	_, err := repository.db.Exec(query,  timeline.Title, timeline.Date, timeline.Body, id); if err != nil {
		return fmt.Errorf("error occurred while updating timeline: %w", err)
	}
	
	return nil
}

func(repository *TimelineRepository) DeleteTimeline(id string) error {
	query := `DELETE FROM "Timeline" WHERE "id" = $1`

	_, err := repository.db.Exec(query, id); if err != nil {
		return fmt.Errorf("error occurred while deleting timeline: %w", err)
	}

	return nil
}