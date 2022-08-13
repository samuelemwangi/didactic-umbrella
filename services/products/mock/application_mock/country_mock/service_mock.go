// Code generated by MockGen. DO NOT EDIT.
// Source: application\country\service.go

// Package country_mock is a generated GoMock package.
package country_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	country "github.com/samuelemwangi/jumia-mds-test/services/products/application/country"
	errorhelper "github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
)

// MockCountryService is a mock of CountryService interface.
type MockCountryService struct {
	ctrl     *gomock.Controller
	recorder *MockCountryServiceMockRecorder
}

// MockCountryServiceMockRecorder is the mock recorder for MockCountryService.
type MockCountryServiceMockRecorder struct {
	mock *MockCountryService
}

// NewMockCountryService creates a new mock instance.
func NewMockCountryService(ctrl *gomock.Controller) *MockCountryService {
	mock := &MockCountryService{ctrl: ctrl}
	mock.recorder = &MockCountryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCountryService) EXPECT() *MockCountryServiceMockRecorder {
	return m.recorder
}

// GetCountries mocks base method.
func (m *MockCountryService) GetCountries() (*country.CountriesResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountries")
	ret0, _ := ret[0].(*country.CountriesResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// GetCountries indicates an expected call of GetCountries.
func (mr *MockCountryServiceMockRecorder) GetCountries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountries", reflect.TypeOf((*MockCountryService)(nil).GetCountries))
}

// SaveCountry mocks base method.
func (m *MockCountryService) SaveCountry(arg0 *country.CountryRequestDTO) (*country.CountryResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCountry", arg0)
	ret0, _ := ret[0].(*country.CountryResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// SaveCountry indicates an expected call of SaveCountry.
func (mr *MockCountryServiceMockRecorder) SaveCountry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCountry", reflect.TypeOf((*MockCountryService)(nil).SaveCountry), arg0)
}
