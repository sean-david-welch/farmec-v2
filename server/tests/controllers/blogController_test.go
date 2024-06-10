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

func BlogControllerTest(test *testing.T) (*gin.Engine, *mocks.MockBlogService, *handlers.BlogHandler, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockBlogService(ctrl)
	controller := handlers.NewBlogController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetBlogs(test *testing.T) {
	router, mockService, controller, recorder, time := BlogControllerTest(test)

	expectedBlogs := []types.Blog{
		{
			ID:         "1",
			Title:      "Blog 1",
			Date:       "01/01/24",
			MainImage:  "image1.jpg",
			Subheading: "subheading 1",
			Body:       "body 1",
			Created:    time,
		},
		{
			ID:         "2",
			Title:      "Blog 2",
			Date:       "01/01/24",
			MainImage:  "image1.jpg",
			Subheading: "subheading 2",
			Body:       "body 2",
			Created:    time,
		},
	}

	mockService.EXPECT().GetBlogs().Return(expectedBlogs, nil)

	router.GET("/api/blogs", controller.GetBlogs)

	mocks.PerformRequest(test, router, "GET", "/api/blogs", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Blog
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedBlogs, actual)

}

func TestCreateBlog(test *testing.T) {
	router, mockService, controller, recorder, time := BlogControllerTest(test)

	newBlog := &types.Blog{
		ID:         "3",
		Title:      "Blog 3",
		Date:       "01/01/24",
		MainImage:  "image3.jpg",
		Subheading: "subheading 3",
		Body:       "body 3",
		Created:    time,
	}
	newBlogJSON, _ := json.Marshal(newBlog)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreateBlog(newBlog).Return(expectedResult, nil)

	router.POST("/api/blogs", controller.CreateBlog)
	mocks.PerformRequest(test, router, "POST", "/api/blogs", bytes.NewBuffer(newBlogJSON), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actual types.ModelResult
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedResult, &actual)
}
