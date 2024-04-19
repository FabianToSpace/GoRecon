package messaging_test

import (
	"errors"
	"testing"

	mock_messaging "github.com/FabianToSpace/GoRecon/messaging/mocks"
	"github.com/golang/mock/gomock"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestRabbitConnection_Connect(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mock_connection := mock_messaging.NewMockRabbitConnect(mc)

	t.Run("Successful Connection", func(t *testing.T) {
		mock_connection.EXPECT().Connect().Return(nil)
		err := mock_connection.Connect()
		assert.NoError(t, err)
	})

	t.Run("Failed Connection", func(t *testing.T) {
		// Modify the connection URL to intentionally fail the connection
		mock_connection.EXPECT().Connect().Return(errors.New("Error"))
		err := mock_connection.Connect()
		assert.Error(t, err)
	})
}

func TestRabbitConnection_ChannelConnect(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mock_connection := mock_messaging.NewMockRabbitConnect(mc)

	t.Run("Successful Channel Connection", func(t *testing.T) {
		mock_connection.EXPECT().ChannelConnect().Return(nil)
		err := mock_connection.ChannelConnect()
		assert.NoError(t, err)
	})

	t.Run("Failed Channel Connection", func(t *testing.T) {
		// Modify the connection URL to intentionally fail the connection
		mock_connection.EXPECT().ChannelConnect().Return(errors.New("Error"))
		err := mock_connection.ChannelConnect()
		assert.Error(t, err)
	})
}

func TestRabbitConnection_QueueConnect(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mock_connection := mock_messaging.NewMockRabbitConnect(mc)

	t.Run("Successful Queue Connection", func(t *testing.T) {
		mock_connection.EXPECT().QueueConnect().Return(nil)
		err := mock_connection.QueueConnect()
		assert.NoError(t, err)
	})

	t.Run("Failed Queue Connection", func(t *testing.T) {
		// Modify the connection URL to intentionally fail the connection
		mock_connection.EXPECT().QueueConnect().Return(errors.New("Error"))
		err := mock_connection.QueueConnect()
		assert.Error(t, err)
	})
}

func TestRabbitConnection_PublishMessage(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mock_connection := mock_messaging.NewMockRabbitConnect(mc)

	t.Run("Successful Publish Message", func(t *testing.T) {
		mock_connection.EXPECT().PublishMessage(gomock.Any()).Return(nil)
		err := mock_connection.PublishMessage([]byte(""))
		assert.NoError(t, err)
	})

	t.Run("Failed Publish Message", func(t *testing.T) {
		// Modify the connection URL to intentionally fail the connection
		mock_connection.EXPECT().PublishMessage(gomock.Any()).Return(errors.New("Error"))
		err := mock_connection.PublishMessage([]byte(""))
		assert.Error(t, err)
	})
}

func TestRabbitConnection_Consume(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mock_connection := mock_messaging.NewMockRabbitConnect(mc)

	t.Run("Successful Consume", func(t *testing.T) {
		returnChan := make(chan amqp.Delivery)
		mock_connection.EXPECT().Consume().Return(returnChan, nil)
		channel, err := mock_connection.Consume()

		assert.NotNil(t, channel)
		assert.Empty(t, channel)
		assert.NoError(t, err)
	})

	t.Run("Failed Consume", func(t *testing.T) {
		mock_connection.EXPECT().Consume().Return(nil, errors.New("Error"))
		_, err := mock_connection.Consume()
		assert.Error(t, err)
	})
}

func TestRabbitConnection_DeclareExchange(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mock_connection := mock_messaging.NewMockRabbitConnect(mc)

	t.Run("Successful Declare Exchange", func(t *testing.T) {
		mock_connection.EXPECT().DeclareExchange().Return(nil)
		err := mock_connection.DeclareExchange()
		assert.NoError(t, err)
	})

	t.Run("Failed Declare Exchange", func(t *testing.T) {
		// Modify the connection URL to intentionally fail the connection
		mock_connection.EXPECT().DeclareExchange().Return(errors.New("Error"))
		err := mock_connection.DeclareExchange()
		assert.Error(t, err)
	})
}
