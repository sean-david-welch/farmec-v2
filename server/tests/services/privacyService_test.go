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

func PrivacyTestService(test *testing.T) (*mocks.MockPrivacyRepository, services.PrivacyService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockPrivacyRepository(controller)
	service := services.NewPrivacyService(mockRepo)

	return mockRepo, service
}

func TestGetPrivacy(test *testing.T) {
	mockRepo, service := PrivacyTestService(test)

	expectedPrivacy := []types.Privacy{
		{
			ID:      "1",
			Title:   "title 1",
			Body:    "body 1",
			Created: time.Now(),
		},
		{

			ID:      "2",
			Title:   "title 2",
			Body:    "body 2",
			Created: time.Now(),
		},
	}

	mockRepo.EXPECT().GetPrivacy().Return(expectedPrivacy, nil)

	privacys, err := service.GetPrivacys()

	assert.NoError(test, err)
	assert.Equal(test, privacys, expectedPrivacy)
}

func TestCreatePrivacy(test *testing.T) {
	mockRepo, service := PrivacyTestService(test)

	newPrivacy := &types.Privacy{
		ID:      "1",
		Title:   "title 1",
		Body:    "body 1",
		Created: time.Now(),
	}

	mockRepo.EXPECT().CreatePrivacy(newPrivacy).Return(nil)

	err := service.CreatePrivacy(newPrivacy)

	assert.NoError(test, err)
}
