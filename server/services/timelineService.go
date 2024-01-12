package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type TimelineService struct {
	repository *repository.TimelineRepository
}

func NewTimelineService(repository *repository.TimelineRepository) *TimelineService {
	return &TimelineService{repository: repository}
}

func(service *TimelineService) GetTimelines() ([]models.Timeline, error) {
	timelines, err := service.repository.GetTimelines(); if err != nil {
		return nil, err
	}
	
	return timelines, nil
}

func(service *TimelineService) CreateTimeline(timeline *models.Timeline) error {
	if err := service.repository.CreateTimeline(timeline); err != nil {
		return err
	}
	
	return nil
}

func(service *TimelineService) UpdateTimeline(id string, timeline *models.Timeline) error {
	if err := service.repository.UpdateTimeine(id, timeline); err != nil {
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