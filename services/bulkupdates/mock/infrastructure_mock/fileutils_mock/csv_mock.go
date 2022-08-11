// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructure\fileutils\csv.go

// Package fileutils_mock is a generated GoMock package.
package fileutils_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCSVReader is a mock of CSVReader interface.
type MockCSVReader struct {
	ctrl     *gomock.Controller
	recorder *MockCSVReaderMockRecorder
}

// MockCSVReaderMockRecorder is the mock recorder for MockCSVReader.
type MockCSVReaderMockRecorder struct {
	mock *MockCSVReader
}

// NewMockCSVReader creates a new mock instance.
func NewMockCSVReader(ctrl *gomock.Controller) *MockCSVReader {
	mock := &MockCSVReader{ctrl: ctrl}
	mock.recorder = &MockCSVReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCSVReader) EXPECT() *MockCSVReaderMockRecorder {
	return m.recorder
}

// ReadFile mocks base method.
func (m *MockCSVReader) ReadFile(arg0 string) ([][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockCSVReaderMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockCSVReader)(nil).ReadFile), arg0)
}
