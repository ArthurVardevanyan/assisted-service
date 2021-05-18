// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/pkg/ocm (interfaces: OCMAccountsMgmt)

// Package ocm is a generated GoMock package.
package ocm

import (
	context "context"
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	reflect "reflect"
)

// MockOCMAccountsMgmt is a mock of OCMAccountsMgmt interface
type MockOCMAccountsMgmt struct {
	ctrl     *gomock.Controller
	recorder *MockOCMAccountsMgmtMockRecorder
}

// MockOCMAccountsMgmtMockRecorder is the mock recorder for MockOCMAccountsMgmt
type MockOCMAccountsMgmtMockRecorder struct {
	mock *MockOCMAccountsMgmt
}

// NewMockOCMAccountsMgmt creates a new mock instance
func NewMockOCMAccountsMgmt(ctrl *gomock.Controller) *MockOCMAccountsMgmt {
	mock := &MockOCMAccountsMgmt{ctrl: ctrl}
	mock.recorder = &MockOCMAccountsMgmtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOCMAccountsMgmt) EXPECT() *MockOCMAccountsMgmtMockRecorder {
	return m.recorder
}

// CreateSubscription mocks base method
func (m *MockOCMAccountsMgmt) CreateSubscription(arg0 context.Context, arg1 strfmt.UUID, arg2 string) (*v1.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription
func (mr *MockOCMAccountsMgmtMockRecorder) CreateSubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).CreateSubscription), arg0, arg1, arg2)
}

// DeleteSubscription mocks base method
func (m *MockOCMAccountsMgmt) DeleteSubscription(arg0 context.Context, arg1 strfmt.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription
func (mr *MockOCMAccountsMgmtMockRecorder) DeleteSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).DeleteSubscription), arg0, arg1)
}

// GetSubscription mocks base method
func (m *MockOCMAccountsMgmt) GetSubscription(arg0 context.Context, arg1 strfmt.UUID) (*v1.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscription", arg0, arg1)
	ret0, _ := ret[0].(*v1.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription
func (mr *MockOCMAccountsMgmtMockRecorder) GetSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscription", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).GetSubscription), arg0, arg1)
}

// UpdateSubscriptionConsoleUrl mocks base method
func (m *MockOCMAccountsMgmt) UpdateSubscriptionConsoleUrl(arg0 context.Context, arg1 strfmt.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriptionConsoleUrl", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscriptionConsoleUrl indicates an expected call of UpdateSubscriptionConsoleUrl
func (mr *MockOCMAccountsMgmtMockRecorder) UpdateSubscriptionConsoleUrl(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriptionConsoleUrl", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).UpdateSubscriptionConsoleUrl), arg0, arg1, arg2)
}

// UpdateSubscriptionDisplayName mocks base method
func (m *MockOCMAccountsMgmt) UpdateSubscriptionDisplayName(arg0 context.Context, arg1 strfmt.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriptionDisplayName", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscriptionDisplayName indicates an expected call of UpdateSubscriptionDisplayName
func (mr *MockOCMAccountsMgmtMockRecorder) UpdateSubscriptionDisplayName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriptionDisplayName", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).UpdateSubscriptionDisplayName), arg0, arg1, arg2)
}

// UpdateSubscriptionOpenshiftClusterID mocks base method
func (m *MockOCMAccountsMgmt) UpdateSubscriptionOpenshiftClusterID(arg0 context.Context, arg1, arg2 strfmt.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriptionOpenshiftClusterID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscriptionOpenshiftClusterID indicates an expected call of UpdateSubscriptionOpenshiftClusterID
func (mr *MockOCMAccountsMgmtMockRecorder) UpdateSubscriptionOpenshiftClusterID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriptionOpenshiftClusterID", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).UpdateSubscriptionOpenshiftClusterID), arg0, arg1, arg2)
}

// UpdateSubscriptionStatusActive mocks base method
func (m *MockOCMAccountsMgmt) UpdateSubscriptionStatusActive(arg0 context.Context, arg1 strfmt.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriptionStatusActive", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscriptionStatusActive indicates an expected call of UpdateSubscriptionStatusActive
func (mr *MockOCMAccountsMgmtMockRecorder) UpdateSubscriptionStatusActive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriptionStatusActive", reflect.TypeOf((*MockOCMAccountsMgmt)(nil).UpdateSubscriptionStatusActive), arg0, arg1)
}
