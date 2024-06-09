package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineRepository interface {
	GetTimelines() ([]types.Timeline, error)
	CreateTimeline(timeline *types.Timeline) error
	UpdateTimeline(id string, timeline *types.Timeline) error
	DeleteTimeline(id string) error
}

type TimelineRepositoryImpl struct {
	database *sql.DB
}

func NewTimelineRepository(database *sql.DB) *TimelineRepositoryImpl {
	return &TimelineRepositoryImpl{database: database}
}

func (repository *TimelineRepositoryImpl) GetTimelines() ([]types.Timeline, error) {
	var timelines []types.Timeline

	query := `SELECT * FROM "Timeline"`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var timeline types.Timeline

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

func (repository *TimelineRepositoryImpl) CreateTimeline(timeline *types.Timeline) error {
	timeline.ID = uuid.NewString()
	timeline.Created = time.Now().String()

	query := `INSERT INTO "Timeline" (id, title, date, body, created) VALUES ($1, $2, $3, $4, $5)`

	_, err := repository.database.Exec(query, timeline.ID, timeline.Title, timeline.Date, timeline.Body, timeline.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating timeline: %w", err)
	}

	return nil
}

func (repository *TimelineRepositoryImpl) UpdateTimeline(id string, timeline *types.Timeline) error {
	query := `UPDATE "Timeline" SET title = $1, data = %2, body = $3 WHERE "id" = $4`

	_, err := repository.database.Exec(query, timeline.Title, timeline.Date, timeline.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating timeline: %w", err)
	}

	return nil
}

func (repository *TimelineRepositoryImpl) DeleteTimeline(id string) error {
	query := `DELETE FROM "Timeline" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting timeline: %w", err)
	}

	return nil
}
