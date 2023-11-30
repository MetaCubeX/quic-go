// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go (interfaces: StreamManager)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/metacubex/quic-go -destination mock_stream_manager_test.go github.com/metacubex/quic-go StreamManager
//
// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"

	protocol "github.com/metacubex/quic-go/internal/protocol"
	wire "github.com/metacubex/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockStreamManager is a mock of StreamManager interface.
type MockStreamManager struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagerMockRecorder
}

// MockStreamManagerMockRecorder is the mock recorder for MockStreamManager.
type MockStreamManagerMockRecorder struct {
	mock *MockStreamManager
}

// NewMockStreamManager creates a new mock instance.
func NewMockStreamManager(ctrl *gomock.Controller) *MockStreamManager {
	mock := &MockStreamManager{ctrl: ctrl}
	mock.recorder = &MockStreamManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamManager) EXPECT() *MockStreamManagerMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockStreamManager) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockStreamManagerMockRecorder) AcceptStream(arg0 any) *StreamManagerAcceptStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockStreamManager)(nil).AcceptStream), arg0)
	return &StreamManagerAcceptStreamCall{Call: call}
}

// StreamManagerAcceptStreamCall wrap *gomock.Call
type StreamManagerAcceptStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerAcceptStreamCall) Return(arg0 Stream, arg1 error) *StreamManagerAcceptStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerAcceptStreamCall) Do(f func(context.Context) (Stream, error)) *StreamManagerAcceptStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerAcceptStreamCall) DoAndReturn(f func(context.Context) (Stream, error)) *StreamManagerAcceptStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AcceptUniStream mocks base method.
func (m *MockStreamManager) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockStreamManagerMockRecorder) AcceptUniStream(arg0 any) *StreamManagerAcceptUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockStreamManager)(nil).AcceptUniStream), arg0)
	return &StreamManagerAcceptUniStreamCall{Call: call}
}

// StreamManagerAcceptUniStreamCall wrap *gomock.Call
type StreamManagerAcceptUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerAcceptUniStreamCall) Return(arg0 ReceiveStream, arg1 error) *StreamManagerAcceptUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerAcceptUniStreamCall) Do(f func(context.Context) (ReceiveStream, error)) *StreamManagerAcceptUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerAcceptUniStreamCall) DoAndReturn(f func(context.Context) (ReceiveStream, error)) *StreamManagerAcceptUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloseWithError mocks base method.
func (m *MockStreamManager) CloseWithError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseWithError", arg0)
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockStreamManagerMockRecorder) CloseWithError(arg0 any) *StreamManagerCloseWithErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockStreamManager)(nil).CloseWithError), arg0)
	return &StreamManagerCloseWithErrorCall{Call: call}
}

// StreamManagerCloseWithErrorCall wrap *gomock.Call
type StreamManagerCloseWithErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerCloseWithErrorCall) Return() *StreamManagerCloseWithErrorCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerCloseWithErrorCall) Do(f func(error)) *StreamManagerCloseWithErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerCloseWithErrorCall) DoAndReturn(f func(error)) *StreamManagerCloseWithErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteStream mocks base method.
func (m *MockStreamManager) DeleteStream(arg0 protocol.StreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamManagerMockRecorder) DeleteStream(arg0 any) *StreamManagerDeleteStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamManager)(nil).DeleteStream), arg0)
	return &StreamManagerDeleteStreamCall{Call: call}
}

