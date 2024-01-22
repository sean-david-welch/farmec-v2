package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func CarouselTestService(test *testing.T) (*mocks.MockCarouselRepository, *mocks.MockS3Client, services.CarouselService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockCarouselRepository(controller)
	mockS3Client := mocks.NewMockS3Client(controller)
	service := services.NewCarouselService(mockRepo, mockS3Client, "test-folder")

	return mockRepo, mockS3Client, service
}

func TestGetCarousels(test *testing.T) {
	mockRepo, _, service := CarouselTestService(test)

	expectedCarousels := []types.Carousel{
		{
			ID:    "1",
			Name:  "Image 1",
			Image: "image1.jpg",
		},
		{
			ID:    "2",
			Name:  "Image 2",
			Image: "image2.jpg",
		},
	}

	mockRepo.EXPECT().GetCarousels().Return(expectedCarousels, nil)

	carousels, err := service.GetCarousels()

	assert.NoError(test, err)
	assert.Equal(test, expectedCarousels, carousels)
}

func TestCreateCarousel(test *testing.T) {
	mockRepo, mockS3Client, service := CarouselTestService(test)

	newCarousel := &types.Carousel{
		ID:    "2",
		Name:  "Image 2",
		Image: "image3.jpg",
	}

	mockS3Client.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.All()).Return("presigned-url", "image-url", nil)
	mockRepo.EXPECT().CreateCarousel(newCarousel).Return(nil)

	result, err := service.CreateCarousel(newCarousel)

	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, "image-url", result.ImageUrl)
	assert.Equal(test, "presigned-url", result.PresignedUrl)
}
