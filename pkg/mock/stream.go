// Code generated by MockGen. DO NOT EDIT.
// Source: ../types/stream.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "mosn.io/api"
	types "mosn.io/mosn/pkg/types"
	buffer "mosn.io/pkg/buffer"
	reflect "reflect"
)

// MockStream is a mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockStream) ID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockStreamMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockStream)(nil).ID))
}

// AddEventListener mocks base method
func (m *MockStream) AddEventListener(streamEventListener types.StreamEventListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEventListener", streamEventListener)
}

// AddEventListener indicates an expected call of AddEventListener
func (mr *MockStreamMockRecorder) AddEventListener(streamEventListener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventListener", reflect.TypeOf((*MockStream)(nil).AddEventListener), streamEventListener)
}

// RemoveEventListener mocks base method
func (m *MockStream) RemoveEventListener(streamEventListener types.StreamEventListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveEventListener", streamEventListener)
}

// RemoveEventListener indicates an expected call of RemoveEventListener
func (mr *MockStreamMockRecorder) RemoveEventListener(streamEventListener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEventListener", reflect.TypeOf((*MockStream)(nil).RemoveEventListener), streamEventListener)
}

// ResetStream mocks base method
func (m *MockStream) ResetStream(reason types.StreamResetReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetStream", reason)
}

// ResetStream indicates an expected call of ResetStream
func (mr *MockStreamMockRecorder) ResetStream(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStream", reflect.TypeOf((*MockStream)(nil).ResetStream), reason)
}

// DestroyStream mocks base method
func (m *MockStream) DestroyStream() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DestroyStream")
}

// DestroyStream indicates an expected call of DestroyStream
func (mr *MockStreamMockRecorder) DestroyStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyStream", reflect.TypeOf((*MockStream)(nil).DestroyStream))
}

// MockStreamEventListener is a mock of StreamEventListener interface
type MockStreamEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockStreamEventListenerMockRecorder
}

// MockStreamEventListenerMockRecorder is the mock recorder for MockStreamEventListener
type MockStreamEventListenerMockRecorder struct {
	mock *MockStreamEventListener
}

// NewMockStreamEventListener creates a new mock instance
func NewMockStreamEventListener(ctrl *gomock.Controller) *MockStreamEventListener {
	mock := &MockStreamEventListener{ctrl: ctrl}
	mock.recorder = &MockStreamEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamEventListener) EXPECT() *MockStreamEventListenerMockRecorder {
	return m.recorder
}

// OnResetStream mocks base method
func (m *MockStreamEventListener) OnResetStream(reason types.StreamResetReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnResetStream", reason)
}

// OnResetStream indicates an expected call of OnResetStream
func (mr *MockStreamEventListenerMockRecorder) OnResetStream(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnResetStream", reflect.TypeOf((*MockStreamEventListener)(nil).OnResetStream), reason)
}

// OnDestroyStream mocks base method
func (m *MockStreamEventListener) OnDestroyStream() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDestroyStream")
}

// OnDestroyStream indicates an expected call of OnDestroyStream
func (mr *MockStreamEventListenerMockRecorder) OnDestroyStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDestroyStream", reflect.TypeOf((*MockStreamEventListener)(nil).OnDestroyStream))
}

// MockStreamSender is a mock of StreamSender interface
type MockStreamSender struct {
	ctrl     *gomock.Controller
	recorder *MockStreamSenderMockRecorder
}

// MockStreamSenderMockRecorder is the mock recorder for MockStreamSender
type MockStreamSenderMockRecorder struct {
	mock *MockStreamSender
}

// NewMockStreamSender creates a new mock instance
func NewMockStreamSender(ctrl *gomock.Controller) *MockStreamSender {
	mock := &MockStreamSender{ctrl: ctrl}
	mock.recorder = &MockStreamSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamSender) EXPECT() *MockStreamSenderMockRecorder {
	return m.recorder
}