// StreamManagerDeleteStreamCall wrap *gomock.Call
type StreamManagerDeleteStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerDeleteStreamCall) Return(arg0 error) *StreamManagerDeleteStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerDeleteStreamCall) Do(f func(protocol.StreamID) error) *StreamManagerDeleteStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerDeleteStreamCall) DoAndReturn(f func(protocol.StreamID) error) *StreamManagerDeleteStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOrOpenReceiveStream mocks base method.
func (m *MockStreamManager) GetOrOpenReceiveStream(arg0 protocol.StreamID) (receiveStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenReceiveStream", arg0)
	ret0, _ := ret[0].(receiveStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenReceiveStream indicates an expected call of GetOrOpenReceiveStream.
func (mr *MockStreamManagerMockRecorder) GetOrOpenReceiveStream(arg0 any) *StreamManagerGetOrOpenReceiveStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenReceiveStream", reflect.TypeOf((*MockStreamManager)(nil).GetOrOpenReceiveStream), arg0)
	return &StreamManagerGetOrOpenReceiveStreamCall{Call: call}
}

// StreamManagerGetOrOpenReceiveStreamCall wrap *gomock.Call
type StreamManagerGetOrOpenReceiveStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerGetOrOpenReceiveStreamCall) Return(arg0 receiveStreamI, arg1 error) *StreamManagerGetOrOpenReceiveStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerGetOrOpenReceiveStreamCall) Do(f func(protocol.StreamID) (receiveStreamI, error)) *StreamManagerGetOrOpenReceiveStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerGetOrOpenReceiveStreamCall) DoAndReturn(f func(protocol.StreamID) (receiveStreamI, error)) *StreamManagerGetOrOpenReceiveStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOrOpenSendStream mocks base method.
func (m *MockStreamManager) GetOrOpenSendStream(arg0 protocol.StreamID) (sendStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenSendStream", arg0)
	ret0, _ := ret[0].(sendStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenSendStream indicates an expected call of GetOrOpenSendStream.
func (mr *MockStreamManagerMockRecorder) GetOrOpenSendStream(arg0 any) *StreamManagerGetOrOpenSendStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenSendStream", reflect.TypeOf((*MockStreamManager)(nil).GetOrOpenSendStream), arg0)
	return &StreamManagerGetOrOpenSendStreamCall{Call: call}
}

// StreamManagerGetOrOpenSendStreamCall wrap *gomock.Call
type StreamManagerGetOrOpenSendStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerGetOrOpenSendStreamCall) Return(arg0 sendStreamI, arg1 error) *StreamManagerGetOrOpenSendStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerGetOrOpenSendStreamCall) Do(f func(protocol.StreamID) (sendStreamI, error)) *StreamManagerGetOrOpenSendStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerGetOrOpenSendStreamCall) DoAndReturn(f func(protocol.StreamID) (sendStreamI, error)) *StreamManagerGetOrOpenSendStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandleMaxStreamsFrame mocks base method.
func (m *MockStreamManager) HandleMaxStreamsFrame(arg0 *wire.MaxStreamsFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMaxStreamsFrame", arg0)
}

// HandleMaxStreamsFrame indicates an expected call of HandleMaxStreamsFrame.
func (mr *MockStreamManagerMockRecorder) HandleMaxStreamsFrame(arg0 any) *StreamManagerHandleMaxStreamsFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMaxStreamsFrame", reflect.TypeOf((*MockStreamManager)(nil).HandleMaxStreamsFrame), arg0)
	return &StreamManagerHandleMaxStreamsFrameCall{Call: call}
}

// StreamManagerHandleMaxStreamsFrameCall wrap *gomock.Call
type StreamManagerHandleMaxStreamsFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerHandleMaxStreamsFrameCall) Return() *StreamManagerHandleMaxStreamsFrameCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerHandleMaxStreamsFrameCall) Do(f func(*wire.MaxStreamsFrame)) *StreamManagerHandleMaxStreamsFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerHandleMaxStreamsFrameCall) DoAndReturn(f func(*wire.MaxStreamsFrame)) *StreamManagerHandleMaxStreamsFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStream mocks base method.
func (m *MockStreamManager) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockStreamManagerMockRecorder) OpenStream() *StreamManagerOpenStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockStreamManager)(nil).OpenStream))
	return &StreamManagerOpenStreamCall{Call: call}
}

// StreamManagerOpenStreamCall wrap *gomock.Call
type StreamManagerOpenStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerOpenStreamCall) Return(arg0 Stream, arg1 error) *StreamManagerOpenStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerOpenStreamCall) Do(f func() (Stream, error)) *StreamManagerOpenStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerOpenStreamCall) DoAndReturn(f func() (Stream, error)) *StreamManagerOpenStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStreamSync mocks base method.
func (m *MockStreamManager) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockStreamManagerMockRecorder) OpenStreamSync(arg0 any) *StreamManagerOpenStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockStreamManager)(nil).OpenStreamSync), arg0)
	return &StreamManagerOpenStreamSyncCall{Call: call}
}

