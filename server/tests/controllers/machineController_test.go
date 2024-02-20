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

func MachineControllerTest(test *testing.T) (*gin.Engine, *mocks.MockMachineService, *controllers.MachineController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockMachineService(ctrl)
	controller := controllers.NewMachineController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetMachines(test *testing.T) {
	router, mockService, controller, recorder, time := MachineControllerTest(test)

	description := "description 1"
	machineLink := "machineLink 1"

	expectedMachines := []types.Machine{
		{
			ID:           "1",
			SupplierID:   "12",
			Name:         "Machine 1",
			MachineImage: "image1.jpg",
			Description:  &description,
			MachineLink:  &machineLink,
			Created:      time,
		},
		{
			ID:           "2",
			SupplierID:   "12",
			Name:         "Machine 2",
			MachineImage: "image1.jpg",
			Description:  &description,
			MachineLink:  &machineLink,
			Created:      time,
		},
	}

	supplier_id := expectedMachines[0].SupplierID

	mockService.EXPECT().GetMachines(supplier_id).Return(expectedMachines, nil)

	router.GET("/api/machines/:id", controller.GetMachines)

	url := fmt.Sprintf("/api/machines/%s", supplier_id)
	mocks.PerformRequest(test, router, "GET", url, nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Machine
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedMachines, actual)
}

func TestCreateMachine(t *testing.T) {
	router, mockService, controller, recorder, time := MachineControllerTest(t)

	newMachine := &types.Machine{
		ID:           "3",
		SupplierID:   "12",
		Name:         "Machine 3",
		MachineImage: "image1.jpg",
		Description:  nil,
		MachineLink:  nil,
		Created:      time,
	}
	jsonMachine, _ := json.Marshal(newMachine)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreateMachine(newMachine).Return(expectedResult, nil)

	router.POST("/api/machines", controller.CreateMachine)
	mocks.PerformRequest(t, router, "POST", "/api/machines", bytes.NewBuffer(jsonMachine), recorder)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var actual types.ModelResult
	mocks.UnmarshalResponse(t, recorder, &actual)

	assert.Equal(t, expectedResult, &actual)
}
