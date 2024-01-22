// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/videoRepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sean-david-welch/farmec-v2/server/types"
)

// MockVideoRepository is a mock of VideoRepository interface.
type MockVideoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVideoRepositoryMockRecorder
}

// MockVideoRepositoryMockRecorder is the mock recorder for MockVideoRepository.
type MockVideoRepositoryMockRecorder struct {
	mock *MockVideoRepository
}

// NewMockVideoRepository creates a new mock instance.
func NewMockVideoRepository(ctrl *gomock.Controller) *MockVideoRepository {
	mock := &MockVideoRepository{ctrl: ctrl}
	mock.recorder = &MockVideoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVideoRepository) EXPECT() *MockVideoRepositoryMockRecorder {
	return m.recorder
}

// CreateVideo mocks base method.
func (m *MockVideoRepository) CreateVideo(video *types.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVideo", video)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVideo indicates an expected call of CreateVideo.
func (mr *MockVideoRepositoryMockRecorder) CreateVideo(video interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVideo", reflect.TypeOf((*MockVideoRepository)(nil).CreateVideo), video)
}

// DeleteVideo mocks base method.
func (m *MockVideoRepository) DeleteVideo(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVideo", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVideo indicates an expected call of DeleteVideo.
func (mr *MockVideoRepositoryMockRecorder) DeleteVideo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVideo", reflect.TypeOf((*MockVideoRepository)(nil).DeleteVideo), id)
}

// GetVideos mocks base method.
func (m *MockVideoRepository) GetVideos(id string) ([]types.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideos", id)
	ret0, _ := ret[0].([]types.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideos indicates an expected call of GetVideos.
func (mr *MockVideoRepositoryMockRecorder) GetVideos(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideos", reflect.TypeOf((*MockVideoRepository)(nil).GetVideos), id)
}

// UpdateVideo mocks base method.
func (m *MockVideoRepository) UpdateVideo(id string, video *types.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVideo", id, video)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVideo indicates an expected call of UpdateVideo.
func (mr *MockVideoRepositoryMockRecorder) UpdateVideo(id, video interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVideo", reflect.TypeOf((*MockVideoRepository)(nil).UpdateVideo), id, video)
}
