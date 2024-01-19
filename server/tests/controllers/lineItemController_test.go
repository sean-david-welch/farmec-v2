package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestGetLineItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockLineItemService)
	controller := controllers.NewLineItemController(mockService)

	expectedLineItems := []types.LineItem{
	{
        ID: "1",
        Name: "Apple",
        Price: 1.99,
        Image: "apple.jpg",
    },
    {
        ID: "2",
        Name: "Banana",
        Price: 0.99,
        Image: "banana.jpg",
    }}
	mockService.On("GetLineItems").Return(expectedLineItems, nil)

	req, _ := http.NewRequest("GET", "/lineitems", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/lineitems", controller.GetLineItems)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var receivedLineItems []types.LineItem
    err := json.Unmarshal(w.Body.Bytes(), &receivedLineItems)
    assert.NoError(t, err)
    assert.Equal(t, expectedLineItems, receivedLineItems)

	mockService.AssertExpectations(t)
}

func TestCreateLineItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockLineItemService)
	controller := controllers.NewLineItemController(mockService)

	// Create a sample line item to send in the request
	newLineItem := &types.LineItem{
		ID:    "3",
		Name:  "Grape",
		Price: 2.50,
		Image: "grape.jpg",
	}
	newLineItemJson, _ := json.Marshal(newLineItem)

	// Define the expected result
	expectedResult := &types.ModelResult{
		PresginedUrl: "presigned-url",
		ImageUrl:     "image-url",
	}
	mockService.On("CreateLineItem", newLineItem).Return(expectedResult, nil)

	// Create an HTTP request and recorder
	req, _ := http.NewRequest("POST", "/lineitems", bytes.NewBuffer(newLineItemJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Set up the Gin context
	r := gin.Default()
	r.POST("/lineitems", controller.CreateLineItem)
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	// Asserting the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newLineItem.Name, response["lineItem"].(map[string]interface{})["name"])
	assert.Equal(t, expectedResult.PresginedUrl, response["presginedUrl"])
	assert.Equal(t, expectedResult.ImageUrl, response["imageUrl"])

	mockService.AssertExpectations(t)
}
