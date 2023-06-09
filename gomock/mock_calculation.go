// Code generated by MockGen. DO NOT EDIT.
// Source: calculation.go

// Package gomock is a generated GoMock package.
package gomock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockcalculation is a mock of calculation interface.
type Mockcalculation struct {
	ctrl     *gomock.Controller
	recorder *MockcalculationMockRecorder
}

// MockcalculationMockRecorder is the mock recorder for Mockcalculation.
type MockcalculationMockRecorder struct {
	mock *Mockcalculation
}

// NewMockcalculation creates a new mock instance.
func NewMockcalculation(ctrl *gomock.Controller) *Mockcalculation {
	mock := &Mockcalculation{ctrl: ctrl}
	mock.recorder = &MockcalculationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcalculation) EXPECT() *MockcalculationMockRecorder {
	return m.recorder
}

// add mocks base method.
func (m *Mockcalculation) add(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "add", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// add indicates an expected call of add.
func (mr *MockcalculationMockRecorder) add(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "add", reflect.TypeOf((*Mockcalculation)(nil).add), a, b)
}

// addFunc mocks base method.
func (m *Mockcalculation) addFunc(a, b int, f func(int, int) int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addFunc", a, b, f)
	ret0, _ := ret[0].(int)
	return ret0
}

// addFunc indicates an expected call of addFunc.
func (mr *MockcalculationMockRecorder) addFunc(a, b, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addFunc", reflect.TypeOf((*Mockcalculation)(nil).addFunc), a, b, f)
}
