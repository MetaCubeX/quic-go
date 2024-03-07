// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: EarlyConnection)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mockquic -destination quic/early_conn_tmp.go github.com/quic-go/quic-go EarlyConnection
//

// Package mockquic is a generated GoMock package.
package mockquic

import (
	context "context"
	net "net"
	reflect "reflect"

	quic "github.com/quic-go/quic-go"
	congestion "github.com/quic-go/quic-go/congestion"
	qerr "github.com/quic-go/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockEarlyConnection is a mock of EarlyConnection interface.
type MockEarlyConnection struct {
	ctrl     *gomock.Controller
	recorder *MockEarlyConnectionMockRecorder
}

// MockEarlyConnectionMockRecorder is the mock recorder for MockEarlyConnection.
type MockEarlyConnectionMockRecorder struct {
	mock *MockEarlyConnection
}

// NewMockEarlyConnection creates a new mock instance.
func NewMockEarlyConnection(ctrl *gomock.Controller) *MockEarlyConnection {
	mock := &MockEarlyConnection{ctrl: ctrl}
	mock.recorder = &MockEarlyConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEarlyConnection) EXPECT() *MockEarlyConnectionMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockEarlyConnection) AcceptStream(arg0 context.Context) (quic.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(quic.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockEarlyConnectionMockRecorder) AcceptStream(arg0 any) *MockEarlyConnectionAcceptStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockEarlyConnection)(nil).AcceptStream), arg0)
	return &MockEarlyConnectionAcceptStreamCall{Call: call}
}

// MockEarlyConnectionAcceptStreamCall wrap *gomock.Call
type MockEarlyConnectionAcceptStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionAcceptStreamCall) Return(arg0 quic.Stream, arg1 error) *MockEarlyConnectionAcceptStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionAcceptStreamCall) Do(f func(context.Context) (quic.Stream, error)) *MockEarlyConnectionAcceptStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionAcceptStreamCall) DoAndReturn(f func(context.Context) (quic.Stream, error)) *MockEarlyConnectionAcceptStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AcceptUniStream mocks base method.
func (m *MockEarlyConnection) AcceptUniStream(arg0 context.Context) (quic.ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(quic.ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockEarlyConnectionMockRecorder) AcceptUniStream(arg0 any) *MockEarlyConnectionAcceptUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockEarlyConnection)(nil).AcceptUniStream), arg0)
	return &MockEarlyConnectionAcceptUniStreamCall{Call: call}
}

// MockEarlyConnectionAcceptUniStreamCall wrap *gomock.Call
type MockEarlyConnectionAcceptUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionAcceptUniStreamCall) Return(arg0 quic.ReceiveStream, arg1 error) *MockEarlyConnectionAcceptUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionAcceptUniStreamCall) Do(f func(context.Context) (quic.ReceiveStream, error)) *MockEarlyConnectionAcceptUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionAcceptUniStreamCall) DoAndReturn(f func(context.Context) (quic.ReceiveStream, error)) *MockEarlyConnectionAcceptUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloseWithError mocks base method.
func (m *MockEarlyConnection) CloseWithError(arg0 qerr.ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockEarlyConnectionMockRecorder) CloseWithError(arg0, arg1 any) *MockEarlyConnectionCloseWithErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockEarlyConnection)(nil).CloseWithError), arg0, arg1)
	return &MockEarlyConnectionCloseWithErrorCall{Call: call}
}

// MockEarlyConnectionCloseWithErrorCall wrap *gomock.Call
type MockEarlyConnectionCloseWithErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionCloseWithErrorCall) Return(arg0 error) *MockEarlyConnectionCloseWithErrorCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionCloseWithErrorCall) Do(f func(qerr.ApplicationErrorCode, string) error) *MockEarlyConnectionCloseWithErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionCloseWithErrorCall) DoAndReturn(f func(qerr.ApplicationErrorCode, string) error) *MockEarlyConnectionCloseWithErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectionState mocks base method.
func (m *MockEarlyConnection) ConnectionState() quic.ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(quic.ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockEarlyConnectionMockRecorder) ConnectionState() *MockEarlyConnectionConnectionStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockEarlyConnection)(nil).ConnectionState))
	return &MockEarlyConnectionConnectionStateCall{Call: call}
}

// MockEarlyConnectionConnectionStateCall wrap *gomock.Call
type MockEarlyConnectionConnectionStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionConnectionStateCall) Return(arg0 quic.ConnectionState) *MockEarlyConnectionConnectionStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionConnectionStateCall) Do(f func() quic.ConnectionState) *MockEarlyConnectionConnectionStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionConnectionStateCall) DoAndReturn(f func() quic.ConnectionState) *MockEarlyConnectionConnectionStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockEarlyConnection) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockEarlyConnectionMockRecorder) Context() *MockEarlyConnectionContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEarlyConnection)(nil).Context))
	return &MockEarlyConnectionContextCall{Call: call}
}

