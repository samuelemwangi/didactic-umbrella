// Code generated by MockGen. DO NOT EDIT.
// Source: persistence\repositories\uploadmetadata.go

// Package repositories_mock is a generated GoMock package.
package repositories_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

// MockUploadMetadataRepository is a mock of UploadMetadataRepository interface.
type MockUploadMetadataRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUploadMetadataRepositoryMockRecorder
}

// MockUploadMetadataRepositoryMockRecorder is the mock recorder for MockUploadMetadataRepository.
type MockUploadMetadataRepositoryMockRecorder struct {
	mock *MockUploadMetadataRepository
}

// NewMockUploadMetadataRepository creates a new mock instance.
func NewMockUploadMetadataRepository(ctrl *gomock.Controller) *MockUploadMetadataRepository {
	mock := &MockUploadMetadataRepository{ctrl: ctrl}
	mock.recorder = &MockUploadMetadataRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadMetadataRepository) EXPECT() *MockUploadMetadataRepositoryMockRecorder {
	return m.recorder
}

// GetUploadByUploadId mocks base method.
func (m *MockUploadMetadataRepository) GetUploadByUploadId(arg0 string) (*domain.UploadMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUploadByUploadId", arg0)
	ret0, _ := ret[0].(*domain.UploadMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadByUploadId indicates an expected call of GetUploadByUploadId.
func (mr *MockUploadMetadataRepositoryMockRecorder) GetUploadByUploadId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadByUploadId", reflect.TypeOf((*MockUploadMetadataRepository)(nil).GetUploadByUploadId), arg0)
}

// UpdateUploadStatus mocks base method.
func (m *MockUploadMetadataRepository) UpdateUploadStatus(arg0 *domain.UploadMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUploadStatus indicates an expected call of UpdateUploadStatus.
func (mr *MockUploadMetadataRepositoryMockRecorder) UpdateUploadStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadStatus", reflect.TypeOf((*MockUploadMetadataRepository)(nil).UpdateUploadStatus), arg0)
}