// AppendHeaders mocks base method
func (m *MockStreamSender) AppendHeaders(ctx context.Context, headers api.HeaderMap, endStream bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendHeaders", ctx, headers, endStream)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendHeaders indicates an expected call of AppendHeaders
func (mr *MockStreamSenderMockRecorder) AppendHeaders(ctx, headers, endStream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendHeaders", reflect.TypeOf((*MockStreamSender)(nil).AppendHeaders), ctx, headers, endStream)
}

// AppendData mocks base method
func (m *MockStreamSender) AppendData(ctx context.Context, data buffer.IoBuffer, endStream bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendData", ctx, data, endStream)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendData indicates an expected call of AppendData
func (mr *MockStreamSenderMockRecorder) AppendData(ctx, data, endStream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendData", reflect.TypeOf((*MockStreamSender)(nil).AppendData), ctx, data, endStream)
}

// AppendTrailers mocks base method
func (m *MockStreamSender) AppendTrailers(ctx context.Context, trailers api.HeaderMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendTrailers", ctx, trailers)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendTrailers indicates an expected call of AppendTrailers
func (mr *MockStreamSenderMockRecorder) AppendTrailers(ctx, trailers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendTrailers", reflect.TypeOf((*MockStreamSender)(nil).AppendTrailers), ctx, trailers)
}

// GetStream mocks base method
func (m *MockStreamSender) GetStream() types.Stream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStream")
	ret0, _ := ret[0].(types.Stream)
	return ret0
}

// GetStream indicates an expected call of GetStream
func (mr *MockStreamSenderMockRecorder) GetStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStream", reflect.TypeOf((*MockStreamSender)(nil).GetStream))
}

// MockStreamReceiveListener is a mock of StreamReceiveListener interface
type MockStreamReceiveListener struct {
	ctrl     *gomock.Controller
	recorder *MockStreamReceiveListenerMockRecorder
}

// MockStreamReceiveListenerMockRecorder is the mock recorder for MockStreamReceiveListener
type MockStreamReceiveListenerMockRecorder struct {
	mock *MockStreamReceiveListener
}

// NewMockStreamReceiveListener creates a new mock instance
func NewMockStreamReceiveListener(ctrl *gomock.Controller) *MockStreamReceiveListener {
	mock := &MockStreamReceiveListener{ctrl: ctrl}
	mock.recorder = &MockStreamReceiveListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamReceiveListener) EXPECT() *MockStreamReceiveListenerMockRecorder {
	return m.recorder
}

// OnReceive mocks base method
func (m *MockStreamReceiveListener) OnReceive(ctx context.Context, headers api.HeaderMap, data buffer.IoBuffer, trailers api.HeaderMap) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnReceive", ctx, headers, data, trailers)
}

// OnReceive indicates an expected call of OnReceive
func (mr *MockStreamReceiveListenerMockRecorder) OnReceive(ctx, headers, data, trailers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnReceive", reflect.TypeOf((*MockStreamReceiveListener)(nil).OnReceive), ctx, headers, data, trailers)
}

// OnDecodeError mocks base method
func (m *MockStreamReceiveListener) OnDecodeError(ctx context.Context, err error, headers api.HeaderMap) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDecodeError", ctx, err, headers)
}

// OnDecodeError indicates an expected call of OnDecodeError
func (mr *MockStreamReceiveListenerMockRecorder) OnDecodeError(ctx, err, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDecodeError", reflect.TypeOf((*MockStreamReceiveListener)(nil).OnDecodeError), ctx, err, headers)
}

// MockStreamConnection is a mock of StreamConnection interface
type MockStreamConnection struct {
	ctrl     *gomock.Controller
	recorder *MockStreamConnectionMockRecorder
}

// MockStreamConnectionMockRecorder is the mock recorder for MockStreamConnection
type MockStreamConnectionMockRecorder struct {
	mock *MockStreamConnection
}

