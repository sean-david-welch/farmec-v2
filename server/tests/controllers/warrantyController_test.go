package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func WarrantyControllerTest(test *testing.T) (*gin.Engine, *mocks.MockWarrantyService, *controllers.WarrantyController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockWarrantyService(ctrl)
	controller := controllers.NewWarrantyController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetWarranties(test *testing.T) {
	router, mockService, controller, recorder, time := WarrantyControllerTest(test)

	DealerContact := "John Doe"
	OwnerName := "Alice Johnson"
	MachineModel := "Model ABC"
	SerialNumber := "SN123456"
	InstallDate := "2024-01-15"
	FailureDate := "2024-02-10"
	RepairDate := "2024-02-20"
	FailureDetails := "Engine malfunction"
	RepairDetails := "Replaced engine"
	LabourHours := "5"
	CompletedBy := "Technician A"

	expectedWarranty := &types.WarrantyClaim{
		ID:             "warranty123",
		Dealer:         "Dealer XYZ",
		DealerContact:  &DealerContact,
		OwnerName:      &OwnerName,
		MachineModel:   &MachineModel,
		SerialNumber:   &SerialNumber,
		InstallDate:    &InstallDate,
		FailureDate:    &FailureDate,
		RepairDate:     &RepairDate,
		FailureDetails: &FailureDetails,
		RepairDetails:  &RepairDetails,
		LabourHours:    &LabourHours,
		CompletedBy:    &CompletedBy,
		Created:        time,
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

	warranty_id := expectedWarranty.ID

	mockService.EXPECT().GetWarrantyById(warranty_id).Return(expectedWarranty, expectedParts, nil)

	router.GET("/api/warranties/:id", controller.GetWarrantyById)

	url := fmt.Sprintf("/api/warranties/%s", warranty_id)
	mocks.PerformRequest(test, router, "GET", url, nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)
}

func TestCreateWarranty(test *testing.T) {
	router, mockService, controller, recorder, time := WarrantyControllerTest(test)

	DealerContact := "John Doe"
	OwnerName := "Alice Johnson"
	MachineModel := "Model ABC"
	SerialNumber := "SN123456"
	InstallDate := "2024-01-15"
	FailureDate := "2024-02-10"
	RepairDate := "2024-02-20"
	FailureDetails := "Engine malfunction"
	RepairDetails := "Replaced engine"
	LabourHours := "5"
	CompletedBy := "Technician A"

	newWarranty := &types.WarrantyClaim{
		ID:             "warranty123",
		Dealer:         "Dealer XYZ",
		DealerContact:  &DealerContact,
		OwnerName:      &OwnerName,
		MachineModel:   &MachineModel,
		SerialNumber:   &SerialNumber,
		InstallDate:    &InstallDate,
		FailureDate:    &FailureDate,
		RepairDate:     &RepairDate,
		FailureDetails: &FailureDetails,
		RepairDetails:  &RepairDetails,
		LabourHours:    &LabourHours,
		CompletedBy:    &CompletedBy,
		Created:        time,
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

	combinedData := &types.WarranrtyParts{
		Warranty: newWarranty,
		Parts:    newParts,
	}

	jsonCombined, _ := json.Marshal(combinedData)

	mockService.EXPECT().CreateWarranty(newWarranty, newParts).Return(nil)

	router.POST("/api/warranties", controller.CreateWarranty)
	mocks.PerformRequest(test, router, "POST", "/api/warranties", bytes.NewBuffer(jsonCombined), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)
}
