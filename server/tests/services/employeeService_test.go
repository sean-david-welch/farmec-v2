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

func EmployeeTestService(test *testing.T) (*mocks.MockEmployeeRepository, *mocks.MockS3Client, services.EmployeeService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(controller)
	mockS3Client := mocks.NewMockS3Client(controller)
	service := services.NewEmployeeService(mockRepo, mockS3Client, "test-folder")

	return mockRepo, mockS3Client, service
}

func TestGetEmployees(test *testing.T) {
	mockRepo, _, service := EmployeeTestService(test)

	expectedEmployees := []types.Employee{
		{
			ID:           "1",
			Name:         "name 1",
			Email:        "email1",
			Role:         "role 1",
			ProfileImage: "image1.jpg",
			Created:      time.Now(),
		},
		{
			ID:           "2",
			Name:         "name 2",
			Email:        "email 2",
			Role:         "role 2",
			ProfileImage: "image2.jpg",
			Created:      time.Now(),
		},
	}

	mockRepo.EXPECT().GetEmployees().Return(expectedEmployees, nil)

	employees, err := service.GetEmployees()

	assert.NoError(test, err)
	assert.Equal(test, employees, expectedEmployees)
}

func TestCreateEmployeee(test *testing.T) {
	mockRepo, mockS3Client, service := EmployeeTestService(test)

	newEmployee := &types.Employee{
		ID:           "2",
		Name:         "name 2",
		Email:        "email2",
		Role:         "role 2",
		ProfileImage: "image2.jpg",
		Created:      time.Now(),
	}

	mockS3Client.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presigned-url", "image-url", nil)
	mockRepo.EXPECT().CreateEmployee(newEmployee).Return(nil)

	result, err := service.CreateEmployee(newEmployee)

	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, "image-url", result.ImageUrl)
	assert.Equal(test, "presigned-url", result.PresginedUrl)
}
