// Code generated by MockGen. DO NOT EDIT.
// Source: persistence\repositories\uploadmetadata.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/samuelemwangi/jumia-mds-test/services/products/domain"
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

// SaveUploadMetaData mocks base method.
func (m *MockUploadMetadataRepository) SaveUploadMetaData(uploadMetadata *domain.UploadMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUploadMetaData", uploadMetadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUploadMetaData indicates an expected call of SaveUploadMetaData.
func (mr *MockUploadMetadataRepositoryMockRecorder) SaveUploadMetaData(uploadMetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUploadMetaData", reflect.TypeOf((*MockUploadMetadataRepository)(nil).SaveUploadMetaData), uploadMetadata)
}
