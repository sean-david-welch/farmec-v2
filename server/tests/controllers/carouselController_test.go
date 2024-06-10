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

func CarouselControllerTest(test *testing.T) (*gin.Engine, *mocks.MockCarouselService, *handlers.CarouselHandler, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockCarouselService(ctrl)
	controller := handlers.NewCarouselController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetCarousel(test *testing.T) {
	router, mockService, controller, recorder, time := CarouselControllerTest(test)

	expectedCarousel := []types.Carousel{
		{
			ID:      "1",
			Name:    "Carousel 1",
			Image:   "image1.jpg",
			Created: time,
		},
		{
			ID:      "2",
			Name:    "Carousel 2",
			Image:   "01/01/24",
			Created: time,
		},
	}

	mockService.EXPECT().GetCarousels().Return(expectedCarousel, nil)

	router.GET("/api/carousels", controller.GetCarousels)

	mocks.PerformRequest(test, router, "GET", "/api/carousels", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Carousel
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedCarousel, actual)
}

func TestCreateCarousel(test *testing.T) {
	router, mockService, controller, recorder, time := CarouselControllerTest(test)

	newCarousel := &types.Carousel{
		ID:      "1",
		Name:    "Carousel 1",
		Image:   "image1.jpg",
		Created: time,
	}
	jsonCarousel, _ := json.Marshal(newCarousel)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreateCarousel(newCarousel).Return(expectedResult, nil)

	router.POST("/api/carousels", controller.CreateCarousel)
	mocks.PerformRequest(test, router, "POST", "/api/carousels", bytes.NewBuffer(jsonCarousel), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actuals types.ModelResult
	mocks.UnmarshalResponse(test, recorder, &actuals)

	assert.Equal(test, expectedResult, &actuals)
}
