package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineService struct {
	repository *repository.TimelineRepository
}

func NewTimelineService(repository *repository.TimelineRepository) *TimelineService {
	return &TimelineService{repository: repository}
}

func(service *TimelineService) GetTimelines() ([]types.Timeline, error) {
	timelines, err := service.repository.GetTimelines(); if err != nil {
		return nil, err
	}
	
	return timelines, nil
}

func(service *TimelineService) CreateTimeline(timeline *types.Timeline) error {
	if err := service.repository.CreateTimeline(timeline); err != nil {
		return err
	}
	
	return nil
}

func(service *TimelineService) UpdateTimeline(id string, timeline *types.Timeline) error {
	if err := service.repository.UpdateTimeline(id, timeline); err != nil {
		return err
	}

	return nil
}

func(service *TimelineService) DeleteTimeline(id string) error {
	if err := service.repository.DeleteTimeline(id); err != nil {
		return err
	}

	return nil
}