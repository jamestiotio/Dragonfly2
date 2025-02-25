// Code generated by MockGen. DO NOT EDIT.
// Source: client_v2.go
//
// Generated by this command:
//
//	mockgen -destination mocks/client_v2_mock.go -source client_v2.go -package mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	common "d7y.io/api/v2/pkg/apis/common/v2"
	dfdaemon "d7y.io/api/v2/pkg/apis/dfdaemon/v2"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockV2 is a mock of V2 interface.
type MockV2 struct {
	ctrl     *gomock.Controller
	recorder *MockV2MockRecorder
}

// MockV2MockRecorder is the mock recorder for MockV2.
type MockV2MockRecorder struct {
	mock *MockV2
}

// NewMockV2 creates a new mock instance.
func NewMockV2(ctrl *gomock.Controller) *MockV2 {
	mock := &MockV2{ctrl: ctrl}
	mock.recorder = &MockV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockV2) EXPECT() *MockV2MockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockV2) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockV2MockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockV2)(nil).Close))
}

// DeleteTask mocks base method.
func (m *MockV2) DeleteTask(arg0 context.Context, arg1 *dfdaemon.DeleteTaskRequest, arg2 ...grpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockV2MockRecorder) DeleteTask(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockV2)(nil).DeleteTask), varargs...)
}

// DownloadTask mocks base method.
func (m *MockV2) DownloadTask(arg0 context.Context, arg1 *dfdaemon.DownloadTaskRequest, arg2 ...grpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadTask indicates an expected call of DownloadTask.
func (mr *MockV2MockRecorder) DownloadTask(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadTask", reflect.TypeOf((*MockV2)(nil).DownloadTask), varargs...)
}

// StatTask mocks base method.
func (m *MockV2) StatTask(arg0 context.Context, arg1 *dfdaemon.StatTaskRequest, arg2 ...grpc.CallOption) (*common.Task, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatTask", varargs...)
	ret0, _ := ret[0].(*common.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockV2MockRecorder) StatTask(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockV2)(nil).StatTask), varargs...)
}

// SyncPieces mocks base method.
func (m *MockV2) SyncPieces(arg0 context.Context, arg1 ...grpc.CallOption) (dfdaemon.Dfdaemon_SyncPiecesClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncPieces", varargs...)
	ret0, _ := ret[0].(dfdaemon.Dfdaemon_SyncPiecesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncPieces indicates an expected call of SyncPieces.
func (mr *MockV2MockRecorder) SyncPieces(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPieces", reflect.TypeOf((*MockV2)(nil).SyncPieces), varargs...)
}

// UploadTask mocks base method.
func (m *MockV2) UploadTask(arg0 context.Context, arg1 *dfdaemon.UploadTaskRequest, arg2 ...grpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadTask indicates an expected call of UploadTask.
func (mr *MockV2MockRecorder) UploadTask(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTask", reflect.TypeOf((*MockV2)(nil).UploadTask), varargs...)
}
