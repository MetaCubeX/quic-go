// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go (interfaces: QUICConn)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/metacubex/quic-go -destination mock_quic_conn_test.go github.com/metacubex/quic-go QUICConn
//

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	net "net"
	reflect "reflect"

	congestion "github.com/metacubex/quic-go/congestion"
	qerr "github.com/metacubex/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockQUICConn is a mock of QUICConn interface.
type MockQUICConn struct {
	ctrl     *gomock.Controller
	recorder *MockQUICConnMockRecorder
}

// MockQUICConnMockRecorder is the mock recorder for MockQUICConn.
type MockQUICConnMockRecorder struct {
	mock *MockQUICConn
}

// NewMockQUICConn creates a new mock instance.
func NewMockQUICConn(ctrl *gomock.Controller) *MockQUICConn {
	mock := &MockQUICConn{ctrl: ctrl}
	mock.recorder = &MockQUICConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQUICConn) EXPECT() *MockQUICConnMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockQUICConn) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockQUICConnMockRecorder) AcceptStream(arg0 any) *MockQUICConnAcceptStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockQUICConn)(nil).AcceptStream), arg0)
	return &MockQUICConnAcceptStreamCall{Call: call}
}

// MockQUICConnAcceptStreamCall wrap *gomock.Call
type MockQUICConnAcceptStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnAcceptStreamCall) Return(arg0 Stream, arg1 error) *MockQUICConnAcceptStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnAcceptStreamCall) Do(f func(context.Context) (Stream, error)) *MockQUICConnAcceptStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnAcceptStreamCall) DoAndReturn(f func(context.Context) (Stream, error)) *MockQUICConnAcceptStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AcceptUniStream mocks base method.
func (m *MockQUICConn) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockQUICConnMockRecorder) AcceptUniStream(arg0 any) *MockQUICConnAcceptUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockQUICConn)(nil).AcceptUniStream), arg0)
	return &MockQUICConnAcceptUniStreamCall{Call: call}
}

// MockQUICConnAcceptUniStreamCall wrap *gomock.Call
type MockQUICConnAcceptUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnAcceptUniStreamCall) Return(arg0 ReceiveStream, arg1 error) *MockQUICConnAcceptUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnAcceptUniStreamCall) Do(f func(context.Context) (ReceiveStream, error)) *MockQUICConnAcceptUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnAcceptUniStreamCall) DoAndReturn(f func(context.Context) (ReceiveStream, error)) *MockQUICConnAcceptUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloseWithError mocks base method.
func (m *MockQUICConn) CloseWithError(arg0 qerr.ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockQUICConnMockRecorder) CloseWithError(arg0, arg1 any) *MockQUICConnCloseWithErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockQUICConn)(nil).CloseWithError), arg0, arg1)
	return &MockQUICConnCloseWithErrorCall{Call: call}
}

// MockQUICConnCloseWithErrorCall wrap *gomock.Call
type MockQUICConnCloseWithErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnCloseWithErrorCall) Return(arg0 error) *MockQUICConnCloseWithErrorCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnCloseWithErrorCall) Do(f func(qerr.ApplicationErrorCode, string) error) *MockQUICConnCloseWithErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnCloseWithErrorCall) DoAndReturn(f func(qerr.ApplicationErrorCode, string) error) *MockQUICConnCloseWithErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectionState mocks base method.
func (m *MockQUICConn) ConnectionState() ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockQUICConnMockRecorder) ConnectionState() *MockQUICConnConnectionStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockQUICConn)(nil).ConnectionState))
	return &MockQUICConnConnectionStateCall{Call: call}
}

// MockQUICConnConnectionStateCall wrap *gomock.Call
type MockQUICConnConnectionStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnConnectionStateCall) Return(arg0 ConnectionState) *MockQUICConnConnectionStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnConnectionStateCall) Do(f func() ConnectionState) *MockQUICConnConnectionStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnConnectionStateCall) DoAndReturn(f func() ConnectionState) *MockQUICConnConnectionStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockQUICConn) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockQUICConnMockRecorder) Context() *MockQUICConnContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQUICConn)(nil).Context))
	return &MockQUICConnContextCall{Call: call}
}

