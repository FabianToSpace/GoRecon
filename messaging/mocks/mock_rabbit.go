// Code generated by MockGen. DO NOT EDIT.
// Source: .\rabbit_connection.go

// Package mock_messaging is a generated GoMock package.
package mock_messaging

import (
	reflect "reflect"

	messaging "github.com/FabianToSpace/GoRecon/messaging"
	gomock "github.com/golang/mock/gomock"
	amqp091_go "github.com/rabbitmq/amqp091-go"
)

// MockRabbitConnect is a mock of RabbitConnect interface.
type MockRabbitConnect struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitConnectMockRecorder
}

// MockRabbitConnectMockRecorder is the mock recorder for MockRabbitConnect.
type MockRabbitConnectMockRecorder struct {
	mock *MockRabbitConnect
}

// NewMockRabbitConnect creates a new mock instance.
func NewMockRabbitConnect(ctrl *gomock.Controller) *MockRabbitConnect {
	mock := &MockRabbitConnect{ctrl: ctrl}
	mock.recorder = &MockRabbitConnectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitConnect) EXPECT() *MockRabbitConnectMockRecorder {
	return m.recorder
}

// ChannelConnect mocks base method.
func (m *MockRabbitConnect) ChannelConnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelConnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// ChannelConnect indicates an expected call of ChannelConnect.
func (mr *MockRabbitConnectMockRecorder) ChannelConnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelConnect", reflect.TypeOf((*MockRabbitConnect)(nil).ChannelConnect))
}

// Connect mocks base method.
func (m *MockRabbitConnect) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockRabbitConnectMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRabbitConnect)(nil).Connect))
}

// Consume mocks base method.
func (m *MockRabbitConnect) Consume() (<-chan amqp091_go.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume")
	ret0, _ := ret[0].(<-chan amqp091_go.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockRabbitConnectMockRecorder) Consume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockRabbitConnect)(nil).Consume))
}

// PublishMessage mocks base method.
func (m *MockRabbitConnect) PublishMessage(arg0 messaging.ServiceMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishMessage indicates an expected call of PublishMessage.
func (mr *MockRabbitConnectMockRecorder) PublishMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMessage", reflect.TypeOf((*MockRabbitConnect)(nil).PublishMessage), arg0)
}

// QueueConnect mocks base method.
func (m *MockRabbitConnect) QueueConnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueConnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueConnect indicates an expected call of QueueConnect.
func (mr *MockRabbitConnectMockRecorder) QueueConnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueConnect", reflect.TypeOf((*MockRabbitConnect)(nil).QueueConnect))
}
