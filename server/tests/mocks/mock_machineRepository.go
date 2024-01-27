// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/machineRepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sean-david-welch/farmec-v2/server/types"
)

// MockMachineRepository is a mock of MachineRepository interface.
type MockMachineRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMachineRepositoryMockRecorder
}

// MockMachineRepositoryMockRecorder is the mock recorder for MockMachineRepository.
type MockMachineRepositoryMockRecorder struct {
	mock *MockMachineRepository
}

// NewMockMachineRepository creates a new mock instance.
func NewMockMachineRepository(ctrl *gomock.Controller) *MockMachineRepository {
	mock := &MockMachineRepository{ctrl: ctrl}
	mock.recorder = &MockMachineRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineRepository) EXPECT() *MockMachineRepositoryMockRecorder {
	return m.recorder
}

// CreateMachine mocks base method.
func (m *MockMachineRepository) CreateMachine(machine *types.Machine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", machine)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockMachineRepositoryMockRecorder) CreateMachine(machine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockMachineRepository)(nil).CreateMachine), machine)
}

// DeleteMachine mocks base method.
func (m *MockMachineRepository) DeleteMachine(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMachine", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMachine indicates an expected call of DeleteMachine.
func (mr *MockMachineRepositoryMockRecorder) DeleteMachine(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMachine", reflect.TypeOf((*MockMachineRepository)(nil).DeleteMachine), id)
}

// GetMachineById mocks base method.
func (m *MockMachineRepository) GetMachineById(id string) (*types.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineById", id)
	ret0, _ := ret[0].(*types.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineById indicates an expected call of GetMachineById.
func (mr *MockMachineRepositoryMockRecorder) GetMachineById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineById", reflect.TypeOf((*MockMachineRepository)(nil).GetMachineById), id)
}

// GetMachines mocks base method.
func (m *MockMachineRepository) GetMachines(id string) ([]types.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachines", id)
	ret0, _ := ret[0].([]types.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachines indicates an expected call of GetMachines.
func (mr *MockMachineRepositoryMockRecorder) GetMachines(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachines", reflect.TypeOf((*MockMachineRepository)(nil).GetMachines), id)
}

// UpdateMachine mocks base method.
func (m *MockMachineRepository) UpdateMachine(id string, machine *types.Machine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMachine", id, machine)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMachine indicates an expected call of UpdateMachine.
func (mr *MockMachineRepositoryMockRecorder) UpdateMachine(id, machine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMachine", reflect.TypeOf((*MockMachineRepository)(nil).UpdateMachine), id, machine)
}