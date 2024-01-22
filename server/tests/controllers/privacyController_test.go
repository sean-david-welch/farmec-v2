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

func PrivacyControllerTest(test *testing.T) (*gin.Engine, *mocks.MockPrivacyService, *controllers.PrivacyController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockPrivacyService(ctrl)
	controller := controllers.NewPrivacyController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetPrivacy(test *testing.T) {
	router, mockService, controller, recorder, time := PrivacyControllerTest(test)

	expectedPrivacy := []types.Privacy{
		{
			ID:      "1",
			Title:   "Privacy 1",
			Body:    "body 1",
			Created: time,
		},
		{
			ID:      "2",
			Title:   "Privacy 2",
			Body:    "body 2",
			Created: time,
		},
	}

	mockService.EXPECT().GetPrivacys().Return(expectedPrivacy, nil)

	router.GET("/api/privacy", controller.GetPrivacys)

	mocks.PerformRequest(test, router, "GET", "/api/privacy", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Privacy
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedPrivacy, actual)
}

func TestCreatePrivacy(test *testing.T) {
	router, mockService, controller, recorder, time := PrivacyControllerTest(test)

	privacy := &types.Privacy{
		ID:      "1",
		Title:   "Privacy 1",
		Body:    "body 1",
		Created: time,
	}
	jsonPrivacy, _ := json.Marshal(privacy)

	mockService.EXPECT().CreatePrivacy(privacy).Return(nil)

	router.POST("/api/privacy", controller.CreatePrivacy)
	mocks.PerformRequest(test, router, "POST", "/api/privacy", bytes.NewBuffer(jsonPrivacy), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)
}
