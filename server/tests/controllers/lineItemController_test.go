package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func setupControllerTest(test *testing.T) (*gin.Engine, *mocks.MockLineItemService,  *controllers.LineItemController, *httptest.ResponseRecorder) {
    gin.SetMode(gin.TestMode)

    mockService := new(mocks.MockLineItemService)
    controller := controllers.NewLineItemController(mockService)

    router := gin.Default()
    recorder := httptest.NewRecorder()

    return router, mockService, controller, recorder
}

func TestGetLineItems(test *testing.T) {
    router, mockService, controller, recorder := setupControllerTest(test)

    expectedLineItems := []types.LineItem{
        {ID: "1", Name: "Apple", Price: 1.99, Image: "apple.jpg"},
        {ID: "2", Name: "Banana", Price: 0.99, Image: "banana.jpg"},
    }

    mockService.On("GetLineItems").Return(expectedLineItems, nil)

    router.GET("/lineitems", controller.GetLineItems)

    mocks.PerformRequest(test, router, "GET", "/lineitems", nil, recorder)
    
    assert.Equal(test, http.StatusOK, recorder.Code)

    var actual []types.LineItem
    mocks.UnmarshalResponse(test, recorder, &actual)


    assert.Equal(test, expectedLineItems, actual)
    mockService.AssertExpectations(test)
}

func TestCreateLineItem(test *testing.T) {
    router, mockService, controller, recorder := setupControllerTest(test)

    newLineItem := &types.LineItem{ID: "3", Name: "Grape", Price: 2.50, Image: "grape.jpg"}
    newLineItemJSON, err := json.Marshal(newLineItem); if err != nil {
        test.Fatal("Error marshalling new line item:", err)
    }

    expectedResult := &types.ModelResult{PresginedUrl: "presigned-url", ImageUrl: "image-url"}
    mockService.On("CreateLineItem", newLineItem).Return(expectedResult, nil)

    router.POST("/lineitems", controller.CreateLineItem)
    mocks.PerformRequest(test, router, "POST", "/lineitems", bytes.NewBuffer(newLineItemJSON), recorder)

    assert.Equal(test, http.StatusCreated, recorder.Code)

    var actual types.ModelResult
    mocks.UnmarshalResponse(test, recorder, &actual)


    assert.Equal(test, expectedResult, &actual)
    mockService.AssertExpectations(test)
}