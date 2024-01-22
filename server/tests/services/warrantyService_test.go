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

func WarrantyTestService(test *testing.T) (*mocks.MockWarrantyRepository, services.WarrantyService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockWarrantyRepository(controller)
	service := services.NewWarrantyService(mockRepo)

	return mockRepo, service
}

func TestGetWarrantyById(test *testing.T) {
	mockRepo, service := WarrantyTestService(test)

	expectedWarranty := &types.WarrantyClaim{
		ID:             "warranty123",
		Dealer:         "Dealer XYZ",
		DealerContact:  "John Doe",
		OwnerName:      "Alice Johnson",
		MachineModel:   "Model ABC",
		SerialNumber:   "SN123456",
		InstallDate:    "2024-01-15",
		FailureDate:    "2024-02-10",
		RepairDate:     "2024-02-20",
		FailureDetails: "Engine malfunction",
		RepairDetails:  "Replaced engine",
		LabourHours:    "5",
		CompletedBy:    "Technician A",
		Created:        time.Now(),
	}

	expectedParts := []types.PartsRequired{
		{
			ID:             "part1",
			WarrantyID:     "warranty123",
			PartNumber:     "PN123",
			QuantityNeeded: "1",
			InvoiceNumber:  "INV123",
			Description:    "Engine",
		},
		{
			ID:             "part2",
			WarrantyID:     "warranty123",
			PartNumber:     "PN124",
			QuantityNeeded: "2",
			InvoiceNumber:  "INV124",
			Description:    "Oil Filter",
		},
	}

	warrantyId := expectedWarranty.ID

	mockRepo.EXPECT().GetWarrantyById(warrantyId).Return(expectedWarranty, expectedParts, nil)

	warranties, parts, err := service.GetWarrantyById(warrantyId)

	assert.NoError(test, err)
	assert.Equal(test, warranties, expectedWarranty)
	assert.Equal(test, parts, expectedParts)
}

func TestCreateWarranty(test *testing.T) {
	mockRepo, service := WarrantyTestService(test)

	newWarranty := &types.WarrantyClaim{
		ID:             "warranty123",
		Dealer:         "Dealer XYZ",
		DealerContact:  "John Doe",
		OwnerName:      "Alice Johnson",
		MachineModel:   "Model ABC",
		SerialNumber:   "SN123456",
		InstallDate:    "2024-01-15",
		FailureDate:    "2024-02-10",
		RepairDate:     "2024-02-20",
		FailureDetails: "Engine malfunction",
		RepairDetails:  "Replaced engine",
		LabourHours:    "5",
		CompletedBy:    "Technician A",
		Created:        time.Now(),
	}

	newParts := []types.PartsRequired{
		{
			ID:             "part1",
			WarrantyID:     "warranty123",
			PartNumber:     "PN123",
			QuantityNeeded: "1",
			InvoiceNumber:  "INV123",
			Description:    "Engine",
		},
		{
			ID:             "part2",
			WarrantyID:     "warranty123",
			PartNumber:     "PN124",
			QuantityNeeded: "2",
			InvoiceNumber:  "INV124",
			Description:    "Oil Filter",
		},
	}

	mockRepo.EXPECT().CreateWarranty(newWarranty, newParts).Return(nil)

	err := service.CreateWarranty(newWarranty, newParts)

	assert.NoError(test, err)
}
