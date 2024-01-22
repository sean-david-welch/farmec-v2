package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func LineItemTestService(t *testing.T) (*mocks.MockLineItemRepository, *mocks.MockS3Client, services.LineItemService) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockLineItemRepository(ctrl)
	mockS3Client := mocks.NewMockS3Client(ctrl)
	service := services.NewLineItemService(mockRepo, mockS3Client, "test-folder")

	return mockRepo, mockS3Client, service
}

func TestGetLineItems(t *testing.T) {
	mockRepo, _, service := LineItemTestService(t)

	expectedLineItems := []types.LineItem{
		{ID: "1", Name: "Item 1", Price: 10.99, Image: "image1.jpg"},
		{ID: "2", Name: "Item 2", Price: 20.99, Image: "image2.jpg"},
	}

	mockRepo.EXPECT().GetLineItems().Return(expectedLineItems, nil)

	lineItems, err := service.GetLineItems()

	assert.NoError(t, err)
	assert.Equal(t, expectedLineItems, lineItems)
}

func TestCreateLineItem(t *testing.T) {
	mockRepo, mockS3Client, service := LineItemTestService(t)

	newLineItem := &types.LineItem{
		ID:    "3",
		Name:  "Item 3",
		Price: 15.99,
		Image: "image3.jpg",
	}

	mockS3Client.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presigned-url", "image-url", nil)
	mockRepo.EXPECT().CreateLineItem(newLineItem).Return(nil)

	result, err := service.CreateLineItem(newLineItem)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "image-url", result.ImageUrl)
	assert.Equal(t, "presigned-url", result.PresginedUrl)
}