// MockQUICConnContextCall wrap *gomock.Call
type MockQUICConnContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnContextCall) Return(arg0 context.Context) *MockQUICConnContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnContextCall) Do(f func() context.Context) *MockQUICConnContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnContextCall) DoAndReturn(f func() context.Context) *MockQUICConnContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandshakeComplete mocks base method.
func (m *MockQUICConn) HandshakeComplete() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockQUICConnMockRecorder) HandshakeComplete() *MockQUICConnHandshakeCompleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockQUICConn)(nil).HandshakeComplete))
	return &MockQUICConnHandshakeCompleteCall{Call: call}
}

// MockQUICConnHandshakeCompleteCall wrap *gomock.Call
type MockQUICConnHandshakeCompleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnHandshakeCompleteCall) Return(arg0 <-chan struct{}) *MockQUICConnHandshakeCompleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnHandshakeCompleteCall) Do(f func() <-chan struct{}) *MockQUICConnHandshakeCompleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnHandshakeCompleteCall) DoAndReturn(f func() <-chan struct{}) *MockQUICConnHandshakeCompleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockQUICConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockQUICConnMockRecorder) LocalAddr() *MockQUICConnLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockQUICConn)(nil).LocalAddr))
	return &MockQUICConnLocalAddrCall{Call: call}
}

// MockQUICConnLocalAddrCall wrap *gomock.Call
type MockQUICConnLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnLocalAddrCall) Return(arg0 net.Addr) *MockQUICConnLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnLocalAddrCall) Do(f func() net.Addr) *MockQUICConnLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnLocalAddrCall) DoAndReturn(f func() net.Addr) *MockQUICConnLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextConnection mocks base method.
func (m *MockQUICConn) NextConnection() Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextConnection")
	ret0, _ := ret[0].(Connection)
	return ret0
}

// NextConnection indicates an expected call of NextConnection.
func (mr *MockQUICConnMockRecorder) NextConnection() *MockQUICConnNextConnectionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextConnection", reflect.TypeOf((*MockQUICConn)(nil).NextConnection))
	return &MockQUICConnNextConnectionCall{Call: call}
}

// MockQUICConnNextConnectionCall wrap *gomock.Call
type MockQUICConnNextConnectionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnNextConnectionCall) Return(arg0 Connection) *MockQUICConnNextConnectionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnNextConnectionCall) Do(f func() Connection) *MockQUICConnNextConnectionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnNextConnectionCall) DoAndReturn(f func() Connection) *MockQUICConnNextConnectionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStream mocks base method.
func (m *MockQUICConn) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockQUICConnMockRecorder) OpenStream() *MockQUICConnOpenStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockQUICConn)(nil).OpenStream))
	return &MockQUICConnOpenStreamCall{Call: call}
}

// MockQUICConnOpenStreamCall wrap *gomock.Call
type MockQUICConnOpenStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnOpenStreamCall) Return(arg0 Stream, arg1 error) *MockQUICConnOpenStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnOpenStreamCall) Do(f func() (Stream, error)) *MockQUICConnOpenStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnOpenStreamCall) DoAndReturn(f func() (Stream, error)) *MockQUICConnOpenStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStreamSync mocks base method.
func (m *MockQUICConn) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockQUICConnMockRecorder) OpenStreamSync(arg0 any) *MockQUICConnOpenStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockQUICConn)(nil).OpenStreamSync), arg0)
	return &MockQUICConnOpenStreamSyncCall{Call: call}
}

// MockQUICConnOpenStreamSyncCall wrap *gomock.Call
type MockQUICConnOpenStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnOpenStreamSyncCall) Return(arg0 Stream, arg1 error) *MockQUICConnOpenStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnOpenStreamSyncCall) Do(f func(context.Context) (Stream, error)) *MockQUICConnOpenStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnOpenStreamSyncCall) DoAndReturn(f func(context.Context) (Stream, error)) *MockQUICConnOpenStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStream mocks base method.
func (m *MockQUICConn) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockQUICConnMockRecorder) OpenUniStream() *MockQUICConnOpenUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockQUICConn)(nil).OpenUniStream))
	return &MockQUICConnOpenUniStreamCall{Call: call}
}