// StreamManagerOpenStreamSyncCall wrap *gomock.Call
type StreamManagerOpenStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerOpenStreamSyncCall) Return(arg0 Stream, arg1 error) *StreamManagerOpenStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerOpenStreamSyncCall) Do(f func(context.Context) (Stream, error)) *StreamManagerOpenStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerOpenStreamSyncCall) DoAndReturn(f func(context.Context) (Stream, error)) *StreamManagerOpenStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStream mocks base method.
func (m *MockStreamManager) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockStreamManagerMockRecorder) OpenUniStream() *StreamManagerOpenUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockStreamManager)(nil).OpenUniStream))
	return &StreamManagerOpenUniStreamCall{Call: call}
}

// StreamManagerOpenUniStreamCall wrap *gomock.Call
type StreamManagerOpenUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerOpenUniStreamCall) Return(arg0 SendStream, arg1 error) *StreamManagerOpenUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerOpenUniStreamCall) Do(f func() (SendStream, error)) *StreamManagerOpenUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerOpenUniStreamCall) DoAndReturn(f func() (SendStream, error)) *StreamManagerOpenUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStreamSync mocks base method.
func (m *MockStreamManager) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockStreamManagerMockRecorder) OpenUniStreamSync(arg0 any) *StreamManagerOpenUniStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockStreamManager)(nil).OpenUniStreamSync), arg0)
	return &StreamManagerOpenUniStreamSyncCall{Call: call}
}

// StreamManagerOpenUniStreamSyncCall wrap *gomock.Call
type StreamManagerOpenUniStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerOpenUniStreamSyncCall) Return(arg0 SendStream, arg1 error) *StreamManagerOpenUniStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerOpenUniStreamSyncCall) Do(f func(context.Context) (SendStream, error)) *StreamManagerOpenUniStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerOpenUniStreamSyncCall) DoAndReturn(f func(context.Context) (SendStream, error)) *StreamManagerOpenUniStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResetFor0RTT mocks base method.
func (m *MockStreamManager) ResetFor0RTT() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetFor0RTT")
}

// ResetFor0RTT indicates an expected call of ResetFor0RTT.
func (mr *MockStreamManagerMockRecorder) ResetFor0RTT() *StreamManagerResetFor0RTTCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFor0RTT", reflect.TypeOf((*MockStreamManager)(nil).ResetFor0RTT))
	return &StreamManagerResetFor0RTTCall{Call: call}
}

// StreamManagerResetFor0RTTCall wrap *gomock.Call
type StreamManagerResetFor0RTTCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerResetFor0RTTCall) Return() *StreamManagerResetFor0RTTCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerResetFor0RTTCall) Do(f func()) *StreamManagerResetFor0RTTCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerResetFor0RTTCall) DoAndReturn(f func()) *StreamManagerResetFor0RTTCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateLimits mocks base method.
func (m *MockStreamManager) UpdateLimits(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateLimits", arg0)
}

// UpdateLimits indicates an expected call of UpdateLimits.
func (mr *MockStreamManagerMockRecorder) UpdateLimits(arg0 any) *StreamManagerUpdateLimitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLimits", reflect.TypeOf((*MockStreamManager)(nil).UpdateLimits), arg0)
	return &StreamManagerUpdateLimitsCall{Call: call}
}

// StreamManagerUpdateLimitsCall wrap *gomock.Call
type StreamManagerUpdateLimitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerUpdateLimitsCall) Return() *StreamManagerUpdateLimitsCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerUpdateLimitsCall) Do(f func(*wire.TransportParameters)) *StreamManagerUpdateLimitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerUpdateLimitsCall) DoAndReturn(f func(*wire.TransportParameters)) *StreamManagerUpdateLimitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UseResetMaps mocks base method.
func (m *MockStreamManager) UseResetMaps() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UseResetMaps")
}

// UseResetMaps indicates an expected call of UseResetMaps.
func (mr *MockStreamManagerMockRecorder) UseResetMaps() *StreamManagerUseResetMapsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseResetMaps", reflect.TypeOf((*MockStreamManager)(nil).UseResetMaps))
	return &StreamManagerUseResetMapsCall{Call: call}
}

// StreamManagerUseResetMapsCall wrap *gomock.Call
type StreamManagerUseResetMapsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamManagerUseResetMapsCall) Return() *StreamManagerUseResetMapsCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamManagerUseResetMapsCall) Do(f func()) *StreamManagerUseResetMapsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamManagerUseResetMapsCall) DoAndReturn(f func()) *StreamManagerUseResetMapsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
