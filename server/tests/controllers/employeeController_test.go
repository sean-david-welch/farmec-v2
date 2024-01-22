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
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func EmployeeControllerTest(test *testing.T) (*gin.Engine, *mocks.MockEmployeeService, *controllers.EmployeeController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockEmployeeService(ctrl)
	controller := controllers.NewEmployeeController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetEmployees(test *testing.T) {
	router, mockService, controller, recorder, time := EmployeeControllerTest(test)

	expectedEmployees := []types.Employee{
		{
			ID:           "1",
			Name:         "Employee 1",
			Email:        "email 1",
			Role:         "role 1",
			ProfileImage: "profileImage 1",
			Created:      time,
		},
		{
			ID:           "2",
			Name:         "Employee 2",
			Email:        "email 2",
			Role:         "role 2",
			ProfileImage: "profileImage 2",
			Created:      time,
		},
	}

	mockService.EXPECT().GetEmployees().Return(expectedEmployees, nil)

	router.GET("/api/employees", controller.GetEmployees)

	mocks.PerformRequest(test, router, "GET", "/api/employees", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Employee
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedEmployees, actual)
}

func TestCreateEmployee(test *testing.T) {
	router, mockService, controller, recorder, time := EmployeeControllerTest(test)

	newEmployee := &types.Employee{
		ID:           "1",
		Name:         "Employee 1",
		Email:        "email 1",
		Role:         "role 1",
		ProfileImage: "profileImage 1",
		Created:      time,
	}
	newEmployeeJson, _ := json.Marshal(newEmployee)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreateEmployee(newEmployee).Return(expectedResult, nil)

	router.POST("/api/employees", controller.CreateEmployee)
	mocks.PerformRequest(test, router, "POST", "/api/employees", bytes.NewBuffer(newEmployeeJson), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actual types.ModelResult
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedResult, &actual)
}