// NewMockStreamConnection creates a new mock instance
func NewMockStreamConnection(ctrl *gomock.Controller) *MockStreamConnection {
	mock := &MockStreamConnection{ctrl: ctrl}
	mock.recorder = &MockStreamConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamConnection) EXPECT() *MockStreamConnectionMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockStreamConnection) Dispatch(buffer buffer.IoBuffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispatch", buffer)
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockStreamConnectionMockRecorder) Dispatch(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockStreamConnection)(nil).Dispatch), buffer)
}

// Protocol mocks base method
func (m *MockStreamConnection) Protocol() api.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(api.Protocol)
	return ret0
}

// Protocol indicates an expected call of Protocol
func (mr *MockStreamConnectionMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockStreamConnection)(nil).Protocol))
}

// ActiveStreamsNum mocks base method
func (m *MockStreamConnection) ActiveStreamsNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveStreamsNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// ActiveStreamsNum indicates an expected call of ActiveStreamsNum
func (mr *MockStreamConnectionMockRecorder) ActiveStreamsNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveStreamsNum", reflect.TypeOf((*MockStreamConnection)(nil).ActiveStreamsNum))
}

// GoAway mocks base method
func (m *MockStreamConnection) GoAway() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GoAway")
}

// GoAway indicates an expected call of GoAway
func (mr *MockStreamConnectionMockRecorder) GoAway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoAway", reflect.TypeOf((*MockStreamConnection)(nil).GoAway))
}

// Reset mocks base method
func (m *MockStreamConnection) Reset(reason types.StreamResetReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", reason)
}

// Reset indicates an expected call of Reset
func (mr *MockStreamConnectionMockRecorder) Reset(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockStreamConnection)(nil).Reset), reason)
}

// CheckReasonError mocks base method
func (m *MockStreamConnection) CheckReasonError(connected bool, event api.ConnectionEvent) (types.StreamResetReason, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReasonError", connected, event)
	ret0, _ := ret[0].(types.StreamResetReason)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckReasonError indicates an expected call of CheckReasonError
func (mr *MockStreamConnectionMockRecorder) CheckReasonError(connected, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReasonError", reflect.TypeOf((*MockStreamConnection)(nil).CheckReasonError), connected, event)
}

// MockServerStreamConnection is a mock of ServerStreamConnection interface
type MockServerStreamConnection struct {
	ctrl     *gomock.Controller
	recorder *MockServerStreamConnectionMockRecorder
}

// MockServerStreamConnectionMockRecorder is the mock recorder for MockServerStreamConnection
type MockServerStreamConnectionMockRecorder struct {
	mock *MockServerStreamConnection
}

// NewMockServerStreamConnection creates a new mock instance
func NewMockServerStreamConnection(ctrl *gomock.Controller) *MockServerStreamConnection {
	mock := &MockServerStreamConnection{ctrl: ctrl}
	mock.recorder = &MockServerStreamConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerStreamConnection) EXPECT() *MockServerStreamConnectionMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockServerStreamConnection) Dispatch(buffer buffer.IoBuffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispatch", buffer)
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockServerStreamConnectionMockRecorder) Dispatch(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockServerStreamConnection)(nil).Dispatch), buffer)
}

// Protocol mocks base method
func (m *MockServerStreamConnection) Protocol() api.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(api.Protocol)
	return ret0
}

// Protocol indicates an expected call of Protocol
func (mr *MockServerStreamConnectionMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockServerStreamConnection)(nil).Protocol))
}

// ActiveStreamsNum mocks base method
func (m *MockServerStreamConnection) ActiveStreamsNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveStreamsNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// ActiveStreamsNum indicates an expected call of ActiveStreamsNum
func (mr *MockServerStreamConnectionMockRecorder) ActiveStreamsNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveStreamsNum", reflect.TypeOf((*MockServerStreamConnection)(nil).ActiveStreamsNum))
}

// GoAway mocks base method
func (m *MockServerStreamConnection) GoAway() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GoAway")
}

