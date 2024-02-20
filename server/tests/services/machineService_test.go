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

func MachineTestService(test *testing.T) (*mocks.MockMachineRepository, *mocks.MockS3Client, services.MachineService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockMachineRepository(controller)
	mockS3Client := mocks.NewMockS3Client(controller)
	service := services.NewMachineService(mockRepo, mockS3Client, "test-folder")

	return mockRepo, mockS3Client, service
}

func TestGetMachine(test *testing.T) {
	mockRepo, _, service := MachineTestService(test)

	description := "description"
	machinelink := "machinelink"

	expectedMachines := []types.Machine{
		{
			ID:           "id ",
			SupplierID:   "12",
			Name:         "name 1",
			MachineImage: "image.jpg",
			Description:  &description,
			MachineLink:  &machinelink,
			Created:      time.Now(),
		},
		{
			ID:           "id 2",
			SupplierID:   "12",
			Name:         "name 2",
			MachineImage: "image.jpg",
			Description:  &description,
			MachineLink:  &machinelink,
			Created:      time.Now(),
		},
	}

	supplier_id := expectedMachines[0].SupplierID

	mockRepo.EXPECT().GetMachines(supplier_id).Return(expectedMachines, nil)

	machines, err := service.GetMachines(supplier_id)

	assert.NoError(test, err)
	assert.Equal(test, machines, expectedMachines)
}

func TestCreateMachine(test *testing.T) {
	mockRepo, mockS3Client, service := MachineTestService(test)

	newMachine := &types.Machine{
		ID:           "2",
		SupplierID:   "12",
		Name:         "name 2",
		MachineImage: "image2.jpg",
		Created:      time.Now(),
	}

	mockS3Client.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presigned-url", "image-url", nil)
	mockRepo.EXPECT().CreateMachine(newMachine).Return(nil)

	result, err := service.CreateMachine(newMachine)

	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, "presigned-url", result.PresignedUrl)
	assert.Equal(test, "image-url", result.ImageUrl)
}
