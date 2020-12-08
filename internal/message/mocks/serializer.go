// Code generated by MockGen. DO NOT EDIT.
// Source: internal/message/serializer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	gcloud_grpc_boilerplate "github.com/micromaniacs/gcloud-grpc-boilerplate"
	models "github.com/micromaniacs/gcloud-grpc-boilerplate/internal/models"
	reflect "reflect"
)

// MockSerializer is a mock of Serializer interface
type MockSerializer struct {
	ctrl     *gomock.Controller
	recorder *MockSerializerMockRecorder
}

// MockSerializerMockRecorder is the mock recorder for MockSerializer
type MockSerializerMockRecorder struct {
	mock *MockSerializer
}

// NewMockSerializer creates a new mock instance
func NewMockSerializer(ctrl *gomock.Controller) *MockSerializer {
	mock := &MockSerializer{ctrl: ctrl}
	mock.recorder = &MockSerializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSerializer) EXPECT() *MockSerializerMockRecorder {
	return m.recorder
}

// ToMessage mocks base method
func (m *MockSerializer) ToMessage(arg0 *models.Message) *gcloud_grpc_boilerplate.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMessage", arg0)
	ret0, _ := ret[0].(*gcloud_grpc_boilerplate.Message)
	return ret0
}

// ToMessage indicates an expected call of ToMessage
func (mr *MockSerializerMockRecorder) ToMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMessage", reflect.TypeOf((*MockSerializer)(nil).ToMessage), arg0)
}

// ToMessages mocks base method
func (m *MockSerializer) ToMessages(arg0 []*models.Message) []*gcloud_grpc_boilerplate.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMessages", arg0)
	ret0, _ := ret[0].([]*gcloud_grpc_boilerplate.Message)
	return ret0
}

// ToMessages indicates an expected call of ToMessages
func (mr *MockSerializerMockRecorder) ToMessages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMessages", reflect.TypeOf((*MockSerializer)(nil).ToMessages), arg0)
}