// GoAway indicates an expected call of GoAway
func (mr *MockServerStreamConnectionMockRecorder) GoAway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoAway", reflect.TypeOf((*MockServerStreamConnection)(nil).GoAway))
}

// Reset mocks base method
func (m *MockServerStreamConnection) Reset(reason types.StreamResetReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", reason)
}

// Reset indicates an expected call of Reset
func (mr *MockServerStreamConnectionMockRecorder) Reset(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockServerStreamConnection)(nil).Reset), reason)
}

// CheckReasonError mocks base method
func (m *MockServerStreamConnection) CheckReasonError(connected bool, event api.ConnectionEvent) (types.StreamResetReason, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReasonError", connected, event)
	ret0, _ := ret[0].(types.StreamResetReason)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckReasonError indicates an expected call of CheckReasonError
func (mr *MockServerStreamConnectionMockRecorder) CheckReasonError(connected, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReasonError", reflect.TypeOf((*MockServerStreamConnection)(nil).CheckReasonError), connected, event)
}

// MockClientStreamConnection is a mock of ClientStreamConnection interface
type MockClientStreamConnection struct {
	ctrl     *gomock.Controller
	recorder *MockClientStreamConnectionMockRecorder
}

// MockClientStreamConnectionMockRecorder is the mock recorder for MockClientStreamConnection
type MockClientStreamConnectionMockRecorder struct {
	mock *MockClientStreamConnection
}

// NewMockClientStreamConnection creates a new mock instance
func NewMockClientStreamConnection(ctrl *gomock.Controller) *MockClientStreamConnection {
	mock := &MockClientStreamConnection{ctrl: ctrl}
	mock.recorder = &MockClientStreamConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientStreamConnection) EXPECT() *MockClientStreamConnectionMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockClientStreamConnection) Dispatch(buffer buffer.IoBuffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispatch", buffer)
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockClientStreamConnectionMockRecorder) Dispatch(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockClientStreamConnection)(nil).Dispatch), buffer)
}

// Protocol mocks base method
func (m *MockClientStreamConnection) Protocol() api.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(api.Protocol)
	return ret0
}

// Protocol indicates an expected call of Protocol
func (mr *MockClientStreamConnectionMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockClientStreamConnection)(nil).Protocol))
}

// ActiveStreamsNum mocks base method
func (m *MockClientStreamConnection) ActiveStreamsNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveStreamsNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// ActiveStreamsNum indicates an expected call of ActiveStreamsNum
func (mr *MockClientStreamConnectionMockRecorder) ActiveStreamsNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveStreamsNum", reflect.TypeOf((*MockClientStreamConnection)(nil).ActiveStreamsNum))
}

// GoAway mocks base method
func (m *MockClientStreamConnection) GoAway() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GoAway")
}

// GoAway indicates an expected call of GoAway
func (mr *MockClientStreamConnectionMockRecorder) GoAway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoAway", reflect.TypeOf((*MockClientStreamConnection)(nil).GoAway))
}

// Reset mocks base method
func (m *MockClientStreamConnection) Reset(reason types.StreamResetReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", reason)
}

// Reset indicates an expected call of Reset
func (mr *MockClientStreamConnectionMockRecorder) Reset(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockClientStreamConnection)(nil).Reset), reason)
}

// CheckReasonError mocks base method
func (m *MockClientStreamConnection) CheckReasonError(connected bool, event api.ConnectionEvent) (types.StreamResetReason, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReasonError", connected, event)
	ret0, _ := ret[0].(types.StreamResetReason)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckReasonError indicates an expected call of CheckReasonError
func (mr *MockClientStreamConnectionMockRecorder) CheckReasonError(connected, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReasonError", reflect.TypeOf((*MockClientStreamConnection)(nil).CheckReasonError), connected, event)
}

// NewStream mocks base method
func (m *MockClientStreamConnection) NewStream(ctx context.Context, responseReceiveListener types.StreamReceiveListener) types.StreamSender {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStream", ctx, responseReceiveListener)
	ret0, _ := ret[0].(types.StreamSender)
	return ret0
}

// NewStream indicates an expected call of NewStream
func (mr *MockClientStreamConnectionMockRecorder) NewStream(ctx, responseReceiveListener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*MockClientStreamConnection)(nil).NewStream), ctx, responseReceiveListener)
}

