package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineService interface {
	GetTimelines(ctx context.Context) ([]types.Timeline, error)
	CreateTimeline(ctx context.Context, timeline *db.Timeline) error
	UpdateTimeline(ctx context.Context, id string, timeline *db.Timeline) error
	DeleteTimeline(ctx context.Context, id string) error
}

type TimelineServiceImpl struct {
	repo repository.TimelineRepo
}

func NewTimelineService(repo repository.TimelineRepo) *TimelineServiceImpl {
	return &TimelineServiceImpl{repo: repo}
}

func (service *TimelineServiceImpl) GetTimelines(ctx context.Context) ([]types.Timeline, error) {
	timelines, err := service.repo.GetTimelines(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.Timeline
	for _, event := range timelines {
		result = append(result, lib.SerializeTimeline(event))
	}
	return result, nil
}

func (service *TimelineServiceImpl) CreateTimeline(ctx context.Context, timeline *db.Timeline) error {
	if err := service.repo.CreateTimeline(ctx, timeline); err != nil {
		return err
	}
	return nil
}

func (service *TimelineServiceImpl) UpdateTimeline(ctx context.Context, id string, timeline *db.Timeline) error {
	if err := service.repo.UpdateTimeline(ctx, id, timeline); err != nil {
		return err
	}
	return nil
}

func (service *TimelineServiceImpl) DeleteTimeline(ctx context.Context, id string) error {
	if err := service.repo.DeleteTimeline(ctx, id); err != nil {
		return err
	}
	return nil
}
