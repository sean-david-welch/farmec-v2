package service

import (
	"testing"

	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

    mockRepo := new(MockLineItemRepository)

    mockRepo.On("GetLineItems").Return(expectedLineItems, nil)

    service := services.NewLineItemService(mockRepo)

    lineItems, err := service.GetLineItems()

    assert.NoError(t, err)
    assert.Equal(t, expectedLineItems, lineItems)

    mockRepo.AssertExpectations(t)
}