// MockQUICConnOpenUniStreamCall wrap *gomock.Call
type MockQUICConnOpenUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnOpenUniStreamCall) Return(arg0 SendStream, arg1 error) *MockQUICConnOpenUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnOpenUniStreamCall) Do(f func() (SendStream, error)) *MockQUICConnOpenUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnOpenUniStreamCall) DoAndReturn(f func() (SendStream, error)) *MockQUICConnOpenUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStreamSync mocks base method.
func (m *MockQUICConn) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockQUICConnMockRecorder) OpenUniStreamSync(arg0 any) *MockQUICConnOpenUniStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockQUICConn)(nil).OpenUniStreamSync), arg0)
	return &MockQUICConnOpenUniStreamSyncCall{Call: call}
}

// MockQUICConnOpenUniStreamSyncCall wrap *gomock.Call
type MockQUICConnOpenUniStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnOpenUniStreamSyncCall) Return(arg0 SendStream, arg1 error) *MockQUICConnOpenUniStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnOpenUniStreamSyncCall) Do(f func(context.Context) (SendStream, error)) *MockQUICConnOpenUniStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnOpenUniStreamSyncCall) DoAndReturn(f func(context.Context) (SendStream, error)) *MockQUICConnOpenUniStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReceiveDatagram mocks base method.
func (m *MockQUICConn) ReceiveDatagram(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveDatagram", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveDatagram indicates an expected call of ReceiveDatagram.
func (mr *MockQUICConnMockRecorder) ReceiveDatagram(arg0 any) *MockQUICConnReceiveDatagramCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveDatagram", reflect.TypeOf((*MockQUICConn)(nil).ReceiveDatagram), arg0)
	return &MockQUICConnReceiveDatagramCall{Call: call}
}

// MockQUICConnReceiveDatagramCall wrap *gomock.Call
type MockQUICConnReceiveDatagramCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnReceiveDatagramCall) Return(arg0 []byte, arg1 error) *MockQUICConnReceiveDatagramCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnReceiveDatagramCall) Do(f func(context.Context) ([]byte, error)) *MockQUICConnReceiveDatagramCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnReceiveDatagramCall) DoAndReturn(f func(context.Context) ([]byte, error)) *MockQUICConnReceiveDatagramCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoteAddr mocks base method.
func (m *MockQUICConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockQUICConnMockRecorder) RemoteAddr() *MockQUICConnRemoteAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockQUICConn)(nil).RemoteAddr))
	return &MockQUICConnRemoteAddrCall{Call: call}
}

// MockQUICConnRemoteAddrCall wrap *gomock.Call
type MockQUICConnRemoteAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnRemoteAddrCall) Return(arg0 net.Addr) *MockQUICConnRemoteAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnRemoteAddrCall) Do(f func() net.Addr) *MockQUICConnRemoteAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnRemoteAddrCall) DoAndReturn(f func() net.Addr) *MockQUICConnRemoteAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendDatagram mocks base method.
func (m *MockQUICConn) SendDatagram(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDatagram", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendDatagram indicates an expected call of SendDatagram.
func (mr *MockQUICConnMockRecorder) SendDatagram(arg0 any) *MockQUICConnSendDatagramCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDatagram", reflect.TypeOf((*MockQUICConn)(nil).SendDatagram), arg0)
	return &MockQUICConnSendDatagramCall{Call: call}
}

// MockQUICConnSendDatagramCall wrap *gomock.Call
type MockQUICConnSendDatagramCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnSendDatagramCall) Return(arg0 error) *MockQUICConnSendDatagramCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnSendDatagramCall) Do(f func([]byte) error) *MockQUICConnSendDatagramCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnSendDatagramCall) DoAndReturn(f func([]byte) error) *MockQUICConnSendDatagramCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// closeWithTransportError mocks base method.
func (m *MockQUICConn) closeWithTransportError(arg0 qerr.TransportErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeWithTransportError", arg0)
}

// closeWithTransportError indicates an expected call of closeWithTransportError.
func (mr *MockQUICConnMockRecorder) closeWithTransportError(arg0 any) *MockQUICConncloseWithTransportErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeWithTransportError", reflect.TypeOf((*MockQUICConn)(nil).closeWithTransportError), arg0)
	return &MockQUICConncloseWithTransportErrorCall{Call: call}
}

