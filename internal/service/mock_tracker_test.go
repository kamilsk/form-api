// Code generated by MockGen. DO NOT EDIT.
// Source: go.octolab.org/ecosystem/forma/internal/service (interfaces: Tracker)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	domain "go.octolab.org/ecosystem/forma/internal/domain"
)

// MockTracker is a mock of Tracker interface
type MockTracker struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerMockRecorder
}

// MockTrackerMockRecorder is the mock recorder for MockTracker
type MockTrackerMockRecorder struct {
	mock *MockTracker
}

// NewMockTracker creates a new mock instance
func NewMockTracker(ctrl *gomock.Controller) *MockTracker {
	mock := &MockTracker{ctrl: ctrl}
	mock.recorder = &MockTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTracker) EXPECT() *MockTrackerMockRecorder {
	return m.recorder
}

// LogInput mocks base method
func (m *MockTracker) LogInput(arg0 context.Context, arg1 domain.InputEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogInput", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogInput indicates an expected call of LogInput
func (mr *MockTrackerMockRecorder) LogInput(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogInput", reflect.TypeOf((*MockTracker)(nil).LogInput), arg0, arg1)
}
