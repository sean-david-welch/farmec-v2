package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/stores"
)

type TimelineService interface {
	GetTimelines(ctx context.Context) ([]db.Timeline, error)
	CreateTimeline(ctx context.Context, timeline *db.Timeline) error
	UpdateTimeline(ctx context.Context, id string, timeline *db.Timeline) error
	DeleteTimeline(ctx context.Context, id string) error
}

type TimelineServiceImpl struct {
	store stores.TimelineStore
}

func NewTimelineService(store stores.TimelineStore) *TimelineServiceImpl {
	return &TimelineServiceImpl{store: store}
}

func (service *TimelineServiceImpl) GetTimelines(ctx context.Context) ([]db.Timeline, error) {
	timelines, err := service.store.GetTimelines(ctx)
	if err != nil {
		return nil, err
	}
	return timelines, nil
}

func (service *TimelineServiceImpl) CreateTimeline(ctx context.Context, timeline *db.Timeline) error {
	if err := service.store.CreateTimeline(ctx, timeline); err != nil {
		return err
	}
	return nil
}

func (service *TimelineServiceImpl) UpdateTimeline(ctx context.Context, id string, timeline *db.Timeline) error {
	if err := service.store.UpdateTimeline(ctx, id, timeline); err != nil {
		return err
	}
	return nil
}

func (service *TimelineServiceImpl) DeleteTimeline(ctx context.Context, id string) error {
	if err := service.store.DeleteTimeline(ctx, id); err != nil {
		return err
	}
	return nil
}
