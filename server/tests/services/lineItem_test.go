package service_tests

import (
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/mock"
)


type MockLineItemRepository struct {
    mock.Mock
}

func (mock *MockLineItemRepository) GetLineItems() ([]types.LineItem, error) {
    args := mock.Called()
    return args.Get(0).([]types.LineItem), args.Error(1)
}

// func TestGetLineItems(t *testing.T) {
//     // Creating a mock repository
//     mockRepo := new(MockLineItemRepository)
//     lineItems := []types.LineItem{
//         {ID: "1", Name: "Item 1", Price: 10.99, Image: "image1.jpg"},
//         {ID: "2", Name: "Item 2", Price: 20.99, Image: "image2.jpg"},
//     }

//     // Setting expectations
//     mockRepo.On("GetLineItems").Return(lineItems, nil)

//     // Creating the service with the mock repository
//     service := services.NewLineItemService(mockRepo)

//     // Calling the service method
//     result, err := service.GetLineItems()

//     // Asserting expectations
//     assert.NoError(t, err)
//     assert.Equal(t, lineItems, result)

//     // Assert that the expectations were met
//     mockRepo.AssertExpectations(t)
// }
