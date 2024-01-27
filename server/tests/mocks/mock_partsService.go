// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/partsService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sean-david-welch/farmec-v2/server/types"
)

// MockPartsService is a mock of PartsService interface.
type MockPartsService struct {
	ctrl     *gomock.Controller
	recorder *MockPartsServiceMockRecorder
}

// MockPartsServiceMockRecorder is the mock recorder for MockPartsService.
type MockPartsServiceMockRecorder struct {
	mock *MockPartsService
}

// NewMockPartsService creates a new mock instance.
func NewMockPartsService(ctrl *gomock.Controller) *MockPartsService {
	mock := &MockPartsService{ctrl: ctrl}
	mock.recorder = &MockPartsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartsService) EXPECT() *MockPartsServiceMockRecorder {
	return m.recorder
}

// CreatePart mocks base method.
func (m *MockPartsService) CreatePart(part *types.Sparepart) (*types.ModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePart", part)
	ret0, _ := ret[0].(*types.ModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePart indicates an expected call of CreatePart.
func (mr *MockPartsServiceMockRecorder) CreatePart(part interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePart", reflect.TypeOf((*MockPartsService)(nil).CreatePart), part)
}

// DeletePart mocks base method.
func (m *MockPartsService) DeletePart(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePart", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePart indicates an expected call of DeletePart.
func (mr *MockPartsServiceMockRecorder) DeletePart(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePart", reflect.TypeOf((*MockPartsService)(nil).DeletePart), id)
}

// GetParts mocks base method.
func (m *MockPartsService) GetParts(id string) ([]types.Sparepart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParts", id)
	ret0, _ := ret[0].([]types.Sparepart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParts indicates an expected call of GetParts.
func (mr *MockPartsServiceMockRecorder) GetParts(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParts", reflect.TypeOf((*MockPartsService)(nil).GetParts), id)
}

// UpdatePart mocks base method.
func (m *MockPartsService) UpdatePart(id string, part *types.Sparepart) (*types.ModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePart", id, part)
	ret0, _ := ret[0].(*types.ModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePart indicates an expected call of UpdatePart.
func (mr *MockPartsServiceMockRecorder) UpdatePart(id, part interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePart", reflect.TypeOf((*MockPartsService)(nil).UpdatePart), id, part)
}