// MockEarlyConnectionContextCall wrap *gomock.Call
type MockEarlyConnectionContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionContextCall) Return(arg0 context.Context) *MockEarlyConnectionContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionContextCall) Do(f func() context.Context) *MockEarlyConnectionContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionContextCall) DoAndReturn(f func() context.Context) *MockEarlyConnectionContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandshakeComplete mocks base method.
func (m *MockEarlyConnection) HandshakeComplete() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockEarlyConnectionMockRecorder) HandshakeComplete() *MockEarlyConnectionHandshakeCompleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockEarlyConnection)(nil).HandshakeComplete))
	return &MockEarlyConnectionHandshakeCompleteCall{Call: call}
}

// MockEarlyConnectionHandshakeCompleteCall wrap *gomock.Call
type MockEarlyConnectionHandshakeCompleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionHandshakeCompleteCall) Return(arg0 <-chan struct{}) *MockEarlyConnectionHandshakeCompleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionHandshakeCompleteCall) Do(f func() <-chan struct{}) *MockEarlyConnectionHandshakeCompleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionHandshakeCompleteCall) DoAndReturn(f func() <-chan struct{}) *MockEarlyConnectionHandshakeCompleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockEarlyConnection) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockEarlyConnectionMockRecorder) LocalAddr() *MockEarlyConnectionLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockEarlyConnection)(nil).LocalAddr))
	return &MockEarlyConnectionLocalAddrCall{Call: call}
}

// MockEarlyConnectionLocalAddrCall wrap *gomock.Call
type MockEarlyConnectionLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionLocalAddrCall) Return(arg0 net.Addr) *MockEarlyConnectionLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionLocalAddrCall) Do(f func() net.Addr) *MockEarlyConnectionLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionLocalAddrCall) DoAndReturn(f func() net.Addr) *MockEarlyConnectionLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextConnection mocks base method.
func (m *MockEarlyConnection) NextConnection(arg0 context.Context) (quic.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextConnection", arg0)
	ret0, _ := ret[0].(quic.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextConnection indicates an expected call of NextConnection.
func (mr *MockEarlyConnectionMockRecorder) NextConnection(arg0 any) *MockEarlyConnectionNextConnectionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextConnection", reflect.TypeOf((*MockEarlyConnection)(nil).NextConnection), arg0)
	return &MockEarlyConnectionNextConnectionCall{Call: call}
}

// MockEarlyConnectionNextConnectionCall wrap *gomock.Call
type MockEarlyConnectionNextConnectionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionNextConnectionCall) Return(arg0 quic.Connection, arg1 error) *MockEarlyConnectionNextConnectionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionNextConnectionCall) Do(f func(context.Context) (quic.Connection, error)) *MockEarlyConnectionNextConnectionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionNextConnectionCall) DoAndReturn(f func(context.Context) (quic.Connection, error)) *MockEarlyConnectionNextConnectionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStream mocks base method.
func (m *MockEarlyConnection) OpenStream() (quic.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(quic.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockEarlyConnectionMockRecorder) OpenStream() *MockEarlyConnectionOpenStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockEarlyConnection)(nil).OpenStream))
	return &MockEarlyConnectionOpenStreamCall{Call: call}
}

// MockEarlyConnectionOpenStreamCall wrap *gomock.Call
type MockEarlyConnectionOpenStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionOpenStreamCall) Return(arg0 quic.Stream, arg1 error) *MockEarlyConnectionOpenStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionOpenStreamCall) Do(f func() (quic.Stream, error)) *MockEarlyConnectionOpenStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionOpenStreamCall) DoAndReturn(f func() (quic.Stream, error)) *MockEarlyConnectionOpenStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStreamSync mocks base method.
func (m *MockEarlyConnection) OpenStreamSync(arg0 context.Context) (quic.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(quic.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockEarlyConnectionMockRecorder) OpenStreamSync(arg0 any) *MockEarlyConnectionOpenStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockEarlyConnection)(nil).OpenStreamSync), arg0)
	return &MockEarlyConnectionOpenStreamSyncCall{Call: call}
}

// MockEarlyConnectionOpenStreamSyncCall wrap *gomock.Call
type MockEarlyConnectionOpenStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionOpenStreamSyncCall) Return(arg0 quic.Stream, arg1 error) *MockEarlyConnectionOpenStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionOpenStreamSyncCall) Do(f func(context.Context) (quic.Stream, error)) *MockEarlyConnectionOpenStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionOpenStreamSyncCall) DoAndReturn(f func(context.Context) (quic.Stream, error)) *MockEarlyConnectionOpenStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStream mocks base method.
func (m *MockEarlyConnection) OpenUniStream() (quic.SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(quic.SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockEarlyConnectionMockRecorder) OpenUniStream() *MockEarlyConnectionOpenUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockEarlyConnection)(nil).OpenUniStream))
	return &MockEarlyConnectionOpenUniStreamCall{Call: call}
}

