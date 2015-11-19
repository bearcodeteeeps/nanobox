// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/nanobox-io/nanobox/util/file/hosts (interfaces: Host)

package mock_hosts

import gomock "github.com/golang/mock/gomock"

// Mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *_MockHostRecorder
}

// Recorder for MockHost (not exported)
type _MockHostRecorder struct {
	mock *MockHost
}

func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &_MockHostRecorder{mock}
	return mock
}

func (_m *MockHost) EXPECT() *_MockHostRecorder {
	return _m.recorder
}

func (_m *MockHost) AddDomain() {
	_m.ctrl.Call(_m, "AddDomain")
}

func (_mr *_MockHostRecorder) AddDomain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddDomain")
}

func (_m *MockHost) HasDomain() bool {
	ret := _m.ctrl.Call(_m, "HasDomain")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockHostRecorder) HasDomain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasDomain")
}

func (_m *MockHost) RemoveDomain() {
	_m.ctrl.Call(_m, "RemoveDomain")
}

func (_mr *_MockHostRecorder) RemoveDomain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveDomain")
}