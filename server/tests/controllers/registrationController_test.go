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

func RegistrationControllerTest(test *testing.T) (*gin.Engine, *mocks.MockRegistrationService, *handlers.RegistrationController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockRegistrationService(ctrl)
	controller := handlers.NewRegistrationController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetRegistrations(test *testing.T) {
	router, mockService, controller, recorder, time := RegistrationControllerTest(test)

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
			Created:          time,
		},
	}

	mockService.EXPECT().GetRegistrations().Return(expectedRegistrations, nil)

	router.GET("/api/registrations", controller.GetRegistrations)

	mocks.PerformRequest(test, router, "GET", "/api/registrations", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.MachineRegistration
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedRegistrations, actual)
}

func TestCreateRegistration(test *testing.T) {
	router, mockService, controller, recorder, time := RegistrationControllerTest(test)

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
		Created:          time,
	}

	jsonRegistration, _ := json.Marshal(newRegistration)

	mockService.EXPECT().CreateRegistration(newRegistration).Return(nil)

	router.POST("/api/registrations", controller.CreateRegistration)
	mocks.PerformRequest(test, router, "POST", "/api/registrations", bytes.NewBuffer(jsonRegistration), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)
}
