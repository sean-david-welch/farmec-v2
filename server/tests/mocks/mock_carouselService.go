// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/carouselService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sean-david-welch/farmec-v2/server/types"
)

// MockCarouselService is a mock of CarouselService interface.
type MockCarouselService struct {
	ctrl     *gomock.Controller
	recorder *MockCarouselServiceMockRecorder
}

// MockCarouselServiceMockRecorder is the mock recorder for MockCarouselService.
type MockCarouselServiceMockRecorder struct {
	mock *MockCarouselService
}

// NewMockCarouselService creates a new mock instance.
func NewMockCarouselService(ctrl *gomock.Controller) *MockCarouselService {
	mock := &MockCarouselService{ctrl: ctrl}
	mock.recorder = &MockCarouselServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarouselService) EXPECT() *MockCarouselServiceMockRecorder {
	return m.recorder
}

// CreateCarousel mocks base method.
func (m *MockCarouselService) CreateCarousel(carousel *types.Carousel) (*types.ModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCarousel", carousel)
	ret0, _ := ret[0].(*types.ModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCarousel indicates an expected call of CreateCarousel.
func (mr *MockCarouselServiceMockRecorder) CreateCarousel(carousel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCarousel", reflect.TypeOf((*MockCarouselService)(nil).CreateCarousel), carousel)
}

// DeleteCarousel mocks base method.
func (m *MockCarouselService) DeleteCarousel(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCarousel", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCarousel indicates an expected call of DeleteCarousel.
func (mr *MockCarouselServiceMockRecorder) DeleteCarousel(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCarousel", reflect.TypeOf((*MockCarouselService)(nil).DeleteCarousel), id)
}

// GetCarousels mocks base method.
func (m *MockCarouselService) GetCarousels() ([]types.Carousel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCarousels")
	ret0, _ := ret[0].([]types.Carousel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCarousels indicates an expected call of GetCarousels.
func (mr *MockCarouselServiceMockRecorder) GetCarousels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCarousels", reflect.TypeOf((*MockCarouselService)(nil).GetCarousels))
}

// UpdateCarousel mocks base method.
func (m *MockCarouselService) UpdateCarousel(id string, carousel *types.Carousel) (*types.ModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCarousel", id, carousel)
	ret0, _ := ret[0].(*types.ModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCarousel indicates an expected call of UpdateCarousel.
func (mr *MockCarouselServiceMockRecorder) UpdateCarousel(id, carousel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCarousel", reflect.TypeOf((*MockCarouselService)(nil).UpdateCarousel), id, carousel)
}