// MockEarlyConnectionOpenUniStreamCall wrap *gomock.Call
type MockEarlyConnectionOpenUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionOpenUniStreamCall) Return(arg0 quic.SendStream, arg1 error) *MockEarlyConnectionOpenUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionOpenUniStreamCall) Do(f func() (quic.SendStream, error)) *MockEarlyConnectionOpenUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionOpenUniStreamCall) DoAndReturn(f func() (quic.SendStream, error)) *MockEarlyConnectionOpenUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStreamSync mocks base method.
func (m *MockEarlyConnection) OpenUniStreamSync(arg0 context.Context) (quic.SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(quic.SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockEarlyConnectionMockRecorder) OpenUniStreamSync(arg0 any) *MockEarlyConnectionOpenUniStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockEarlyConnection)(nil).OpenUniStreamSync), arg0)
	return &MockEarlyConnectionOpenUniStreamSyncCall{Call: call}
}

// MockEarlyConnectionOpenUniStreamSyncCall wrap *gomock.Call
type MockEarlyConnectionOpenUniStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionOpenUniStreamSyncCall) Return(arg0 quic.SendStream, arg1 error) *MockEarlyConnectionOpenUniStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionOpenUniStreamSyncCall) Do(f func(context.Context) (quic.SendStream, error)) *MockEarlyConnectionOpenUniStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionOpenUniStreamSyncCall) DoAndReturn(f func(context.Context) (quic.SendStream, error)) *MockEarlyConnectionOpenUniStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReceiveDatagram mocks base method.
func (m *MockEarlyConnection) ReceiveDatagram(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveDatagram", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveDatagram indicates an expected call of ReceiveDatagram.
func (mr *MockEarlyConnectionMockRecorder) ReceiveDatagram(arg0 any) *MockEarlyConnectionReceiveDatagramCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveDatagram", reflect.TypeOf((*MockEarlyConnection)(nil).ReceiveDatagram), arg0)
	return &MockEarlyConnectionReceiveDatagramCall{Call: call}
}

// MockEarlyConnectionReceiveDatagramCall wrap *gomock.Call
type MockEarlyConnectionReceiveDatagramCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionReceiveDatagramCall) Return(arg0 []byte, arg1 error) *MockEarlyConnectionReceiveDatagramCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionReceiveDatagramCall) Do(f func(context.Context) ([]byte, error)) *MockEarlyConnectionReceiveDatagramCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionReceiveDatagramCall) DoAndReturn(f func(context.Context) ([]byte, error)) *MockEarlyConnectionReceiveDatagramCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoteAddr mocks base method.
func (m *MockEarlyConnection) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockEarlyConnectionMockRecorder) RemoteAddr() *MockEarlyConnectionRemoteAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockEarlyConnection)(nil).RemoteAddr))
	return &MockEarlyConnectionRemoteAddrCall{Call: call}
}

// MockEarlyConnectionRemoteAddrCall wrap *gomock.Call
type MockEarlyConnectionRemoteAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionRemoteAddrCall) Return(arg0 net.Addr) *MockEarlyConnectionRemoteAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionRemoteAddrCall) Do(f func() net.Addr) *MockEarlyConnectionRemoteAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionRemoteAddrCall) DoAndReturn(f func() net.Addr) *MockEarlyConnectionRemoteAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendDatagram mocks base method.
func (m *MockEarlyConnection) SendDatagram(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDatagram", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendDatagram indicates an expected call of SendDatagram.
func (mr *MockEarlyConnectionMockRecorder) SendDatagram(arg0 any) *MockEarlyConnectionSendDatagramCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDatagram", reflect.TypeOf((*MockEarlyConnection)(nil).SendDatagram), arg0)
	return &MockEarlyConnectionSendDatagramCall{Call: call}
}

// MockEarlyConnectionSendDatagramCall wrap *gomock.Call
type MockEarlyConnectionSendDatagramCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEarlyConnectionSendDatagramCall) Return(arg0 error) *MockEarlyConnectionSendDatagramCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEarlyConnectionSendDatagramCall) Do(f func([]byte) error) *MockEarlyConnectionSendDatagramCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEarlyConnectionSendDatagramCall) DoAndReturn(f func([]byte) error) *MockEarlyConnectionSendDatagramCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCongestionControl mocks base method.
func (m *MockEarlyConnection) SetCongestionControl(arg0 congestion.CongestionControl) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCongestionControl", arg0)
}

// SetRemoteAddr mocks base method.
func (m *MockEarlyConnection) SetRemoteAddr(arg0 net.Addr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemoteAddr", arg0)
}

// SetCongestionControl indicates an expected call of SetCongestionControl.
func (mr *MockEarlyConnectionMockRecorder) SetCongestionControl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCongestionControl", reflect.TypeOf((*MockEarlyConnection)(nil).SetCongestionControl), arg0)
}
