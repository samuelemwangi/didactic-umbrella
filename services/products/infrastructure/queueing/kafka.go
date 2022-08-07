package queueing

import (
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaProducer interface {
	ProduceMessage(string, string)
}

type kafkaProducer struct {
	bootStrapServer string
}

func NewKafkaProducer() *kafkaProducer {
	bootStrapServer := os.Getenv("BOOTSTRAP_SERVERS")
	if bootStrapServer == "" {
		bootStrapServer = "localhost:29092"
	}
	return &kafkaProducer{
		bootStrapServer: bootStrapServer,
	}
}

func (kafkaPproducer *kafkaProducer) ProduceMessage(topic string, message string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaPproducer.bootStrapServer})
	if err != nil {
		return
	}

	defer p.Close()

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(message),
	}, nil)

	p.Flush(15 * 1000)

}
