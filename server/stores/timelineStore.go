package stores

import (
	"context"
	"database/sql"
	"fmt"
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
	queries *db.Queries
}

func NewTimelineStore(sql *sql.DB) *TimelineStoreImpl {
	queries := db.New(sql)
	return &TimelineStoreImpl{queries: queries}
}

func (store *TimelineStoreImpl) GetTimelines(ctx context.Context) ([]db.Timeline, error) {
	timelines, err := store.queries.GetTimelines(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting timelines from the db: %w", err)
	}

	var result []db.Timeline
	for _, timeline := range timelines {
		result = append(result, db.Timeline{
			ID:      timeline.ID,
			Title:   timeline.Title,
			Date:    timeline.Date,
			Body:    timeline.Body,
			Created: timeline.Created,
		})
	}
	return result, nil
}

func (store *TimelineStoreImpl) CreateTimeline(ctx context.Context, timeline *db.Timeline) error {
	timeline.ID = uuid.NewString()
	timeline.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreateTimelineParams{
		ID:      timeline.ID,
		Title:   timeline.Title,
		Date:    timeline.Date,
		Body:    timeline.Body,
		Created: timeline.Created,
	}

	if err := store.queries.CreateTimeline(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating a timeline: %w", err)
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
