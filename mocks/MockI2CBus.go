// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kidoman/embd (interfaces: I2CBus)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockI2CBus is a mock of I2CBus interface
type MockI2CBus struct {
	ctrl     *gomock.Controller
	recorder *MockI2CBusMockRecorder
}

// MockI2CBusMockRecorder is the mock recorder for MockI2CBus
type MockI2CBusMockRecorder struct {
	mock *MockI2CBus
}

// NewMockI2CBus creates a new mock instance
func NewMockI2CBus(ctrl *gomock.Controller) *MockI2CBus {
	mock := &MockI2CBus{ctrl: ctrl}
	mock.recorder = &MockI2CBusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockI2CBus) EXPECT() *MockI2CBusMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockI2CBus) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockI2CBusMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockI2CBus)(nil).Close))
}

// ReadByte mocks base method
func (_m *MockI2CBus) ReadByte(_param0 byte) (byte, error) {
	ret := _m.ctrl.Call(_m, "ReadByte", _param0)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByte indicates an expected call of ReadByte
func (_mr *MockI2CBusMockRecorder) ReadByte(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadByte", reflect.TypeOf((*MockI2CBus)(nil).ReadByte), arg0)
}

// ReadByteFromReg mocks base method
func (_m *MockI2CBus) ReadByteFromReg(_param0 byte, _param1 byte) (byte, error) {
	ret := _m.ctrl.Call(_m, "ReadByteFromReg", _param0, _param1)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByteFromReg indicates an expected call of ReadByteFromReg
func (_mr *MockI2CBusMockRecorder) ReadByteFromReg(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadByteFromReg", reflect.TypeOf((*MockI2CBus)(nil).ReadByteFromReg), arg0, arg1)
}

// ReadBytes mocks base method
func (_m *MockI2CBus) ReadBytes(_param0 byte, _param1 int) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "ReadBytes", _param0, _param1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBytes indicates an expected call of ReadBytes
func (_mr *MockI2CBusMockRecorder) ReadBytes(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadBytes", reflect.TypeOf((*MockI2CBus)(nil).ReadBytes), arg0, arg1)
}

// ReadFromReg mocks base method
func (_m *MockI2CBus) ReadFromReg(_param0 byte, _param1 byte, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "ReadFromReg", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFromReg indicates an expected call of ReadFromReg
func (_mr *MockI2CBusMockRecorder) ReadFromReg(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadFromReg", reflect.TypeOf((*MockI2CBus)(nil).ReadFromReg), arg0, arg1, arg2)
}

// ReadWordFromReg mocks base method
func (_m *MockI2CBus) ReadWordFromReg(_param0 byte, _param1 byte) (uint16, error) {
	ret := _m.ctrl.Call(_m, "ReadWordFromReg", _param0, _param1)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWordFromReg indicates an expected call of ReadWordFromReg
func (_mr *MockI2CBusMockRecorder) ReadWordFromReg(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadWordFromReg", reflect.TypeOf((*MockI2CBus)(nil).ReadWordFromReg), arg0, arg1)
}

// WriteByte mocks base method
func (_m *MockI2CBus) WriteByte(_param0 byte, _param1 byte) error {
	ret := _m.ctrl.Call(_m, "WriteByte", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteByte indicates an expected call of WriteByte
func (_mr *MockI2CBusMockRecorder) WriteByte(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteByte", reflect.TypeOf((*MockI2CBus)(nil).WriteByte), arg0, arg1)
}

// WriteByteToReg mocks base method
func (_m *MockI2CBus) WriteByteToReg(_param0 byte, _param1 byte, _param2 byte) error {
	ret := _m.ctrl.Call(_m, "WriteByteToReg", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteByteToReg indicates an expected call of WriteByteToReg
func (_mr *MockI2CBusMockRecorder) WriteByteToReg(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteByteToReg", reflect.TypeOf((*MockI2CBus)(nil).WriteByteToReg), arg0, arg1, arg2)
}

// WriteBytes mocks base method
func (_m *MockI2CBus) WriteBytes(_param0 byte, _param1 []byte) error {
	ret := _m.ctrl.Call(_m, "WriteBytes", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteBytes indicates an expected call of WriteBytes
func (_mr *MockI2CBusMockRecorder) WriteBytes(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteBytes", reflect.TypeOf((*MockI2CBus)(nil).WriteBytes), arg0, arg1)
}

// WriteToReg mocks base method
func (_m *MockI2CBus) WriteToReg(_param0 byte, _param1 byte, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "WriteToReg", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToReg indicates an expected call of WriteToReg
func (_mr *MockI2CBusMockRecorder) WriteToReg(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteToReg", reflect.TypeOf((*MockI2CBus)(nil).WriteToReg), arg0, arg1, arg2)
}

// WriteWordToReg mocks base method
func (_m *MockI2CBus) WriteWordToReg(_param0 byte, _param1 byte, _param2 uint16) error {
	ret := _m.ctrl.Call(_m, "WriteWordToReg", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteWordToReg indicates an expected call of WriteWordToReg
func (_mr *MockI2CBusMockRecorder) WriteWordToReg(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteWordToReg", reflect.TypeOf((*MockI2CBus)(nil).WriteWordToReg), arg0, arg1, arg2)
}
