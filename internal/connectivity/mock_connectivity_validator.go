// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package connectivity is a generated GoMock package.
package connectivity

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// GetHostValidInterfaces mocks base method
func (m *MockValidator) GetHostValidInterfaces(host *models.Host) ([]*models.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostValidInterfaces", host)
	ret0, _ := ret[0].([]*models.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostValidInterfaces indicates an expected call of GetHostValidInterfaces
func (mr *MockValidatorMockRecorder) GetHostValidInterfaces(host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostValidInterfaces", reflect.TypeOf((*MockValidator)(nil).GetHostValidInterfaces), host)
}
