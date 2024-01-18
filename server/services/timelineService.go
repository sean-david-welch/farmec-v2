package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TimelineService interface {
	GetTimelines() ([]types.Timeline, error) 
	CreateTimeline(timeline *types.Timeline) error 
	UpdateTimeline(id string, timeline *types.Timeline) error 
	DeleteTimeline(id string) error 
}

type TimelineServiceImpl struct {
	repository repository.TimelineRepository
}

func NewTimelineService(repository repository.TimelineRepository) *TimelineServiceImpl {
	return &TimelineServiceImpl{repository: repository}
}

func(service *TimelineServiceImpl) GetTimelines() ([]types.Timeline, error) {
	timelines, err := service.repository.GetTimelines(); if err != nil {
		return nil, err
	}
	
	return timelines, nil
}

func(service *TimelineServiceImpl) CreateTimeline(timeline *types.Timeline) error {
	if err := service.repository.CreateTimeline(timeline); err != nil {
		return err
	}
	
	return nil
}

func(service *TimelineServiceImpl) UpdateTimeline(id string, timeline *types.Timeline) error {
	if err := service.repository.UpdateTimeline(id, timeline); err != nil {
		return err
	}

	return nil
}

func(service *TimelineServiceImpl) DeleteTimeline(id string) error {
	if err := service.repository.DeleteTimeline(id); err != nil {
		return err
	}

	return nil
}