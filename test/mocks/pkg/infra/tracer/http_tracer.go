// Code generated by MockGen. DO NOT EDIT.
// Source: ./http_tracer.go

// Package tracerMock is a generated GoMock package.
package tracerMock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	trace "go.opentelemetry.io/otel/trace"
)

// MockTracer is a mock of Tracer interface.
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
}

// MockTracerMockRecorder is the mock recorder for MockTracer.
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance.
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockTracer) Error(span trace.Span, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", span, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockTracerMockRecorder) Error(span, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockTracer)(nil).Error), span, err)
}

// ShutdownLogger mocks base method.
func (m *MockTracer) ShutdownLogger(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownLogger", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutdownLogger indicates an expected call of ShutdownLogger.
func (mr *MockTracerMockRecorder) ShutdownLogger(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownLogger", reflect.TypeOf((*MockTracer)(nil).ShutdownLogger), ctx)
}

// ShutdownMeter mocks base method.
func (m *MockTracer) ShutdownMeter(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownMeter", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutdownMeter indicates an expected call of ShutdownMeter.
func (mr *MockTracerMockRecorder) ShutdownMeter(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownMeter", reflect.TypeOf((*MockTracer)(nil).ShutdownMeter), ctx)
}

// ShutdownTracer mocks base method.
func (m *MockTracer) ShutdownTracer(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownTracer", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutdownTracer indicates an expected call of ShutdownTracer.
func (mr *MockTracerMockRecorder) ShutdownTracer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownTracer", reflect.TypeOf((*MockTracer)(nil).ShutdownTracer), ctx)
}

// Start mocks base method.
func (m *MockTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, spanName}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(trace.Span)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockTracerMockRecorder) Start(ctx, spanName interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, spanName}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTracer)(nil).Start), varargs...)
}

// StartRoot mocks base method.
func (m *MockTracer) StartRoot(ctx context.Context, request *http.Request, spanName string) (context.Context, trace.Span) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartRoot", ctx, request, spanName)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(trace.Span)
	return ret0, ret1
}

// StartRoot indicates an expected call of StartRoot.
func (mr *MockTracerMockRecorder) StartRoot(ctx, request, spanName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRoot", reflect.TypeOf((*MockTracer)(nil).StartRoot), ctx, request, spanName)
}
