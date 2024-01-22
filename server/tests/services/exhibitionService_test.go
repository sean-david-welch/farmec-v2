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

func ExhibitionTestService(test *testing.T) (*mocks.MockExhibitionRepository, services.ExhibitionService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockExhibitionRepository(controller)
	service := services.NewExhibitionService(mockRepo)

	return mockRepo, service
}

func TestGetExhibitions(test *testing.T) {
	mockRepo, service := ExhibitionTestService(test)

	expectedExhibitions := []types.Exhibition{
		{
			ID:       "1",
			Title:    "Title 1",
			Date:     "01/01/24",
			Location: "Location 1",
			Info:     "Info 1",
			Created:  time.Now(),
		},
		{
			ID:       "1",
			Title:    "Title 1",
			Date:     "01/01/24",
			Location: "Location 1",
			Info:     "Info 1",
			Created:  time.Now(),
		},
	}

	mockRepo.EXPECT().GetExhibitions().Return(expectedExhibitions, nil)

	exhibitions, err := service.GetExhibitions()

	assert.NoError(test, err)
	assert.Equal(test, exhibitions, expectedExhibitions)
}

func TestCreateExhibitions(test *testing.T) {
	mockRepo, service := ExhibitionTestService(test)

	newExhibition := &types.Exhibition{
		ID:       "3",
		Title:    "Title 3",
		Date:     "01/01/24",
		Location: "Location 3",
		Info:     "Info 3",
		Created:  time.Now(),
	}

	mockRepo.EXPECT().CreateExhibition(newExhibition)

	err := service.CreateExhibition(newExhibition)

	assert.NoError(test, err)
}
