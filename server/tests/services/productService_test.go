package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func ProductTestService(test *testing.T) (*mocks.MockProductRepository, *mocks.MockS3Client, services.ProductService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockProductRepository(controller)
	mocks3 := mocks.NewMockS3Client(controller)
	service := services.NewProductService(mockRepo, mocks3, "test-folder")

	return mockRepo, mocks3, service
}

func TestGetProducts(test *testing.T) {
	mockRepo, _, service := ProductTestService(test)

	expectedProducts := []types.Product{
		{
			ID:           "id ",
			MachineID:    "12",
			Name:         "name 1",
			ProductImage: "image.jpg",
			Description:  "description",
			ProductLink:  "productlink",
		},
		{
			ID:           "id 2",
			MachineID:    "12",
			Name:         "name 2",
			ProductImage: "image.jpg",
			Description:  "description",
			ProductLink:  "productlink",
		},
	}

	machineId := expectedProducts[0].MachineID

	mockRepo.EXPECT().GetProducts(machineId).Return(expectedProducts, nil)

	products, err := service.GetProducts(machineId)

	assert.NoError(test, err)
	assert.Equal(test, products, expectedProducts)
}

func TestCreateProduct(test *testing.T) {
	mockRepo, mockS3, service := ProductTestService(test)

	newProduct := &types.Product{
		ID:           "id ",
		MachineID:    "12",
		Name:         "name 1",
		ProductImage: "image.jpg",
		Description:  "description",
		ProductLink:  "productlink",
	}

	mockS3.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presignedurl", "imageurl", nil)
	mockRepo.EXPECT().CreateProduct(newProduct).Return(nil)

	result, err := service.CreateProduct(newProduct)

	assert.NoError(test, err)
	assert.NotNil(test, result)

	assert.Equal(test, "presignedurl", result.PresignedUrl)
	assert.Equal(test, "imageurl", result.ImageUrl)
}
