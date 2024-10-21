// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go (interfaces: SendStreamI)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/metacubex/quic-go -destination mock_send_stream_internal_test.go github.com/metacubex/quic-go SendStreamI
//

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"
	time "time"

	ackhandler "github.com/metacubex/quic-go/internal/ackhandler"
	protocol "github.com/metacubex/quic-go/internal/protocol"
	qerr "github.com/metacubex/quic-go/internal/qerr"
	wire "github.com/metacubex/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockSendStreamI is a mock of SendStreamI interface.
type MockSendStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockSendStreamIMockRecorder
}

// MockSendStreamIMockRecorder is the mock recorder for MockSendStreamI.
type MockSendStreamIMockRecorder struct {
	mock *MockSendStreamI
}

// NewMockSendStreamI creates a new mock instance.
func NewMockSendStreamI(ctrl *gomock.Controller) *MockSendStreamI {
	mock := &MockSendStreamI{ctrl: ctrl}
	mock.recorder = &MockSendStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendStreamI) EXPECT() *MockSendStreamIMockRecorder {
	return m.recorder
}

// CancelWrite mocks base method.
func (m *MockSendStreamI) CancelWrite(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockSendStreamIMockRecorder) CancelWrite(arg0 any) *MockSendStreamICancelWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockSendStreamI)(nil).CancelWrite), arg0)
	return &MockSendStreamICancelWriteCall{Call: call}
}

// MockSendStreamICancelWriteCall wrap *gomock.Call
type MockSendStreamICancelWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamICancelWriteCall) Return() *MockSendStreamICancelWriteCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamICancelWriteCall) Do(f func(qerr.StreamErrorCode)) *MockSendStreamICancelWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamICancelWriteCall) DoAndReturn(f func(qerr.StreamErrorCode)) *MockSendStreamICancelWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockSendStreamI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSendStreamIMockRecorder) Close() *MockSendStreamICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSendStreamI)(nil).Close))
	return &MockSendStreamICloseCall{Call: call}
}

// MockSendStreamICloseCall wrap *gomock.Call
type MockSendStreamICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamICloseCall) Return(arg0 error) *MockSendStreamICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamICloseCall) Do(f func() error) *MockSendStreamICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamICloseCall) DoAndReturn(f func() error) *MockSendStreamICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockSendStreamI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSendStreamIMockRecorder) Context() *MockSendStreamIContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSendStreamI)(nil).Context))
	return &MockSendStreamIContextCall{Call: call}
}

// MockSendStreamIContextCall wrap *gomock.Call
type MockSendStreamIContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIContextCall) Return(arg0 context.Context) *MockSendStreamIContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIContextCall) Do(f func() context.Context) *MockSendStreamIContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIContextCall) DoAndReturn(f func() context.Context) *MockSendStreamIContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockSendStreamI) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockSendStreamIMockRecorder) SetWriteDeadline(arg0 any) *MockSendStreamISetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockSendStreamI)(nil).SetWriteDeadline), arg0)
	return &MockSendStreamISetWriteDeadlineCall{Call: call}
}

// MockSendStreamISetWriteDeadlineCall wrap *gomock.Call
type MockSendStreamISetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamISetWriteDeadlineCall) Return(arg0 error) *MockSendStreamISetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamISetWriteDeadlineCall) Do(f func(time.Time) error) *MockSendStreamISetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamISetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *MockSendStreamISetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamID mocks base method.
func (m *MockSendStreamI) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockSendStreamIMockRecorder) StreamID() *MockSendStreamIStreamIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockSendStreamI)(nil).StreamID))
	return &MockSendStreamIStreamIDCall{Call: call}
}

// MockSendStreamIStreamIDCall wrap *gomock.Call
type MockSendStreamIStreamIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIStreamIDCall) Return(arg0 protocol.StreamID) *MockSendStreamIStreamIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIStreamIDCall) Do(f func() protocol.StreamID) *MockSendStreamIStreamIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIStreamIDCall) DoAndReturn(f func() protocol.StreamID) *MockSendStreamIStreamIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockSendStreamI) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockSendStreamIMockRecorder) Write(arg0 any) *MockSendStreamIWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSendStreamI)(nil).Write), arg0)
	return &MockSendStreamIWriteCall{Call: call}
}

