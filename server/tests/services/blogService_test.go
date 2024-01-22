package services_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func BlogTestService(test *testing.T) (*mocks.MockBlogRepository, *mocks.MockS3Client, services.BlogService) {
	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogRepository(ctrl)
	mockS3Client := mocks.NewMockS3Client(ctrl)
	service := services.NewBlogService(mockRepo, mockS3Client, "test-folder")

	return mockRepo, mockS3Client, service
}

func TestGetBlogs(test *testing.T) {
	mockRepo, _, service := BlogTestService(test)

	expectedBlogs := []types.Blog{
		{
			ID:         "1",
			Title:      "Blog 1",
			Date:       "01/01/24",
			MainImage:  "image1.jpg",
			Subheading: "subheading1",
			Body:       "body 1",
			Created:    time.Now(),
		},
		{
			ID:         "1",
			Title:      "Blog 1",
			Date:       "01/01/24",
			MainImage:  "image1.jpg",
			Subheading: "subheading1",
			Body:       "body 1",
			Created:    time.Now(),
		},
	}

	mockRepo.EXPECT().GetBlogs().Return(expectedBlogs, nil)

	blogs, err := service.GetBlogs()

	assert.NoError(test, err)
	assert.Equal(test, expectedBlogs, blogs)
}

func TestCreateBlog(test *testing.T) {
	mockRepo, mockS3Client, service := BlogTestService(test)

	newBlog := &types.Blog{
		ID:         "1",
		Title:      "Blog 1",
		Date:       "01/01/24",
		MainImage:  "image3.jpg",
		Subheading: "subheading1",
		Body:       "body 1",
		Created:    time.Now(),
	}

	mockS3Client.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presigned-url", "image-url", nil)
	mockRepo.EXPECT().CreateBlog(newBlog).Return(nil)

	result, err := service.CreateBlog(newBlog)

	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, "image-url", result.ImageUrl)
	assert.Equal(test, "presigned-url", result.PresignedUrl)
}
