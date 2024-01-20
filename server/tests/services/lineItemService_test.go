package services_test

import (
	"testing"

	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestEnvironment() (*mocks.MockLineItemRepository, *mocks.MockS3Client, services.LineItemService) {
	folder := "test-folder"
    
	mockS3Client := new(mocks.MockS3Client)
	mockRepo := new(mocks.MockLineItemRepository)
	service := services.NewLineItemService(mockRepo, mockS3Client, folder)

	return mockRepo, mockS3Client, service
}

func TestGetLineItems(t *testing.T) {
    expectedLineItems := []types.LineItem{
        {ID: "1", Name: "Item 1", Price: 10.99, Image: "image1.jpg"},
        {ID: "2", Name: "Item 2", Price: 20.99, Image: "image2.jpg"},
    }
    
    mockRepo, _, service := setupTestEnvironment()

    mockRepo.On("GetLineItems").Return(expectedLineItems, nil)

    lineItems, err := service.GetLineItems()

    assert.NoError(t, err)
    assert.Equal(t, expectedLineItems, lineItems)

    mockRepo.AssertExpectations(t)
}

func TestCreateLineItem(t *testing.T) {

    mockRepo, mockS3Client, service := setupTestEnvironment()

    mockS3Client.On("GeneratePresignedUrl", mock.Anything, mock.Anything).Return("presigned-url", "image-url", nil)

    newLineItem := &types.LineItem{
        ID: "3", 
        Name: "Item 3", 
        Price: 15.99, 
        Image: "image3.jpg",
    }

    mockRepo.On("CreateLineItem", newLineItem).Return(nil)

    result, err := service.CreateLineItem(newLineItem)

    assert.NoError(t, err)

    assert.NotNil(t, result)
    assert.Equal(t, "image-url", result.ImageUrl)
    assert.Equal(t, "presigned-url", result.PresginedUrl)

    mockRepo.AssertExpectations(t)
}

