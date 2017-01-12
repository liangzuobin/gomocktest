// Automatically generated by MockGen. DO NOT EDIT!
// Source: mock.go

package mock_firsttry

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Face interface
type MockFace struct {
	ctrl     *gomock.Controller
	recorder *_MockFaceRecorder
}

// Recorder for MockFace (not exported)
type _MockFaceRecorder struct {
	mock *MockFace
}

func NewMockFace(ctrl *gomock.Controller) *MockFace {
	mock := &MockFace{ctrl: ctrl}
	mock.recorder = &_MockFaceRecorder{mock}
	return mock
}

func (_m *MockFace) EXPECT() *_MockFaceRecorder {
	return _m.recorder
}

func (_m *MockFace) Bar() (string, error) {
	ret := _m.ctrl.Call(_m, "Bar")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFaceRecorder) Bar() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bar")
}

func (_m *MockFace) Baz() (string, error) {
	ret := _m.ctrl.Call(_m, "Baz")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFaceRecorder) Baz() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Baz")
}
