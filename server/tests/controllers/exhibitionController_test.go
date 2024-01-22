package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func ExhibitionControllerTest(test *testing.T) (*gin.Engine, *mocks.MockExhibitionService, *controllers.ExhibitionController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockExhibitionService(ctrl)
	controller := controllers.NewExhibitionController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetExhibitions(test *testing.T) {
	router, mockService, controller, recorder, time := ExhibitionControllerTest(test)

	expectedExhibitions := []types.Exhibition{
		{
			ID:       "1",
			Title:    "Exhibition 1",
			Date:     "01/01/24",
			Location: "Dublin",
			Info:     "description 1",
			Created:  time,
		},
		{
			ID:       "2",
			Title:    "Exhibition 2",
			Date:     "01/01/24",
			Location: "Dublin",
			Info:     "description 2",
			Created:  time,
		},
	}

	mockService.EXPECT().GetExhibitions().Return(expectedExhibitions, nil)

	router.GET("/api/exhibitions", controller.GetExhibitions)

	mocks.PerformRequest(test, router, "GET", "/api/exhibitions", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Exhibition
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedExhibitions, actual)
}

func TestCreateExhibition(test *testing.T) {
	router, mockService, controller, recorder, time := ExhibitionControllerTest(test)

	newExhibition := &types.Exhibition{
		ID:       "1",
		Title:    "Exhibition 1",
		Date:     "01/01/24",
		Location: "Dublin",
		Info:     "description 1",
		Created:  time,
	}
	newExhibitionJson, _ := json.Marshal(newExhibition)

	mockService.EXPECT().CreateExhibition(newExhibition).Return(nil)

	router.POST("/api/exhibitions", controller.CreateExhibition)
	mocks.PerformRequest(test, router, "POST", "/api/exhibitions", bytes.NewBuffer(newExhibitionJson), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actual types.ModelResult
	mocks.UnmarshalResponse(test, recorder, &actual)
}
