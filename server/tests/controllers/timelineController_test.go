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

func TimelineControllerTest(test *testing.T) (*gin.Engine, *mocks.MockTimelineService, *controllers.TimelineController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimelineService(ctrl)
	controller := controllers.NewTimelineController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetTimelines(test *testing.T) {
	router, mockService, controller, recorder, time := TimelineControllerTest(test)

	expectedTimelines := []types.Timeline{
		{
			ID:      "1",
			Title:   "Timeline 1",
			Date:    "01/01/24",
			Body:    "body 1",
			Created: time,
		},
		{
			ID:      "2",
			Title:   "Timeline 2",
			Date:    "01/01/24",
			Body:    "body 2",
			Created: time,
		},
	}

	mockService.EXPECT().GetTimelines().Return(expectedTimelines, nil)

	router.GET("/api/timelines", controller.GetTimelines)

	mocks.PerformRequest(test, router, "GET", "/api/timelines", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Timeline
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedTimelines, actual)
}

func TestCreateTimeline(test *testing.T) {
	router, mockService, controller, recorder, time := TimelineControllerTest(test)

	newTimeline := &types.Timeline{
		ID:      "1",
		Title:   "Timeline 1",
		Date:    "01/01/24",
		Body:    "body 1",
		Created: time,
	}
	jsonTimeline, _ := json.Marshal(newTimeline)

	mockService.EXPECT().CreateTimeline(newTimeline).Return(nil)

	router.POST("/api/timelines", controller.CreateTimeline)

	mocks.PerformRequest(test, router, "POST", "/api/timelines", bytes.NewBuffer(jsonTimeline), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actual types.Timeline
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, *newTimeline, actual)
}
