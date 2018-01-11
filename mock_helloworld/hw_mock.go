// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hg2c/hellogrpc/helloworld (interfaces: GreeterClient)

// Package mock_helloworld is a generated GoMock package.
package mock_helloworld

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	helloworld "github.com/hg2c/hellogrpc/helloworld"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGreeterClient is a mock of GreeterClient interface
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterClientMockRecorder
}

// MockGreeterClientMockRecorder is the mock recorder for MockGreeterClient
type MockGreeterClientMockRecorder struct {
	mock *MockGreeterClient
}

// NewMockGreeterClient creates a new mock instance
func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &MockGreeterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGreeterClient) EXPECT() *MockGreeterClientMockRecorder {
	return m.recorder
}

// SayHello mocks base method
func (m *MockGreeterClient) SayHello(arg0 context.Context, arg1 *helloworld.HelloRequest, arg2 ...grpc.CallOption) (*helloworld.HelloReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*helloworld.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello
func (mr *MockGreeterClientMockRecorder) SayHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockGreeterClient)(nil).SayHello), varargs...)
}