// MockStreamConnectionEventListener is a mock of StreamConnectionEventListener interface
type MockStreamConnectionEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockStreamConnectionEventListenerMockRecorder
}

// MockStreamConnectionEventListenerMockRecorder is the mock recorder for MockStreamConnectionEventListener
type MockStreamConnectionEventListenerMockRecorder struct {
	mock *MockStreamConnectionEventListener
}

// NewMockStreamConnectionEventListener creates a new mock instance
func NewMockStreamConnectionEventListener(ctrl *gomock.Controller) *MockStreamConnectionEventListener {
	mock := &MockStreamConnectionEventListener{ctrl: ctrl}
	mock.recorder = &MockStreamConnectionEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamConnectionEventListener) EXPECT() *MockStreamConnectionEventListenerMockRecorder {
	return m.recorder
}

// OnGoAway mocks base method
func (m *MockStreamConnectionEventListener) OnGoAway() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnGoAway")
}

// OnGoAway indicates an expected call of OnGoAway
func (mr *MockStreamConnectionEventListenerMockRecorder) OnGoAway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnGoAway", reflect.TypeOf((*MockStreamConnectionEventListener)(nil).OnGoAway))
}

// MockServerStreamConnectionEventListener is a mock of ServerStreamConnectionEventListener interface
type MockServerStreamConnectionEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockServerStreamConnectionEventListenerMockRecorder
}

// MockServerStreamConnectionEventListenerMockRecorder is the mock recorder for MockServerStreamConnectionEventListener
type MockServerStreamConnectionEventListenerMockRecorder struct {
	mock *MockServerStreamConnectionEventListener
}

// NewMockServerStreamConnectionEventListener creates a new mock instance
func NewMockServerStreamConnectionEventListener(ctrl *gomock.Controller) *MockServerStreamConnectionEventListener {
	mock := &MockServerStreamConnectionEventListener{ctrl: ctrl}
	mock.recorder = &MockServerStreamConnectionEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerStreamConnectionEventListener) EXPECT() *MockServerStreamConnectionEventListenerMockRecorder {
	return m.recorder
}

// OnGoAway mocks base method
func (m *MockServerStreamConnectionEventListener) OnGoAway() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnGoAway")
}

// OnGoAway indicates an expected call of OnGoAway
func (mr *MockServerStreamConnectionEventListenerMockRecorder) OnGoAway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnGoAway", reflect.TypeOf((*MockServerStreamConnectionEventListener)(nil).OnGoAway))
}

// NewStreamDetect mocks base method
func (m *MockServerStreamConnectionEventListener) NewStreamDetect(context context.Context, sender types.StreamSender, span types.Span) types.StreamReceiveListener {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStreamDetect", context, sender, span)
	ret0, _ := ret[0].(types.StreamReceiveListener)
	return ret0
}

// NewStreamDetect indicates an expected call of NewStreamDetect
func (mr *MockServerStreamConnectionEventListenerMockRecorder) NewStreamDetect(context, sender, span interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStreamDetect", reflect.TypeOf((*MockServerStreamConnectionEventListener)(nil).NewStreamDetect), context, sender, span)
}

// MockConnectionPool is a mock of ConnectionPool interface
type MockConnectionPool struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionPoolMockRecorder
}

// MockConnectionPoolMockRecorder is the mock recorder for MockConnectionPool
type MockConnectionPoolMockRecorder struct {
	mock *MockConnectionPool
}

// NewMockConnectionPool creates a new mock instance
func NewMockConnectionPool(ctrl *gomock.Controller) *MockConnectionPool {
	mock := &MockConnectionPool{ctrl: ctrl}
	mock.recorder = &MockConnectionPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnectionPool) EXPECT() *MockConnectionPoolMockRecorder {
	return m.recorder
}

