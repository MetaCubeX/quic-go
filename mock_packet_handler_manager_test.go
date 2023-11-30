// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metacubex/quic-go (interfaces: PacketHandlerManager)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/metacubex/quic-go -destination mock_packet_handler_manager_test.go github.com/metacubex/quic-go PacketHandlerManager
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/metacubex/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockPacketHandlerManager is a mock of PacketHandlerManager interface.
type MockPacketHandlerManager struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerManagerMockRecorder
}

// MockPacketHandlerManagerMockRecorder is the mock recorder for MockPacketHandlerManager.
type MockPacketHandlerManagerMockRecorder struct {
	mock *MockPacketHandlerManager
}

// NewMockPacketHandlerManager creates a new mock instance.
func NewMockPacketHandlerManager(ctrl *gomock.Controller) *MockPacketHandlerManager {
	mock := &MockPacketHandlerManager{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketHandlerManager) EXPECT() *MockPacketHandlerManagerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPacketHandlerManager) Add(arg0 protocol.ConnectionID, arg1 packetHandler) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPacketHandlerManagerMockRecorder) Add(arg0, arg1 any) *PacketHandlerManagerAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPacketHandlerManager)(nil).Add), arg0, arg1)
	return &PacketHandlerManagerAddCall{Call: call}
}

// PacketHandlerManagerAddCall wrap *gomock.Call
type PacketHandlerManagerAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerAddCall) Return(arg0 bool) *PacketHandlerManagerAddCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerAddCall) Do(f func(protocol.ConnectionID, packetHandler) bool) *PacketHandlerManagerAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerAddCall) DoAndReturn(f func(protocol.ConnectionID, packetHandler) bool) *PacketHandlerManagerAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddResetToken mocks base method.
func (m *MockPacketHandlerManager) AddResetToken(arg0 protocol.StatelessResetToken, arg1 packetHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResetToken", arg0, arg1)
}

// AddResetToken indicates an expected call of AddResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) AddResetToken(arg0, arg1 any) *PacketHandlerManagerAddResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).AddResetToken), arg0, arg1)
	return &PacketHandlerManagerAddResetTokenCall{Call: call}
}

// PacketHandlerManagerAddResetTokenCall wrap *gomock.Call
type PacketHandlerManagerAddResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerAddResetTokenCall) Return() *PacketHandlerManagerAddResetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerAddResetTokenCall) Do(f func(protocol.StatelessResetToken, packetHandler)) *PacketHandlerManagerAddResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerAddResetTokenCall) DoAndReturn(f func(protocol.StatelessResetToken, packetHandler)) *PacketHandlerManagerAddResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddWithConnID mocks base method.
func (m *MockPacketHandlerManager) AddWithConnID(arg0, arg1 protocol.ConnectionID, arg2 func() (packetHandler, bool)) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWithConnID", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddWithConnID indicates an expected call of AddWithConnID.
func (mr *MockPacketHandlerManagerMockRecorder) AddWithConnID(arg0, arg1, arg2 any) *PacketHandlerManagerAddWithConnIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWithConnID", reflect.TypeOf((*MockPacketHandlerManager)(nil).AddWithConnID), arg0, arg1, arg2)
	return &PacketHandlerManagerAddWithConnIDCall{Call: call}
}

// PacketHandlerManagerAddWithConnIDCall wrap *gomock.Call
type PacketHandlerManagerAddWithConnIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerAddWithConnIDCall) Return(arg0 bool) *PacketHandlerManagerAddWithConnIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerAddWithConnIDCall) Do(f func(protocol.ConnectionID, protocol.ConnectionID, func() (packetHandler, bool)) bool) *PacketHandlerManagerAddWithConnIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerAddWithConnIDCall) DoAndReturn(f func(protocol.ConnectionID, protocol.ConnectionID, func() (packetHandler, bool)) bool) *PacketHandlerManagerAddWithConnIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockPacketHandlerManager) Close(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", arg0)
}

// Close indicates an expected call of Close.
func (mr *MockPacketHandlerManagerMockRecorder) Close(arg0 any) *PacketHandlerManagerCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPacketHandlerManager)(nil).Close), arg0)
	return &PacketHandlerManagerCloseCall{Call: call}
}

// PacketHandlerManagerCloseCall wrap *gomock.Call
type PacketHandlerManagerCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerCloseCall) Return() *PacketHandlerManagerCloseCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerCloseCall) Do(f func(error)) *PacketHandlerManagerCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerCloseCall) DoAndReturn(f func(error)) *PacketHandlerManagerCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockPacketHandlerManager) Get(arg0 protocol.ConnectionID) (packetHandler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(packetHandler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPacketHandlerManagerMockRecorder) Get(arg0 any) *PacketHandlerManagerGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPacketHandlerManager)(nil).Get), arg0)
	return &PacketHandlerManagerGetCall{Call: call}
}

// PacketHandlerManagerGetCall wrap *gomock.Call
type PacketHandlerManagerGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerGetCall) Return(arg0 packetHandler, arg1 bool) *PacketHandlerManagerGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerGetCall) Do(f func(protocol.ConnectionID) (packetHandler, bool)) *PacketHandlerManagerGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerGetCall) DoAndReturn(f func(protocol.ConnectionID) (packetHandler, bool)) *PacketHandlerManagerGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByResetToken mocks base method.
func (m *MockPacketHandlerManager) GetByResetToken(arg0 protocol.StatelessResetToken) (packetHandler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByResetToken", arg0)
	ret0, _ := ret[0].(packetHandler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetByResetToken indicates an expected call of GetByResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) GetByResetToken(arg0 any) *PacketHandlerManagerGetByResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).GetByResetToken), arg0)
	return &PacketHandlerManagerGetByResetTokenCall{Call: call}
}

