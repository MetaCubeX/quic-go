// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go/internal/handshake (interfaces: LongHeaderOpener)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mocks -destination long_header_opener.go github.com/metacubex/quic-go/internal/handshake LongHeaderOpener
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	protocol "github.com/metacubex/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockLongHeaderOpener is a mock of LongHeaderOpener interface.
type MockLongHeaderOpener struct {
	ctrl     *gomock.Controller
	recorder *MockLongHeaderOpenerMockRecorder
	isgomock struct{}
}

// MockLongHeaderOpenerMockRecorder is the mock recorder for MockLongHeaderOpener.
type MockLongHeaderOpenerMockRecorder struct {
	mock *MockLongHeaderOpener
}

// NewMockLongHeaderOpener creates a new mock instance.
func NewMockLongHeaderOpener(ctrl *gomock.Controller) *MockLongHeaderOpener {
	mock := &MockLongHeaderOpener{ctrl: ctrl}
	mock.recorder = &MockLongHeaderOpenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLongHeaderOpener) EXPECT() *MockLongHeaderOpenerMockRecorder {
	return m.recorder
}

// DecodePacketNumber mocks base method.
func (m *MockLongHeaderOpener) DecodePacketNumber(wirePN protocol.PacketNumber, wirePNLen protocol.PacketNumberLen) protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodePacketNumber", wirePN, wirePNLen)
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// DecodePacketNumber indicates an expected call of DecodePacketNumber.
func (mr *MockLongHeaderOpenerMockRecorder) DecodePacketNumber(wirePN, wirePNLen any) *MockLongHeaderOpenerDecodePacketNumberCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodePacketNumber", reflect.TypeOf((*MockLongHeaderOpener)(nil).DecodePacketNumber), wirePN, wirePNLen)
	return &MockLongHeaderOpenerDecodePacketNumberCall{Call: call}
}

// MockLongHeaderOpenerDecodePacketNumberCall wrap *gomock.Call
type MockLongHeaderOpenerDecodePacketNumberCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLongHeaderOpenerDecodePacketNumberCall) Return(arg0 protocol.PacketNumber) *MockLongHeaderOpenerDecodePacketNumberCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLongHeaderOpenerDecodePacketNumberCall) Do(f func(protocol.PacketNumber, protocol.PacketNumberLen) protocol.PacketNumber) *MockLongHeaderOpenerDecodePacketNumberCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLongHeaderOpenerDecodePacketNumberCall) DoAndReturn(f func(protocol.PacketNumber, protocol.PacketNumberLen) protocol.PacketNumber) *MockLongHeaderOpenerDecodePacketNumberCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DecryptHeader mocks base method.
func (m *MockLongHeaderOpener) DecryptHeader(sample []byte, firstByte *byte, pnBytes []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DecryptHeader", sample, firstByte, pnBytes)
}

// DecryptHeader indicates an expected call of DecryptHeader.
func (mr *MockLongHeaderOpenerMockRecorder) DecryptHeader(sample, firstByte, pnBytes any) *MockLongHeaderOpenerDecryptHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptHeader", reflect.TypeOf((*MockLongHeaderOpener)(nil).DecryptHeader), sample, firstByte, pnBytes)
	return &MockLongHeaderOpenerDecryptHeaderCall{Call: call}
}

// MockLongHeaderOpenerDecryptHeaderCall wrap *gomock.Call
type MockLongHeaderOpenerDecryptHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLongHeaderOpenerDecryptHeaderCall) Return() *MockLongHeaderOpenerDecryptHeaderCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLongHeaderOpenerDecryptHeaderCall) Do(f func([]byte, *byte, []byte)) *MockLongHeaderOpenerDecryptHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLongHeaderOpenerDecryptHeaderCall) DoAndReturn(f func([]byte, *byte, []byte)) *MockLongHeaderOpenerDecryptHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Open mocks base method.
func (m *MockLongHeaderOpener) Open(dst, src []byte, pn protocol.PacketNumber, associatedData []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", dst, src, pn, associatedData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockLongHeaderOpenerMockRecorder) Open(dst, src, pn, associatedData any) *MockLongHeaderOpenerOpenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockLongHeaderOpener)(nil).Open), dst, src, pn, associatedData)
	return &MockLongHeaderOpenerOpenCall{Call: call}
}

// MockLongHeaderOpenerOpenCall wrap *gomock.Call
type MockLongHeaderOpenerOpenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLongHeaderOpenerOpenCall) Return(arg0 []byte, arg1 error) *MockLongHeaderOpenerOpenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLongHeaderOpenerOpenCall) Do(f func([]byte, []byte, protocol.PacketNumber, []byte) ([]byte, error)) *MockLongHeaderOpenerOpenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLongHeaderOpenerOpenCall) DoAndReturn(f func([]byte, []byte, protocol.PacketNumber, []byte) ([]byte, error)) *MockLongHeaderOpenerOpenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
