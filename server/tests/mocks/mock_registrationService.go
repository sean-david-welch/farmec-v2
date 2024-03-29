// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/registrationService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sean-david-welch/farmec-v2/server/types"
)

// MockRegistrationService is a mock of RegistrationService interface.
type MockRegistrationService struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationServiceMockRecorder
}

// MockRegistrationServiceMockRecorder is the mock recorder for MockRegistrationService.
type MockRegistrationServiceMockRecorder struct {
	mock *MockRegistrationService
}

// NewMockRegistrationService creates a new mock instance.
func NewMockRegistrationService(ctrl *gomock.Controller) *MockRegistrationService {
	mock := &MockRegistrationService{ctrl: ctrl}
	mock.recorder = &MockRegistrationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistrationService) EXPECT() *MockRegistrationServiceMockRecorder {
	return m.recorder
}

// CreateRegistration mocks base method.
func (m *MockRegistrationService) CreateRegistration(registration *types.MachineRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRegistration", registration)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRegistration indicates an expected call of CreateRegistration.
func (mr *MockRegistrationServiceMockRecorder) CreateRegistration(registration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRegistration", reflect.TypeOf((*MockRegistrationService)(nil).CreateRegistration), registration)
}

// DeleteRegistration mocks base method.
func (m *MockRegistrationService) DeleteRegistration(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRegistration", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRegistration indicates an expected call of DeleteRegistration.
func (mr *MockRegistrationServiceMockRecorder) DeleteRegistration(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistration", reflect.TypeOf((*MockRegistrationService)(nil).DeleteRegistration), id)
}

// GetRegistrationById mocks base method.
func (m *MockRegistrationService) GetRegistrationById(id string) (*types.MachineRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistrationById", id)
	ret0, _ := ret[0].(*types.MachineRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistrationById indicates an expected call of GetRegistrationById.
func (mr *MockRegistrationServiceMockRecorder) GetRegistrationById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistrationById", reflect.TypeOf((*MockRegistrationService)(nil).GetRegistrationById), id)
}

// GetRegistrations mocks base method.
func (m *MockRegistrationService) GetRegistrations() ([]types.MachineRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistrations")
	ret0, _ := ret[0].([]types.MachineRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistrations indicates an expected call of GetRegistrations.
func (mr *MockRegistrationServiceMockRecorder) GetRegistrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistrations", reflect.TypeOf((*MockRegistrationService)(nil).GetRegistrations))
}

// UpdateRegistration mocks base method.
func (m *MockRegistrationService) UpdateRegistration(id string, registration *types.MachineRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegistration", id, registration)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRegistration indicates an expected call of UpdateRegistration.
func (mr *MockRegistrationServiceMockRecorder) UpdateRegistration(id, registration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegistration", reflect.TypeOf((*MockRegistrationService)(nil).UpdateRegistration), id, registration)
}
