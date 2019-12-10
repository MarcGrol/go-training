// Code generated by MockGen. DO NOT EDIT.
// Source: patients.pb.go

// Package patientinfoapi is a generated GoMock package.
package patientinfoapi

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockPatientInfoClient is a mock of PatientInfoClient interface
type MockPatientInfoClient struct {
	ctrl     *gomock.Controller
	recorder *MockPatientInfoClientMockRecorder
}

// MockPatientInfoClientMockRecorder is the mock recorder for MockPatientInfoClient
type MockPatientInfoClientMockRecorder struct {
	mock *MockPatientInfoClient
}

// NewMockPatientInfoClient creates a new mock instance
func NewMockPatientInfoClient(ctrl *gomock.Controller) *MockPatientInfoClient {
	mock := &MockPatientInfoClient{ctrl: ctrl}
	mock.recorder = &MockPatientInfoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPatientInfoClient) EXPECT() *MockPatientInfoClientMockRecorder {
	return m.recorder
}

// GetPatientOnUid mocks base method
func (m *MockPatientInfoClient) GetPatientOnUid(ctx context.Context, in *GetPatientOnUidRequest, opts ...grpc.CallOption) (*GetPatientOnUidReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPatientOnUid", varargs...)
	ret0, _ := ret[0].(*GetPatientOnUidReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatientOnUid indicates an expected call of GetPatientOnUid
func (mr *MockPatientInfoClientMockRecorder) GetPatientOnUid(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatientOnUid", reflect.TypeOf((*MockPatientInfoClient)(nil).GetPatientOnUid), varargs...)
}

// MockPatientInfoServer is a mock of PatientInfoServer interface
type MockPatientInfoServer struct {
	ctrl     *gomock.Controller
	recorder *MockPatientInfoServerMockRecorder
}

// MockPatientInfoServerMockRecorder is the mock recorder for MockPatientInfoServer
type MockPatientInfoServerMockRecorder struct {
	mock *MockPatientInfoServer
}

// NewMockPatientInfoServer creates a new mock instance
func NewMockPatientInfoServer(ctrl *gomock.Controller) *MockPatientInfoServer {
	mock := &MockPatientInfoServer{ctrl: ctrl}
	mock.recorder = &MockPatientInfoServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPatientInfoServer) EXPECT() *MockPatientInfoServerMockRecorder {
	return m.recorder
}

// GetPatientOnUid mocks base method
func (m *MockPatientInfoServer) GetPatientOnUid(arg0 context.Context, arg1 *GetPatientOnUidRequest) (*GetPatientOnUidReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPatientOnUid", arg0, arg1)
	ret0, _ := ret[0].(*GetPatientOnUidReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatientOnUid indicates an expected call of GetPatientOnUid
func (mr *MockPatientInfoServerMockRecorder) GetPatientOnUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatientOnUid", reflect.TypeOf((*MockPatientInfoServer)(nil).GetPatientOnUid), arg0, arg1)
}