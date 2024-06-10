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
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TermsControllerTest(test *testing.T) (*gin.Engine, *mocks.MockTermsService, *handlers.TermsController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockTermsService(ctrl)
	controller := handlers.NewTermsController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetTerms(test *testing.T) {
	router, mockService, controller, recorder, time := TermsControllerTest(test)

	expectedTerms := []types.Terms{
		{
			ID:      "1",
			Title:   "Terms 1",
			Body:    "content 1",
			Created: time,
		},
		{
			ID:      "2",
			Title:   "Terms 2",
			Body:    "content 2",
			Created: time,
		},
	}

	mockService.EXPECT().GetTerms().Return(expectedTerms, nil)

	router.GET("/api/terms", controller.GetTerms)

	mocks.PerformRequest(test, router, "GET", "/api/terms", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Terms
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedTerms, actual)
}

func TestCreateTerms(test *testing.T) {
	router, mockService, controller, recorder, time := TermsControllerTest(test)

	expectedTerms := &types.Terms{
		ID:      "1",
		Title:   "Terms 1",
		Body:    "content 1",
		Created: time,
	}
	jsonTerms, _ := json.Marshal(expectedTerms)

	mockService.EXPECT().CreateTerm(expectedTerms).Return(nil)

	router.POST("/api/terms", controller.CreateTerm)

	mocks.PerformRequest(test, router, "POST", "/api/terms", bytes.NewBuffer(jsonTerms), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actual types.Terms
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedTerms, &actual)
}
