package queueing

import (
	"context"
	"os"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	ConsumeMessage(string) (*string, error)
}

type kafkaConsumer struct {
	bootStrapServer string
}

func NewKafkaConsumer() *kafkaConsumer {
	bootStrapServer := os.Getenv("BOOTSTRAP_SERVERS")
	if bootStrapServer == "" {
		bootStrapServer = "localhost:29092"
	}
	return &kafkaConsumer{
		bootStrapServer: bootStrapServer,
	}
}

func (kafkaConsumer *kafkaConsumer) ConsumeMessage(topic string) (*string, error) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{kafkaConsumer.bootStrapServer},
		Topic:       topic,
		GroupID:     "kafka-consumer-group-1",
		StartOffset: kafka.FirstOffset,
	})

	m, err := r.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	message := string(m.Value)
	r.Close()
	return &message, nil

}
