package stores

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type TimelineStore interface {
	GetTimelines(ctx context.Context) ([]db.Timeline, error)
	CreateTimeline(ctx context.Context, timeline *db.Timeline) error
	UpdateTimeline(ctx context.Context, id string, timeline *db.Timeline) error
	DeleteTimeline(ctx context.Context, id string) error
}

type TimelineStoreImpl struct {
	database *sql.DB
}

func NewTimelineStore(database *sql.DB) *TimelineStoreImpl {
	return &TimelineStoreImpl{database: database}
}

func (store *TimelineStoreImpl) GetTimelines() ([]db.Timeline, error) {
	var timelines []db.Timeline

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
		var timeline db.Timeline

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

func (store *TimelineStoreImpl) CreateTimeline(timeline *db.Timeline) error {
	timeline.ID = uuid.NewString()
	timeline.Created = time.Now().String()

	query := `INSERT INTO "Timeline" (id, title, date, body, created) VALUES (?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, timeline.ID, timeline.Title, timeline.Date, timeline.Body, timeline.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating timeline: %w", err)
	}

	return nil
}

func (store *TimelineStoreImpl) UpdateTimeline(id string, timeline *db.Timeline) error {
	query := `UPDATE "Timeline" SET title = ?, date = ?, body = ? WHERE "id" = ?`

	_, err := store.database.Exec(query, timeline.Title, timeline.Date, timeline.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating timeline: %w", err)
	}

	return nil
}

func (store *TimelineStoreImpl) DeleteTimeline(id string) error {
	query := `DELETE FROM "Timeline" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting timeline: %w", err)
	}

	return nil
}
