package mocks

import (
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/mock"
)

// Repository
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

// Service
type MockLineItemService struct {
    mock.Mock
}

func (m *MockLineItemService) GetLineItems() ([]types.LineItem, error) {
    args := m.Called()
    return args.Get(0).([]types.LineItem), args.Error(1)
}

func (m *MockLineItemService) GetLineItemById(id string) (*types.LineItem, error) {
	return nil, nil
}

func (m *MockLineItemService) CreateLineItem(lineItem *types.LineItem) (*types.ModelResult, error) {
	args := m.Called(lineItem)
	return args.Get(0).(*types.ModelResult), args.Error(1)
}


func (m *MockLineItemService) UpdateLineItem(id string, lineItem *types.LineItem) (*types.ModelResult, error) {
	return nil, nil
}

func (m *MockLineItemService) DeleteLineItem(id string) error {
	return nil
}