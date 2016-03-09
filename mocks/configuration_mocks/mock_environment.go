// Automatically generated by MockGen. DO NOT EDIT!
// Source: configuration/environment.go

package mock_configuration

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of EnvironmentService interface
type MockEnvironmentService struct {
	ctrl     *gomock.Controller
	recorder *_MockEnvironmentServiceRecorder
}

// Recorder for MockEnvironmentService (not exported)
type _MockEnvironmentServiceRecorder struct {
	mock *MockEnvironmentService
}

func NewMockEnvironmentService(ctrl *gomock.Controller) *MockEnvironmentService {
	mock := &MockEnvironmentService{ctrl: ctrl}
	mock.recorder = &_MockEnvironmentServiceRecorder{mock}
	return mock
}

func (_m *MockEnvironmentService) EXPECT() *_MockEnvironmentServiceRecorder {
	return _m.recorder
}

func (_m *MockEnvironmentService) Get(id string) (*Environment, error) {
	ret := _m.ctrl.Call(_m, "Get", id)
	ret0, _ := ret[0].(*Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEnvironmentServiceRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockEnvironmentService) List() ([]Environment, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEnvironmentServiceRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockEnvironmentService) ListWithOptions(options map[string]string) ([]Environment, error) {
	ret := _m.ctrl.Call(_m, "ListWithOptions", options)
	ret0, _ := ret[0].([]Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEnvironmentServiceRecorder) ListWithOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListWithOptions", arg0)
}

func (_m *MockEnvironmentService) Create(environment Environment) (*Environment, error) {
	ret := _m.ctrl.Call(_m, "Create", environment)
	ret0, _ := ret[0].(*Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEnvironmentServiceRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockEnvironmentService) Update(id string, environment Environment) (*Environment, error) {
	ret := _m.ctrl.Call(_m, "Update", id, environment)
	ret0, _ := ret[0].(*Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEnvironmentServiceRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1)
}

func (_m *MockEnvironmentService) Delete(id string) (bool, error) {
	ret := _m.ctrl.Call(_m, "Delete", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEnvironmentServiceRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}
