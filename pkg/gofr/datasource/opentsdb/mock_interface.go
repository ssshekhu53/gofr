// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock_interface.go -package=opentsdb
//

// Package opentsdb is a generated GoMock package.
package opentsdb

import (
	context "context"
	net "net"
	http "net/http"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockhttpClient is a mock of httpClient interface.
type MockhttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockhttpClientMockRecorder
}

// MockhttpClientMockRecorder is the mock recorder for MockhttpClient.
type MockhttpClientMockRecorder struct {
	mock *MockhttpClient
}

// NewMockhttpClient creates a new mock instance.
func NewMockhttpClient(ctrl *gomock.Controller) *MockhttpClient {
	mock := &MockhttpClient{ctrl: ctrl}
	mock.recorder = &MockhttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockhttpClient) EXPECT() *MockhttpClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockhttpClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockhttpClientMockRecorder) Do(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockhttpClient)(nil).Do), req)
}

// Mockconnection is a mock of connection interface.
type Mockconnection struct {
	ctrl     *gomock.Controller
	recorder *MockconnectionMockRecorder
}

// MockconnectionMockRecorder is the mock recorder for Mockconnection.
type MockconnectionMockRecorder struct {
	mock *Mockconnection
}

// NewMockconnection creates a new mock instance.
func NewMockconnection(ctrl *gomock.Controller) *Mockconnection {
	mock := &Mockconnection{ctrl: ctrl}
	mock.recorder = &MockconnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockconnection) EXPECT() *MockconnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *Mockconnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockconnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Mockconnection)(nil).Close))
}

// LocalAddr mocks base method.
func (m *Mockconnection) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockconnectionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*Mockconnection)(nil).LocalAddr))
}

// Read mocks base method.
func (m *Mockconnection) Read(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockconnectionMockRecorder) Read(b any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*Mockconnection)(nil).Read), b)
}

// RemoteAddr mocks base method.
func (m *Mockconnection) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockconnectionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*Mockconnection)(nil).RemoteAddr))
}

// SetDeadline mocks base method.
func (m *Mockconnection) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockconnectionMockRecorder) SetDeadline(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*Mockconnection)(nil).SetDeadline), t)
}

// SetReadDeadline mocks base method.
func (m *Mockconnection) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockconnectionMockRecorder) SetReadDeadline(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*Mockconnection)(nil).SetReadDeadline), t)
}

// SetWriteDeadline mocks base method.
func (m *Mockconnection) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockconnectionMockRecorder) SetWriteDeadline(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*Mockconnection)(nil).SetWriteDeadline), t)
}

// Write mocks base method.
func (m *Mockconnection) Write(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockconnectionMockRecorder) Write(b any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*Mockconnection)(nil).Write), b)
}

// Mockresponse is a mock of response interface.
type Mockresponse struct {
	ctrl     *gomock.Controller
	recorder *MockresponseMockRecorder
}

// MockresponseMockRecorder is the mock recorder for Mockresponse.
type MockresponseMockRecorder struct {
	mock *Mockresponse
}

// NewMockresponse creates a new mock instance.
func NewMockresponse(ctrl *gomock.Controller) *Mockresponse {
	mock := &Mockresponse{ctrl: ctrl}
	mock.recorder = &MockresponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockresponse) EXPECT() *MockresponseMockRecorder {
	return m.recorder
}

// getCustomParser mocks base method.
func (m *Mockresponse) getCustomParser(arg0 Logger) func([]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getCustomParser", arg0)
	ret0, _ := ret[0].(func([]byte) error)
	return ret0
}

// getCustomParser indicates an expected call of getCustomParser.
func (mr *MockresponseMockRecorder) getCustomParser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getCustomParser", reflect.TypeOf((*Mockresponse)(nil).getCustomParser), arg0)
}

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

// Debug mocks base method.
func (m *MockLogger) Debug(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Fatal mocks base method.
func (m *MockLogger) Fatal(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder) Fatal(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), args...)
}

// Log mocks base method.
func (m *MockLogger) Log(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log.
func (mr *MockLoggerMockRecorder) Log(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), args...)
}

// Logf mocks base method.
func (m *MockLogger) Logf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockLoggerMockRecorder) Logf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockLogger)(nil).Logf), varargs...)
}

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// NewHistogram mocks base method.
func (m *MockMetrics) NewHistogram(name, desc string, buckets ...float64) {
	m.ctrl.T.Helper()
	varargs := []any{name, desc}
	for _, a := range buckets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NewHistogram", varargs...)
}

// NewHistogram indicates an expected call of NewHistogram.
func (mr *MockMetricsMockRecorder) NewHistogram(name, desc any, buckets ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, desc}, buckets...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistogram", reflect.TypeOf((*MockMetrics)(nil).NewHistogram), varargs...)
}

// RecordHistogram mocks base method.
func (m *MockMetrics) RecordHistogram(ctx context.Context, name string, value float64, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name, value}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RecordHistogram", varargs...)
}

// RecordHistogram indicates an expected call of RecordHistogram.
func (mr *MockMetricsMockRecorder) RecordHistogram(ctx, name, value any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name, value}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordHistogram", reflect.TypeOf((*MockMetrics)(nil).RecordHistogram), varargs...)
}