// Protocol mocks base method
func (m *MockConnectionPool) Protocol() api.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(api.Protocol)
	return ret0
}

// Protocol indicates an expected call of Protocol
func (mr *MockConnectionPoolMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockConnectionPool)(nil).Protocol))
}

// NewStream mocks base method
func (m *MockConnectionPool) NewStream(ctx context.Context, receiver types.StreamReceiveListener, listener types.PoolEventListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewStream", ctx, receiver, listener)
}

// NewStream indicates an expected call of NewStream
func (mr *MockConnectionPoolMockRecorder) NewStream(ctx, receiver, listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*MockConnectionPool)(nil).NewStream), ctx, receiver, listener)
}

// CheckAndInit mocks base method
func (m *MockConnectionPool) CheckAndInit(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndInit", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckAndInit indicates an expected call of CheckAndInit
func (mr *MockConnectionPoolMockRecorder) CheckAndInit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndInit", reflect.TypeOf((*MockConnectionPool)(nil).CheckAndInit), ctx)
}

// TLSHashValue mocks base method
func (m *MockConnectionPool) TLSHashValue() *types.HashValue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSHashValue")
	ret0, _ := ret[0].(*types.HashValue)
	return ret0
}

// TLSHashValue indicates an expected call of TLSHashValue
func (mr *MockConnectionPoolMockRecorder) TLSHashValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSHashValue", reflect.TypeOf((*MockConnectionPool)(nil).TLSHashValue))
}

// Shutdown mocks base method
func (m *MockConnectionPool) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockConnectionPoolMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockConnectionPool)(nil).Shutdown))
}

// Close mocks base method
func (m *MockConnectionPool) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockConnectionPoolMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnectionPool)(nil).Close))
}

// Host mocks base method
func (m *MockConnectionPool) Host() types.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(types.Host)
	return ret0
}

// Host indicates an expected call of Host
func (mr *MockConnectionPoolMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockConnectionPool)(nil).Host))
}

// UpdateHost mocks base method
func (m *MockConnectionPool) UpdateHost(arg0 types.Host) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateHost", arg0)
}

// UpdateHost indicates an expected call of UpdateHost
func (mr *MockConnectionPoolMockRecorder) UpdateHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHost", reflect.TypeOf((*MockConnectionPool)(nil).UpdateHost), arg0)
}

// MockPoolEventListener is a mock of PoolEventListener interface
type MockPoolEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockPoolEventListenerMockRecorder
}

// MockPoolEventListenerMockRecorder is the mock recorder for MockPoolEventListener
type MockPoolEventListenerMockRecorder struct {
	mock *MockPoolEventListener
}

// NewMockPoolEventListener creates a new mock instance
func NewMockPoolEventListener(ctrl *gomock.Controller) *MockPoolEventListener {
	mock := &MockPoolEventListener{ctrl: ctrl}
	mock.recorder = &MockPoolEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPoolEventListener) EXPECT() *MockPoolEventListenerMockRecorder {
	return m.recorder
}

// OnFailure mocks base method
func (m *MockPoolEventListener) OnFailure(reason types.PoolFailureReason, host types.Host) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnFailure", reason, host)
}

// OnFailure indicates an expected call of OnFailure
func (mr *MockPoolEventListenerMockRecorder) OnFailure(reason, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFailure", reflect.TypeOf((*MockPoolEventListener)(nil).OnFailure), reason, host)
}

// OnReady mocks base method
func (m *MockPoolEventListener) OnReady(sender types.StreamSender, host types.Host) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnReady", sender, host)
}

// OnReady indicates an expected call of OnReady
func (mr *MockPoolEventListenerMockRecorder) OnReady(sender, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnReady", reflect.TypeOf((*MockPoolEventListener)(nil).OnReady), sender, host)
}
