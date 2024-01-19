package service

import (
	"testing"

	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockS3Client struct {
    mock.Mock
}

func (m *MockS3Client) GeneratePresignedUrl(folder, image string) (string, string, error) {
    args := m.Called(folder, image)
    return args.String(0), args.String(1), args.Error(2)
}

func (m *MockS3Client) DeleteImageFromS3(imageUrl string) error {
    args := m.Called(imageUrl)
    return args.Error(0)
}


type MockLineItemRepository struct {
    mock.Mock
}

func (mock *MockLineItemRepository) GetLineItems() ([]types.LineItem, error) {
    args := mock.Called()
    return args.Get(0).([]types.LineItem), args.Error(1)
}

func (mock *MockLineItemRepository) GetLineItemById(id string) (*types.LineItem, error) {
    args := mock.Called(id)
    return args.Get(0).(*types.LineItem), args.Error(1)
}

func (mock *MockLineItemRepository) CreateLineItem(lineItem *types.LineItem) error {
    args := mock.Called(lineItem)
    return args.Error(0)
}

func (mock *MockLineItemRepository) UpdateLineItem(id string, lineItem *types.LineItem) error {
    args := mock.Called(id, lineItem)
    return args.Error(0)
}

func (mock *MockLineItemRepository) DeleteLineItem(id string) error {
    args := mock.Called(id)
    return args.Error(0)
}

func TestGetLineItems(t *testing.T) {
    expectedLineItems := []types.LineItem{
        {ID: "1", Name: "Item 1", Price: 10.99, Image: "image1.jpg"},
        {ID: "2", Name: "Item 2", Price: 20.99, Image: "image2.jpg"},
    }

    folder := "test-folder"
    mockS3Client := new(MockS3Client)
    mockRepo := new(MockLineItemRepository)

    mockRepo.On("GetLineItems").Return(expectedLineItems, nil)
    
    service := services.NewLineItemService(mockRepo, mockS3Client, folder)    

    lineItems, err := service.GetLineItems()

    assert.NoError(t, err)
    assert.Equal(t, expectedLineItems, lineItems)

    mockRepo.AssertExpectations(t)
}

func TestCreateLineItem(t *testing.T) {
    // Create a mock repository
    folder := "test-folder"
    mockS3Client := new(MockS3Client)
    mockRepo := new(MockLineItemRepository)

    service := services.NewLineItemService(mockRepo, mockS3Client, folder)    

    mockS3Client.On("GeneratePresignedUrl", mock.Anything, mock.Anything).Return("presigned-url", "image-url", nil)

    // Define a line item to be created
    newLineItem := &types.LineItem{
        ID: "3", 
        Name: "Item 3", 
        Price: 15.99, 
        Image: "image3.jpg",
    }

    // Set up expectations
    mockRepo.On("CreateLineItem", newLineItem).Return(nil)

    // Call the method
    result, err := service.CreateLineItem(newLineItem)

    // Assert that no error occurred
    assert.NoError(t, err)

    assert.NotNil(t, result)
    assert.Equal(t, "image-url", result.ImageUrl)
    assert.Equal(t, "presigned-url", result.PresginedUrl)

    // Assert that the expectations were met
    mockRepo.AssertExpectations(t)
}

