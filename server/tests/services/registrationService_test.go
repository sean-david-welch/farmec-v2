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

func RegistrationTestService(test *testing.T) (*mocks.MockRegistrationRepository, services.RegistrationService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockRegistrationRepository(controller)
	service := services.NewRegistrationService(mockRepo)

	return mockRepo, service
}

func TestGetRegistrations(test *testing.T) {
	mockRepo, service := RegistrationTestService(test)

	expectedRegistrations := []types.MachineRegistration{
		{
			ID:               "id ",
			DealerName:       "dealer name",
			DealerAddress:    "dealer address",
			OwnerName:        "owner name",
			OwnerAddress:     "owner address",
			MachineModel:     "machine model",
			SerialNumber:     "serial number",
			InstallDate:      "install date",
			InvoiceNumber:    "invoice number",
			CompleteSupply:   true,
			PdiComplete:      true,
			PtoCorrect:       true,
			MachineTestRun:   true,
			SafetyInduction:  true,
			OperatorHandbook: true,
			Date:             "date",
			CompletedBy:      "completed by",
			Created:          time.Now(),
		},
		{
			ID:               "id 2",
			DealerName:       "dealer name",
			DealerAddress:    "dealer address",
			OwnerName:        "owner name",
			OwnerAddress:     "owner address",
			MachineModel:     "machine model",
			SerialNumber:     "serial number",
			InstallDate:      "install date",
			InvoiceNumber:    "invoice number",
			CompleteSupply:   true,
			PdiComplete:      true,
			PtoCorrect:       true,
			MachineTestRun:   true,
			SafetyInduction:  true,
			OperatorHandbook: true,
			Date:             "date",
			CompletedBy:      "completed by",
			Created:          time.Now(),
		},
	}

	mockRepo.EXPECT().GetRegistrations().Return(expectedRegistrations, nil)

	registrations, err := service.GetRegistrations()

	assert.NoError(test, err)
	assert.Equal(test, registrations, expectedRegistrations)
}

func TestCreateRegistration(test *testing.T) {
	mockRepo, service := RegistrationTestService(test)

	newRegistration := &types.MachineRegistration{
		ID:               "id ",
		DealerName:       "dealer name",
		DealerAddress:    "dealer address",
		OwnerName:        "owner name",
		OwnerAddress:     "owner address",
		MachineModel:     "machine model",
		SerialNumber:     "serial number",
		InstallDate:      "install date",
		InvoiceNumber:    "invoice number",
		CompleteSupply:   true,
		PdiComplete:      true,
		PtoCorrect:       true,
		MachineTestRun:   true,
		SafetyInduction:  true,
		OperatorHandbook: true,
		Date:             "date",
		CompletedBy:      "completed by",
		Created:          time.Now(),
	}

	mockRepo.EXPECT().CreateRegistration(newRegistration).Return(nil)

	err := service.CreateRegistration(newRegistration)

	assert.NoError(test, err)
}
