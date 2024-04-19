package messaging

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaProducer() error {
	Init()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		Logger.Error("Kafka Producer", "", "Failed to create producer: "+err.Error())
		return err
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					Logger.Error("Kafka Producer", "", fmt.Sprintf("Delivery failed: %v\n", ev.TopicPartition.Error))
				} else {
					Logger.Debug("Kafka Producer", "", fmt.Sprintf("Delivered message to %v\n", ev.TopicPartition))
				}
			}
		}
	}()

	topic := "test"
	for i := 0; i < 10; i++ {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("Message %d", i)),
		}, nil)
	}

	p.Flush(15 * 1000)
	return nil
}
