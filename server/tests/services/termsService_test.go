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

func TermsTestService(test *testing.T) (*mocks.MockTermsRepository, services.TermsService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockTermsRepository(controller)
	service := services.NewTermsService(mockRepo)

	return mockRepo, service
}

func TestGetTerms(test *testing.T) {
	mockRepo, service := TermsTestService(test)

	expectedTerms := []types.Terms{
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

	mockRepo.EXPECT().GetTerms().Return(expectedTerms, nil)

	terms, err := service.GetTerms()

	assert.NoError(test, err)
	assert.Equal(test, terms, expectedTerms)
}

func TestCreateTerms(test *testing.T) {
	mockRepo, service := TermsTestService(test)

	newTerms := &types.Terms{
		ID:      "1",
		Title:   "title 1",
		Body:    "body 1",
		Created: time.Now(),
	}

	mockRepo.EXPECT().CreateTerm(newTerms).Return(nil)

	err := service.CreateTerm(newTerms)

	assert.NoError(test, err)
}
