package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func PartsControllerTest(test *testing.T) (*gin.Engine, *mocks.MockPartsService, *controllers.PartsController, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockPartsService(ctrl)
	controller := controllers.NewPartsController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	return router, mockService, controller, recorder
}

func TestGetParts(test *testing.T) {
	router, mockService, controller, recorder := PartsControllerTest(test)

	expectedParts := []types.Sparepart{
		{
			ID:             "1",
			SupplierID:     "12",
			Name:           "Sparepart 1",
			PartsImage:     "image1.jpg",
			SparePartsLink: "sparePartsLink 1",
		},
		{
			ID:             "2",
			SupplierID:     "12",
			Name:           "Sparepart 2",
			PartsImage:     "image1.jpg",
			SparePartsLink: "sparePartsLink 2",
		},
	}

	supplierId := expectedParts[0].SupplierID

	mockService.EXPECT().GetParts(supplierId).Return(expectedParts, nil)

	router.GET("/api/spareparts/:id", controller.GetParts)

	url := fmt.Sprintf("/api/spareparts/%s", supplierId)
	mocks.PerformRequest(test, router, "GET", url, nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Sparepart
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedParts, actual)
}

func TestCreatePart(test *testing.T) {
	router, mockService, controller, recorder := PartsControllerTest(test)

	newPart := &types.Sparepart{
		ID:             "1",
		SupplierID:     "12",
		Name:           "Sparepart 1",
		PartsImage:     "image1.jpg",
		SparePartsLink: "sparePartsLink 1",
	}
	jsonPart, _ := json.Marshal(newPart)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreatePart(newPart).Return(expectedResult, nil)

	router.POST("/api/spareparts", controller.CreateParts)
	mocks.PerformRequest(test, router, "POST", "/api/spareparts", bytes.NewBuffer(jsonPart), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actuals types.ModelResult
	mocks.UnmarshalResponse(test, recorder, &actuals)

	assert.Equal(test, expectedResult, &actuals)
}
