package controllers_test

import (
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

func (m *MockLineItemService) CreateLineItem(lineItem *types.LineItem) error {
	return nil
}

func (m *MockLineItemService) UpdateLineItem(id string, lineItem *types.LineItem) error {
	return nil
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