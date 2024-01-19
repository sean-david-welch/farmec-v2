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

func setupControllerTest(t *testing.T) (*gin.Engine, *mocks.MockLineItemService,  *controllers.LineItemController, *httptest.ResponseRecorder) {
    gin.SetMode(gin.TestMode)

    mockService := new(mocks.MockLineItemService)
    controller := controllers.NewLineItemController(mockService)

    router := gin.Default()
    recorder := httptest.NewRecorder()

    return router, mockService, controller, recorder
}

func TestGetLineItems(t *testing.T) {
    // Setup the test environment
    router, mockService, controller, recorder := setupControllerTest(t)

    // Define expected data
    expectedLineItems := []types.LineItem{
        {ID: "1", Name: "Apple", Price: 1.99, Image: "apple.jpg"},
        {ID: "2", Name: "Banana", Price: 0.99, Image: "banana.jpg"},
    }

    // Setup mock expectations
    mockService.On("GetLineItems").Return(expectedLineItems, nil)

    // Register route and make the test request
    router.GET("/lineitems", controller.GetLineItems)
    req, _ := http.NewRequest("GET", "/lineitems", nil)
    router.ServeHTTP(recorder, req)

    // Assert status code
    assert.Equal(t, http.StatusOK, recorder.Code)

    // Deserialize response and assert equality
    var actual []types.LineItem
    err := json.Unmarshal(recorder.Body.Bytes(), &actual)
    assert.NoError(t, err)
    assert.Equal(t, expectedLineItems, actual)

    // Verify that the expectations were met
    mockService.AssertExpectations(t)
}

func TestCreateLineItem(t *testing.T) {
    // Setup the test environment
    router, mockService, controller, recorder := setupControllerTest(t)

    // Define input and expected result
    newLineItem := &types.LineItem{ID: "3", Name: "Grape", Price: 2.50, Image: "grape.jpg"}
    newLineItemJson, _ := json.Marshal(newLineItem)
    expectedResult := &types.ModelResult{PresginedUrl: "presigned-url", ImageUrl: "image-url"}

    // Setup mock expectations
    mockService.On("CreateLineItem", newLineItem).Return(expectedResult, nil)

    // Register route and make the test request
    router.POST("/lineitems", controller.CreateLineItem)
    req, _ := http.NewRequest("POST", "/lineitems", bytes.NewBuffer(newLineItemJson))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(recorder, req)

    // Assert status code
    assert.Equal(t, http.StatusCreated, recorder.Code)

    // Deserialize response and assert specific fields
    var actual map[string]interface{}
    err := json.Unmarshal(recorder.Body.Bytes(), &actual)
    assert.NoError(t, err)
    assert.Equal(t, expectedResult.PresginedUrl, actual["presginedUrl"])
    assert.Equal(t, expectedResult.ImageUrl, actual["imageUrl"])

    // Verify that the expectations were met
    mockService.AssertExpectations(t)
}


