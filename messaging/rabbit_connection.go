package messaging

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConnection struct {
	Queue          *amqp.Queue
	Channel        *amqp.Channel
	Connection     *amqp.Connection
	QueueInfo      QueueInfo
	ConnectionInfo ConnectionInfo
}

type RabbitConnect interface {
	Connect() error
	ChannelConnect() error
	QueueConnect() error
	PublishMessage([]byte) error
	Consume() (<-chan amqp.Delivery, error)
	DeclareExchange() error
}

type ConnectionInfo struct {
	User     string
	Password string
	Host     string
	Port     string
}

type QueueInfo struct {
	QueueName    string
	ExchangeName string
	Durable      bool
	AutoDelete   bool
	Exclusive    bool
	NoWait       bool
	Arguments    amqp.Table
}

func (r *RabbitConnection) Connect() error {
	connectionLink := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		r.ConnectionInfo.User,
		r.ConnectionInfo.Password,
		r.ConnectionInfo.Host,
		r.ConnectionInfo.Port)
	conn, err := amqp.Dial(connectionLink)
	if err != nil {
		return err
	}

	r.Connection = conn

	return nil
}

func (r *RabbitConnection) ChannelConnect() error {
	ch, err := r.Connection.Channel()
	if err != nil {
		return err
	}

	r.Channel = ch

	return nil
}

func (r *RabbitConnection) QueueConnect() error {
	q, err := r.Channel.QueueDeclare(
		r.QueueInfo.QueueName,  // name
		r.QueueInfo.Durable,    // durable
		r.QueueInfo.AutoDelete, // delete when unused
		r.QueueInfo.Exclusive,  // exclusive
		r.QueueInfo.NoWait,     // no-wait
		r.QueueInfo.Arguments,  // arguments
	)
	if err != nil {
		return err
	}

	r.Queue = &q

	return nil
}

func (r *RabbitConnection) PublishMessage(message []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := message

	err := r.Channel.PublishWithContext(ctx,
		"",           // exchange
		r.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitConnection) Consume() (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		r.Queue.Name, // queue
		"",           // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)
}

func (r *RabbitConnection) DeclareExchange() error {
	err := r.Channel.ExchangeDeclare(
		r.QueueInfo.ExchangeName, // name
		"fanout",                 // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		return err
	}

	return nil
}