// MockQUICConncloseWithTransportErrorCall wrap *gomock.Call
type MockQUICConncloseWithTransportErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConncloseWithTransportErrorCall) Return() *MockQUICConncloseWithTransportErrorCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConncloseWithTransportErrorCall) Do(f func(qerr.TransportErrorCode)) *MockQUICConncloseWithTransportErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConncloseWithTransportErrorCall) DoAndReturn(f func(qerr.TransportErrorCode)) *MockQUICConncloseWithTransportErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCongestionControl mocks base method.
func (m *MockQUICConn) SetCongestionControl(arg0 congestion.CongestionControl) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCongestionControl", arg0)
}

// SetRemoteAddr mocks base method.
func (m *MockQUICConn) SetRemoteAddr(arg0 net.Addr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemoteAddr", arg0)
}

// Config mocks base method.
func (m *MockQUICConn) Config() *Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*Config)
	return ret0
}

// SetCongestionControl indicates an expected call of SetCongestionControl.
func (mr *MockQUICConnMockRecorder) SetCongestionControl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCongestionControl", reflect.TypeOf((*MockQUICConn)(nil).SetCongestionControl), arg0)
}

// destroy mocks base method.
func (m *MockQUICConn) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockQUICConnMockRecorder) destroy(arg0 any) *MockQUICConndestroyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockQUICConn)(nil).destroy), arg0)
	return &MockQUICConndestroyCall{Call: call}
}

// MockQUICConndestroyCall wrap *gomock.Call
type MockQUICConndestroyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConndestroyCall) Return() *MockQUICConndestroyCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConndestroyCall) Do(f func(error)) *MockQUICConndestroyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConndestroyCall) DoAndReturn(f func(error)) *MockQUICConndestroyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// earlyConnReady mocks base method.
func (m *MockQUICConn) earlyConnReady() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "earlyConnReady")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// earlyConnReady indicates an expected call of earlyConnReady.
func (mr *MockQUICConnMockRecorder) earlyConnReady() *MockQUICConnearlyConnReadyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "earlyConnReady", reflect.TypeOf((*MockQUICConn)(nil).earlyConnReady))
	return &MockQUICConnearlyConnReadyCall{Call: call}
}

// MockQUICConnearlyConnReadyCall wrap *gomock.Call
type MockQUICConnearlyConnReadyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnearlyConnReadyCall) Return(arg0 <-chan struct{}) *MockQUICConnearlyConnReadyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnearlyConnReadyCall) Do(f func() <-chan struct{}) *MockQUICConnearlyConnReadyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnearlyConnReadyCall) DoAndReturn(f func() <-chan struct{}) *MockQUICConnearlyConnReadyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handlePacket mocks base method.
func (m *MockQUICConn) handlePacket(arg0 receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockQUICConnMockRecorder) handlePacket(arg0 any) *MockQUICConnhandlePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockQUICConn)(nil).handlePacket), arg0)
	return &MockQUICConnhandlePacketCall{Call: call}
}

// MockQUICConnhandlePacketCall wrap *gomock.Call
type MockQUICConnhandlePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnhandlePacketCall) Return() *MockQUICConnhandlePacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnhandlePacketCall) Do(f func(receivedPacket)) *MockQUICConnhandlePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnhandlePacketCall) DoAndReturn(f func(receivedPacket)) *MockQUICConnhandlePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// run mocks base method.
func (m *MockQUICConn) run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "run")
	ret0, _ := ret[0].(error)
	return ret0
}

// run indicates an expected call of run.
func (mr *MockQUICConnMockRecorder) run() *MockQUICConnrunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "run", reflect.TypeOf((*MockQUICConn)(nil).run))
	return &MockQUICConnrunCall{Call: call}
}

// MockQUICConnrunCall wrap *gomock.Call
type MockQUICConnrunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICConnrunCall) Return(arg0 error) *MockQUICConnrunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICConnrunCall) Do(f func() error) *MockQUICConnrunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICConnrunCall) DoAndReturn(f func() error) *MockQUICConnrunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
