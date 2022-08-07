package queueing

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConsumer.bootStrapServer,
		"group.id":          "kafka-consumer-group-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	msg, err := c.ReadMessage(-1)

	c.Close()

	message := string(msg.Value)

	return &message, err

}
