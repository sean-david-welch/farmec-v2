package services

import (
	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineService interface {
	GetTimelines() ([]types.Timeline, error)
	CreateTimeline(timeline *types.Timeline) error
	UpdateTimeline(id string, timeline *types.Timeline) error
	DeleteTimeline(id string) error
}

type TimelineServiceImpl struct {
	store stores.TimelineStore
}

func NewTimelineService(store stores.TimelineStore) *TimelineServiceImpl {
	return &TimelineServiceImpl{store: store}
}

func (service *TimelineServiceImpl) GetTimelines() ([]types.Timeline, error) {
	timelines, err := service.store.GetTimelines()
	if err != nil {
		return nil, err
	}

	return timelines, nil
}

func (service *TimelineServiceImpl) CreateTimeline(timeline *types.Timeline) error {
	if err := service.store.CreateTimeline(timeline); err != nil {
		return err
	}

	return nil
}

func (service *TimelineServiceImpl) UpdateTimeline(id string, timeline *types.Timeline) error {
	if err := service.store.UpdateTimeline(id, timeline); err != nil {
		return err
	}

	return nil
}

func (service *TimelineServiceImpl) DeleteTimeline(id string) error {
	if err := service.store.DeleteTimeline(id); err != nil {
		return err
	}

	return nil
}
