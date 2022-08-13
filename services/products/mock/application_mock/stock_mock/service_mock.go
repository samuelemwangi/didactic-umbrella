// Code generated by MockGen. DO NOT EDIT.
// Source: application\stock\service.go

// Package stock_mock is a generated GoMock package.
package stock_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errorhelper "github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	stock "github.com/samuelemwangi/jumia-mds-test/services/products/application/stock"
)

// MockStockService is a mock of StockService interface.
type MockStockService struct {
	ctrl     *gomock.Controller
	recorder *MockStockServiceMockRecorder
}

// MockStockServiceMockRecorder is the mock recorder for MockStockService.
type MockStockServiceMockRecorder struct {
	mock *MockStockService
}

// NewMockStockService creates a new mock instance.
func NewMockStockService(ctrl *gomock.Controller) *MockStockService {
	mock := &MockStockService{ctrl: ctrl}
	mock.recorder = &MockStockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStockService) EXPECT() *MockStockServiceMockRecorder {
	return m.recorder
}

// ConsumeStock mocks base method.
func (m *MockStockService) ConsumeStock(arg0 *stock.ConsumeStockRequestDTO) (*stock.ConsumeStockResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeStock", arg0)
	ret0, _ := ret[0].(*stock.ConsumeStockResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// ConsumeStock indicates an expected call of ConsumeStock.
func (mr *MockStockServiceMockRecorder) ConsumeStock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeStock", reflect.TypeOf((*MockStockService)(nil).ConsumeStock), arg0)
}