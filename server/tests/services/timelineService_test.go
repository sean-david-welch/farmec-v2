package services_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TimelineTestService(test *testing.T) (*mocks.MockTimelineRepository, services.TimelineService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockTimelineRepository(controller)
	service := services.NewTimelineService(mockRepo)

	return mockRepo, service
}

func TestGetTimeline(test *testing.T) {
	mockRepo, service := TimelineTestService(test)

	expectedTimeline := []types.Timeline{
		{
			ID:      "1",
			Title:   "title 1",
			Date:    "01/01/24",
			Body:    "body 1",
			Created: time.Now(),
		},
		{

			ID:      "2",
			Title:   "title 2",
			Date:    "01/01/24",
			Body:    "body 2",
			Created: time.Now(),
		},
	}

	mockRepo.EXPECT().GetTimelines().Return(expectedTimeline, nil)

	timelines, err := service.GetTimelines()

	assert.NoError(test, err)
	assert.Equal(test, timelines, expectedTimeline)
}

func TestCreateTimeline(test *testing.T) {
	mockRepo, service := TimelineTestService(test)

	newTimeline := &types.Timeline{
		ID:      "1",
		Title:   "title 1",
		Date:    "01/01/24",
		Body:    "body 1",
		Created: time.Now(),
	}

	mockRepo.EXPECT().CreateTimeline(newTimeline).Return(nil)

	err := service.CreateTimeline(newTimeline)

	assert.NoError(test, err)
}
