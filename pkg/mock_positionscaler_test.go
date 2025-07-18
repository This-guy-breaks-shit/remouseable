// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kevinconway/remouseable/pkg (interfaces: PositionScaler)

// Package remouseable is a generated GoMock package.
package remouseable

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPositionScaler is a mock of PositionScaler interface
type MockPositionScaler struct {
	ctrl     *gomock.Controller
	recorder *MockPositionScalerMockRecorder
}

// MockPositionScalerMockRecorder is the mock recorder for MockPositionScaler
type MockPositionScalerMockRecorder struct {
	mock *MockPositionScaler
}

// NewMockPositionScaler creates a new mock instance
func NewMockPositionScaler(ctrl *gomock.Controller) *MockPositionScaler {
	mock := &MockPositionScaler{ctrl: ctrl}
	mock.recorder = &MockPositionScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPositionScaler) EXPECT() *MockPositionScalerMockRecorder {
	return m.recorder
}

// ScalePosition mocks base method
func (m *MockPositionScaler) ScalePosition(arg0, arg1 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScalePosition", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// ScalePosition indicates an expected call of ScalePosition
func (mr *MockPositionScalerMockRecorder) ScalePosition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScalePosition", reflect.TypeOf((*MockPositionScaler)(nil).ScalePosition), arg0, arg1)
}
