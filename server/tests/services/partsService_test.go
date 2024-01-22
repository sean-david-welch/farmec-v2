package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func PartsTestService(test *testing.T) (*mocks.MockPartsRepository, *mocks.MockS3Client, services.PartsService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockPartsRepository(controller)
	mockS3Client := mocks.NewMockS3Client(controller)
	service := services.NewPartsService(mockRepo, mockS3Client, "test-folder")

	return mockRepo, mockS3Client, service
}

func TestGetParts(test *testing.T) {
	mockRepo, _, service := PartsTestService(test)

	expectedParts := []types.Sparepart{
		{
			ID:             "id ",
			SupplierID:     "12",
			Name:           "name 1",
			PartsImage:     "image.jpg",
			SparePartsLink: "link",
		},
		{
			ID:             "id 2",
			SupplierID:     "12",
			Name:           "name 2",
			PartsImage:     "image.jpg",
			SparePartsLink: "link",
		},
	}

	supplierId := expectedParts[0].SupplierID

	mockRepo.EXPECT().GetParts(supplierId).Return(expectedParts, nil)

	parts, err := service.GetParts(supplierId)

	assert.NoError(test, err)
	assert.Equal(test, parts, expectedParts)
}

func TestCreatePart(test *testing.T) {
	mockRepo, mocks3, service := PartsTestService(test)

	newPart := &types.Sparepart{
		ID:             "id ",
		SupplierID:     "12",
		Name:           "name 1",
		PartsImage:     "image.jpg",
		SparePartsLink: "link",
	}

	mocks3.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presignedurl", "imageurl", nil)
	mockRepo.EXPECT().CreatePart(newPart).Return(nil)

	result, err := service.CreatePart(newPart)

	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, "presignedurl", result.PresginedUrl)
	assert.Equal(test, "imageurl", result.ImageUrl)
}
