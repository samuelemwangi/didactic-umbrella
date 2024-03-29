// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructure\queueing\kafka.go

// Package queueing_mock is a generated GoMock package.
package queueing_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKafkaConsumer is a mock of KafkaConsumer interface.
type MockKafkaConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaConsumerMockRecorder
}

// MockKafkaConsumerMockRecorder is the mock recorder for MockKafkaConsumer.
type MockKafkaConsumerMockRecorder struct {
	mock *MockKafkaConsumer
}

// NewMockKafkaConsumer creates a new mock instance.
func NewMockKafkaConsumer(ctrl *gomock.Controller) *MockKafkaConsumer {
	mock := &MockKafkaConsumer{ctrl: ctrl}
	mock.recorder = &MockKafkaConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaConsumer) EXPECT() *MockKafkaConsumerMockRecorder {
	return m.recorder
}

// ConsumeMessage mocks base method.
func (m *MockKafkaConsumer) ConsumeMessage(arg0 string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeMessage", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsumeMessage indicates an expected call of ConsumeMessage.
func (mr *MockKafkaConsumerMockRecorder) ConsumeMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeMessage", reflect.TypeOf((*MockKafkaConsumer)(nil).ConsumeMessage), arg0)
}