// MockSendStreamIWriteCall wrap *gomock.Call
type MockSendStreamIWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIWriteCall) Return(arg0 int, arg1 error) *MockSendStreamIWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIWriteCall) Do(f func([]byte) (int, error)) *MockSendStreamIWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIWriteCall) DoAndReturn(f func([]byte) (int, error)) *MockSendStreamIWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// closeForShutdown mocks base method.
func (m *MockSendStreamI) closeForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown.
func (mr *MockSendStreamIMockRecorder) closeForShutdown(arg0 any) *MockSendStreamIcloseForShutdownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockSendStreamI)(nil).closeForShutdown), arg0)
	return &MockSendStreamIcloseForShutdownCall{Call: call}
}

// MockSendStreamIcloseForShutdownCall wrap *gomock.Call
type MockSendStreamIcloseForShutdownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIcloseForShutdownCall) Return() *MockSendStreamIcloseForShutdownCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIcloseForShutdownCall) Do(f func(error)) *MockSendStreamIcloseForShutdownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIcloseForShutdownCall) DoAndReturn(f func(error)) *MockSendStreamIcloseForShutdownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleStopSendingFrame mocks base method.
func (m *MockSendStreamI) handleStopSendingFrame(arg0 *wire.StopSendingFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStopSendingFrame", arg0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame.
func (mr *MockSendStreamIMockRecorder) handleStopSendingFrame(arg0 any) *MockSendStreamIhandleStopSendingFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockSendStreamI)(nil).handleStopSendingFrame), arg0)
	return &MockSendStreamIhandleStopSendingFrameCall{Call: call}
}

// MockSendStreamIhandleStopSendingFrameCall wrap *gomock.Call
type MockSendStreamIhandleStopSendingFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIhandleStopSendingFrameCall) Return() *MockSendStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIhandleStopSendingFrameCall) Do(f func(*wire.StopSendingFrame)) *MockSendStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIhandleStopSendingFrameCall) DoAndReturn(f func(*wire.StopSendingFrame)) *MockSendStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// hasData mocks base method.
func (m *MockSendStreamI) hasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasData indicates an expected call of hasData.
func (mr *MockSendStreamIMockRecorder) hasData() *MockSendStreamIhasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasData", reflect.TypeOf((*MockSendStreamI)(nil).hasData))
	return &MockSendStreamIhasDataCall{Call: call}
}

// MockSendStreamIhasDataCall wrap *gomock.Call
type MockSendStreamIhasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIhasDataCall) Return(arg0 bool) *MockSendStreamIhasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIhasDataCall) Do(f func() bool) *MockSendStreamIhasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIhasDataCall) DoAndReturn(f func() bool) *MockSendStreamIhasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// popStreamFrame mocks base method.
func (m *MockSendStreamI) popStreamFrame(arg0 protocol.ByteCount, arg1 protocol.Version) (ackhandler.StreamFrame, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "popStreamFrame", arg0, arg1)
	ret0, _ := ret[0].(ackhandler.StreamFrame)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// popStreamFrame indicates an expected call of popStreamFrame.
func (mr *MockSendStreamIMockRecorder) popStreamFrame(arg0, arg1 any) *MockSendStreamIpopStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "popStreamFrame", reflect.TypeOf((*MockSendStreamI)(nil).popStreamFrame), arg0, arg1)
	return &MockSendStreamIpopStreamFrameCall{Call: call}
}

// MockSendStreamIpopStreamFrameCall wrap *gomock.Call
type MockSendStreamIpopStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIpopStreamFrameCall) Return(arg0 ackhandler.StreamFrame, arg1, arg2 bool) *MockSendStreamIpopStreamFrameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIpopStreamFrameCall) Do(f func(protocol.ByteCount, protocol.Version) (ackhandler.StreamFrame, bool, bool)) *MockSendStreamIpopStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIpopStreamFrameCall) DoAndReturn(f func(protocol.ByteCount, protocol.Version) (ackhandler.StreamFrame, bool, bool)) *MockSendStreamIpopStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// updateSendWindow mocks base method.
func (m *MockSendStreamI) updateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateSendWindow", arg0)
}

// updateSendWindow indicates an expected call of updateSendWindow.
func (mr *MockSendStreamIMockRecorder) updateSendWindow(arg0 any) *MockSendStreamIupdateSendWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateSendWindow", reflect.TypeOf((*MockSendStreamI)(nil).updateSendWindow), arg0)
	return &MockSendStreamIupdateSendWindowCall{Call: call}
}

// MockSendStreamIupdateSendWindowCall wrap *gomock.Call
type MockSendStreamIupdateSendWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendStreamIupdateSendWindowCall) Return() *MockSendStreamIupdateSendWindowCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendStreamIupdateSendWindowCall) Do(f func(protocol.ByteCount)) *MockSendStreamIupdateSendWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendStreamIupdateSendWindowCall) DoAndReturn(f func(protocol.ByteCount)) *MockSendStreamIupdateSendWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
