// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/form-api/pkg/server/grpc (interfaces: ProtectedStorage)

// Package grpc_test is a generated GoMock package.
package grpc_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	domain "github.com/kamilsk/form-api/pkg/domain"
	query "github.com/kamilsk/form-api/pkg/storage/query"
	types "github.com/kamilsk/form-api/pkg/storage/types"
)

// MockProtectedStorage is a mock of ProtectedStorage interface
type MockProtectedStorage struct {
	ctrl     *gomock.Controller
	recorder *MockProtectedStorageMockRecorder
}

// MockProtectedStorageMockRecorder is the mock recorder for MockProtectedStorage
type MockProtectedStorageMockRecorder struct {
	mock *MockProtectedStorage
}

// NewMockProtectedStorage creates a new mock instance
func NewMockProtectedStorage(ctrl *gomock.Controller) *MockProtectedStorage {
	mock := &MockProtectedStorage{ctrl: ctrl}
	mock.recorder = &MockProtectedStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProtectedStorage) EXPECT() *MockProtectedStorageMockRecorder {
	return m.recorder
}

// CreateSchema mocks base method
func (m *MockProtectedStorage) CreateSchema(arg0 context.Context, arg1 domain.ID, arg2 query.CreateSchema) (types.Schema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchema", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Schema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchema indicates an expected call of CreateSchema
func (mr *MockProtectedStorageMockRecorder) CreateSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchema", reflect.TypeOf((*MockProtectedStorage)(nil).CreateSchema), arg0, arg1, arg2)
}

// CreateTemplate mocks base method
func (m *MockProtectedStorage) CreateTemplate(arg0 context.Context, arg1 domain.ID, arg2 query.CreateTemplate) (types.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemplate indicates an expected call of CreateTemplate
func (mr *MockProtectedStorageMockRecorder) CreateTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplate", reflect.TypeOf((*MockProtectedStorage)(nil).CreateTemplate), arg0, arg1, arg2)
}

// DeleteSchema mocks base method
func (m *MockProtectedStorage) DeleteSchema(arg0 context.Context, arg1 domain.ID, arg2 query.DeleteSchema) (types.Schema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSchema", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Schema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSchema indicates an expected call of DeleteSchema
func (mr *MockProtectedStorageMockRecorder) DeleteSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSchema", reflect.TypeOf((*MockProtectedStorage)(nil).DeleteSchema), arg0, arg1, arg2)
}

// DeleteTemplate mocks base method
func (m *MockProtectedStorage) DeleteTemplate(arg0 context.Context, arg1 domain.ID, arg2 query.DeleteTemplate) (types.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTemplate indicates an expected call of DeleteTemplate
func (mr *MockProtectedStorageMockRecorder) DeleteTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockProtectedStorage)(nil).DeleteTemplate), arg0, arg1, arg2)
}

// ReadInputByFilter mocks base method
func (m *MockProtectedStorage) ReadInputByFilter(arg0 context.Context, arg1 domain.ID, arg2 query.InputFilter) ([]types.Input, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInputByFilter", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.Input)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInputByFilter indicates an expected call of ReadInputByFilter
func (mr *MockProtectedStorageMockRecorder) ReadInputByFilter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInputByFilter", reflect.TypeOf((*MockProtectedStorage)(nil).ReadInputByFilter), arg0, arg1, arg2)
}

// ReadInputByID mocks base method
func (m *MockProtectedStorage) ReadInputByID(arg0 context.Context, arg1, arg2 domain.ID) (types.Input, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInputByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Input)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInputByID indicates an expected call of ReadInputByID
func (mr *MockProtectedStorageMockRecorder) ReadInputByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInputByID", reflect.TypeOf((*MockProtectedStorage)(nil).ReadInputByID), arg0, arg1, arg2)
}

// ReadSchema mocks base method
func (m *MockProtectedStorage) ReadSchema(arg0 context.Context, arg1 domain.ID, arg2 query.ReadSchema) (types.Schema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSchema", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Schema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSchema indicates an expected call of ReadSchema
func (mr *MockProtectedStorageMockRecorder) ReadSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSchema", reflect.TypeOf((*MockProtectedStorage)(nil).ReadSchema), arg0, arg1, arg2)
}

// ReadTemplate mocks base method
func (m *MockProtectedStorage) ReadTemplate(arg0 context.Context, arg1 domain.ID, arg2 query.ReadTemplate) (types.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTemplate indicates an expected call of ReadTemplate
func (mr *MockProtectedStorageMockRecorder) ReadTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTemplate", reflect.TypeOf((*MockProtectedStorage)(nil).ReadTemplate), arg0, arg1, arg2)
}

// UpdateSchema mocks base method
func (m *MockProtectedStorage) UpdateSchema(arg0 context.Context, arg1 domain.ID, arg2 query.UpdateSchema) (types.Schema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchema", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Schema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchema indicates an expected call of UpdateSchema
func (mr *MockProtectedStorageMockRecorder) UpdateSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchema", reflect.TypeOf((*MockProtectedStorage)(nil).UpdateSchema), arg0, arg1, arg2)
}

// UpdateTemplate mocks base method
func (m *MockProtectedStorage) UpdateTemplate(arg0 context.Context, arg1 domain.ID, arg2 query.UpdateTemplate) (types.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTemplate indicates an expected call of UpdateTemplate
func (mr *MockProtectedStorageMockRecorder) UpdateTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockProtectedStorage)(nil).UpdateTemplate), arg0, arg1, arg2)
}