// PacketHandlerManagerGetByResetTokenCall wrap *gomock.Call
type PacketHandlerManagerGetByResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerGetByResetTokenCall) Return(arg0 packetHandler, arg1 bool) *PacketHandlerManagerGetByResetTokenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerGetByResetTokenCall) Do(f func(protocol.StatelessResetToken) (packetHandler, bool)) *PacketHandlerManagerGetByResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerGetByResetTokenCall) DoAndReturn(f func(protocol.StatelessResetToken) (packetHandler, bool)) *PacketHandlerManagerGetByResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetStatelessResetToken mocks base method.
func (m *MockPacketHandlerManager) GetStatelessResetToken(arg0 protocol.ConnectionID) protocol.StatelessResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessResetToken", arg0)
	ret0, _ := ret[0].(protocol.StatelessResetToken)
	return ret0
}

// GetStatelessResetToken indicates an expected call of GetStatelessResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) GetStatelessResetToken(arg0 any) *PacketHandlerManagerGetStatelessResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).GetStatelessResetToken), arg0)
	return &PacketHandlerManagerGetStatelessResetTokenCall{Call: call}
}

// PacketHandlerManagerGetStatelessResetTokenCall wrap *gomock.Call
type PacketHandlerManagerGetStatelessResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerGetStatelessResetTokenCall) Return(arg0 protocol.StatelessResetToken) *PacketHandlerManagerGetStatelessResetTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerGetStatelessResetTokenCall) Do(f func(protocol.ConnectionID) protocol.StatelessResetToken) *PacketHandlerManagerGetStatelessResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerGetStatelessResetTokenCall) DoAndReturn(f func(protocol.ConnectionID) protocol.StatelessResetToken) *PacketHandlerManagerGetStatelessResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Remove mocks base method.
func (m *MockPacketHandlerManager) Remove(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockPacketHandlerManagerMockRecorder) Remove(arg0 any) *PacketHandlerManagerRemoveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPacketHandlerManager)(nil).Remove), arg0)
	return &PacketHandlerManagerRemoveCall{Call: call}
}

// PacketHandlerManagerRemoveCall wrap *gomock.Call
type PacketHandlerManagerRemoveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerRemoveCall) Return() *PacketHandlerManagerRemoveCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerRemoveCall) Do(f func(protocol.ConnectionID)) *PacketHandlerManagerRemoveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerRemoveCall) DoAndReturn(f func(protocol.ConnectionID)) *PacketHandlerManagerRemoveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveResetToken mocks base method.
func (m *MockPacketHandlerManager) RemoveResetToken(arg0 protocol.StatelessResetToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveResetToken", arg0)
}

// RemoveResetToken indicates an expected call of RemoveResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) RemoveResetToken(arg0 any) *PacketHandlerManagerRemoveResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).RemoveResetToken), arg0)
	return &PacketHandlerManagerRemoveResetTokenCall{Call: call}
}

// PacketHandlerManagerRemoveResetTokenCall wrap *gomock.Call
type PacketHandlerManagerRemoveResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerRemoveResetTokenCall) Return() *PacketHandlerManagerRemoveResetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerRemoveResetTokenCall) Do(f func(protocol.StatelessResetToken)) *PacketHandlerManagerRemoveResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerRemoveResetTokenCall) DoAndReturn(f func(protocol.StatelessResetToken)) *PacketHandlerManagerRemoveResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReplaceWithClosed mocks base method.
func (m *MockPacketHandlerManager) ReplaceWithClosed(arg0 []protocol.ConnectionID, arg1 protocol.Perspective, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceWithClosed", arg0, arg1, arg2)
}

// ReplaceWithClosed indicates an expected call of ReplaceWithClosed.
func (mr *MockPacketHandlerManagerMockRecorder) ReplaceWithClosed(arg0, arg1, arg2 any) *PacketHandlerManagerReplaceWithClosedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceWithClosed", reflect.TypeOf((*MockPacketHandlerManager)(nil).ReplaceWithClosed), arg0, arg1, arg2)
	return &PacketHandlerManagerReplaceWithClosedCall{Call: call}
}

// PacketHandlerManagerReplaceWithClosedCall wrap *gomock.Call
type PacketHandlerManagerReplaceWithClosedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerReplaceWithClosedCall) Return() *PacketHandlerManagerReplaceWithClosedCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerReplaceWithClosedCall) Do(f func([]protocol.ConnectionID, protocol.Perspective, []byte)) *PacketHandlerManagerReplaceWithClosedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerReplaceWithClosedCall) DoAndReturn(f func([]protocol.ConnectionID, protocol.Perspective, []byte)) *PacketHandlerManagerReplaceWithClosedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Retire mocks base method.
func (m *MockPacketHandlerManager) Retire(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Retire", arg0)
}

// Retire indicates an expected call of Retire.
func (mr *MockPacketHandlerManagerMockRecorder) Retire(arg0 any) *PacketHandlerManagerRetireCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retire", reflect.TypeOf((*MockPacketHandlerManager)(nil).Retire), arg0)
	return &PacketHandlerManagerRetireCall{Call: call}
}

// PacketHandlerManagerRetireCall wrap *gomock.Call
type PacketHandlerManagerRetireCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerManagerRetireCall) Return() *PacketHandlerManagerRetireCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerManagerRetireCall) Do(f func(protocol.ConnectionID)) *PacketHandlerManagerRetireCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerManagerRetireCall) DoAndReturn(f func(protocol.ConnectionID)) *PacketHandlerManagerRetireCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
