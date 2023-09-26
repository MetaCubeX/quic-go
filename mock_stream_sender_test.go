// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go (interfaces: StreamSender)
//
// Generated by this command:
//
//	mockgen -build_flags=-tags=gomock -package quic -self_package github.com/metacubex/quic-go -destination mock_stream_sender_test.go github.com/metacubex/quic-go StreamSender
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/metacubex/quic-go/internal/protocol"
	wire "github.com/metacubex/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockStreamSender is a mock of StreamSender interface.
type MockStreamSender struct {
	ctrl     *gomock.Controller
	recorder *MockStreamSenderMockRecorder
}

// MockStreamSenderMockRecorder is the mock recorder for MockStreamSender.
type MockStreamSenderMockRecorder struct {
	mock *MockStreamSender
}

// NewMockStreamSender creates a new mock instance.
func NewMockStreamSender(ctrl *gomock.Controller) *MockStreamSender {
	mock := &MockStreamSender{ctrl: ctrl}
	mock.recorder = &MockStreamSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamSender) EXPECT() *MockStreamSenderMockRecorder {
	return m.recorder
}

// onHasStreamData mocks base method.
func (m *MockStreamSender) onHasStreamData(arg0 protocol.StreamID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "onHasStreamData", arg0)
}

// onHasStreamData indicates an expected call of onHasStreamData.
func (mr *MockStreamSenderMockRecorder) onHasStreamData(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onHasStreamData", reflect.TypeOf((*MockStreamSender)(nil).onHasStreamData), arg0)
}

// onStreamCompleted mocks base method.
func (m *MockStreamSender) onStreamCompleted(arg0 protocol.StreamID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "onStreamCompleted", arg0)
}

// onStreamCompleted indicates an expected call of onStreamCompleted.
func (mr *MockStreamSenderMockRecorder) onStreamCompleted(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onStreamCompleted", reflect.TypeOf((*MockStreamSender)(nil).onStreamCompleted), arg0)
}

// queueControlFrame mocks base method.
func (m *MockStreamSender) queueControlFrame(arg0 wire.Frame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "queueControlFrame", arg0)
}

// queueControlFrame indicates an expected call of queueControlFrame.
func (mr *MockStreamSenderMockRecorder) queueControlFrame(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "queueControlFrame", reflect.TypeOf((*MockStreamSender)(nil).queueControlFrame), arg0)
}
