// Code generated by MockGen. DO NOT EDIT.
// Source: infra/logger/contract.go

// Package mock_logger is a generated GoMock package.
package mock_logger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockLogger) Error(format string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", format)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), format)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Fatal mocks base method.
func (m *MockLogger) Fatal(err interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fatal", err)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder) Fatal(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), err)
}

// Fatalf mocks base method.
func (m *MockLogger) Fatalf(format string, err interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fatalf", format, err)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockLoggerMockRecorder) Fatalf(format, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), format, err)
}

// Infof mocks base method.
func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}
