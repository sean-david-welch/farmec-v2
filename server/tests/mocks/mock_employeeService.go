// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/employeeService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sean-david-welch/farmec-v2/server/types"
)

// MockEmployeeService is a mock of EmployeeService interface.
type MockEmployeeService struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeServiceMockRecorder
}

// MockEmployeeServiceMockRecorder is the mock recorder for MockEmployeeService.
type MockEmployeeServiceMockRecorder struct {
	mock *MockEmployeeService
}

// NewMockEmployeeService creates a new mock instance.
func NewMockEmployeeService(ctrl *gomock.Controller) *MockEmployeeService {
	mock := &MockEmployeeService{ctrl: ctrl}
	mock.recorder = &MockEmployeeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeService) EXPECT() *MockEmployeeServiceMockRecorder {
	return m.recorder
}

// CreateEmployee mocks base method.
func (m *MockEmployeeService) CreateEmployee(employee *types.Employee) (*types.ModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", employee)
	ret0, _ := ret[0].(*types.ModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockEmployeeServiceMockRecorder) CreateEmployee(employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockEmployeeService)(nil).CreateEmployee), employee)
}

// DeleteEmployee mocks base method.
func (m *MockEmployeeService) DeleteEmployee(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee.
func (mr *MockEmployeeServiceMockRecorder) DeleteEmployee(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockEmployeeService)(nil).DeleteEmployee), id)
}

// GetEmployees mocks base method.
func (m *MockEmployeeService) GetEmployees() ([]types.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployees")
	ret0, _ := ret[0].([]types.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployees indicates an expected call of GetEmployees.
func (mr *MockEmployeeServiceMockRecorder) GetEmployees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployees", reflect.TypeOf((*MockEmployeeService)(nil).GetEmployees))
}

// UpdateEmployee mocks base method.
func (m *MockEmployeeService) UpdateEmployee(id string, employee *types.Employee) (*types.ModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", id, employee)
	ret0, _ := ret[0].(*types.ModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockEmployeeServiceMockRecorder) UpdateEmployee(id, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmployeeService)(nil).UpdateEmployee), id, employee)
}