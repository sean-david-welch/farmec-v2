package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func LineItemControllerTesr(t *testing.T) (*gin.Engine, *mocks.MockLineItemService, *controllers.LineItemController, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockLineItemService(ctrl)
	controller := controllers.NewLineItemController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	return router, mockService, controller, recorder
}

func TestGetLineItems(t *testing.T) {
	router, mockService, controller, recorder := LineItemControllerTesr(t)

	expectedLineItems := []types.LineItem{
		{ID: "1", Name: "Apple", Price: 1.99, Image: "apple.jpg"},
		{ID: "2", Name: "Banana", Price: 0.99, Image: "banana.jpg"},
	}

	mockService.EXPECT().GetLineItems().Return(expectedLineItems, nil)

	router.GET("/api/lineitems", controller.GetLineItems)

	mocks.PerformRequest(t, router, "GET", "/api/lineitems", nil, recorder)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var actual []types.LineItem
	mocks.UnmarshalResponse(t, recorder, &actual)

	assert.Equal(t, expectedLineItems, actual)
}

func TestCreateLineItem(t *testing.T) {
	router, mockService, controller, recorder := LineItemControllerTesr(t)

	newLineItem := &types.LineItem{ID: "3", Name: "Grape", Price: 2.50, Image: "grape.jpg"}
	newLineItemJSON, _ := json.Marshal(newLineItem)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreateLineItem(newLineItem).Return(expectedResult, nil)

	router.POST("/api/lineitems", controller.CreateLineItem)
	mocks.PerformRequest(t, router, "POST", "/api/lineitems", bytes.NewBuffer(newLineItemJSON), recorder)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var actual types.ModelResult
	mocks.UnmarshalResponse(t, recorder, &actual)

	assert.Equal(t, expectedResult, &actual)
}
