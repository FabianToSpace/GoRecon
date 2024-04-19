package messaging

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConsumer() error {
	Init()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "test",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		Logger.Error("Kafka Consumer", "", "Failed to create consumer: "+err.Error())
		return err
	}

	c.SubscribeTopics([]string{"test"}, nil)

	for {
		msg, err := c.ReadMessage(100 * time.Millisecond)
		if err == nil {
			Logger.Debug("Kafka Consumer", "", fmt.Sprintf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value)))
		} else {
			Logger.Error("Kafka Consumer", "", fmt.Sprintf("Consumer error: %v (%v)\n", err, msg))
			break
		}
	}

	c.Close()
	return nil
}
