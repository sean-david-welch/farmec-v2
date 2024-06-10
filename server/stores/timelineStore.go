package stores

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineStore interface {
	GetTimelines() ([]types.Timeline, error)
	CreateTimeline(timeline *types.Timeline) error
	UpdateTimeline(id string, timeline *types.Timeline) error
	DeleteTimeline(id string) error
}

type TimelineStoreImpl struct {
	database *sql.DB
}

func NewTimelineStore(database *sql.DB) *TimelineStoreImpl {
	return &TimelineStoreImpl{database: database}
}

func (store *TimelineStoreImpl) GetTimelines() ([]types.Timeline, error) {
	var timelines []types.Timeline

	query := `SELECT * FROM "Timeline"`
	rows, err := store.database.Query(query)
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

func (store *TimelineStoreImpl) CreateTimeline(timeline *types.Timeline) error {
	timeline.ID = uuid.NewString()
	timeline.Created = time.Now().String()

	query := `INSERT INTO "Timeline" (id, title, date, body, created) VALUES ($1, $2, $3, $4, $5)`

	_, err := store.database.Exec(query, timeline.ID, timeline.Title, timeline.Date, timeline.Body, timeline.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating timeline: %w", err)
	}

	return nil
}

func (store *TimelineStoreImpl) UpdateTimeline(id string, timeline *types.Timeline) error {
	query := `UPDATE "Timeline" SET title = $1, data = %2, body = $3 WHERE "id" = $4`

	_, err := store.database.Exec(query, timeline.Title, timeline.Date, timeline.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating timeline: %w", err)
	}

	return nil
}

func (store *TimelineStoreImpl) DeleteTimeline(id string) error {
	query := `DELETE FROM "Timeline" WHERE "id" = $1`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting timeline: %w", err)
	}

	return nil
}
