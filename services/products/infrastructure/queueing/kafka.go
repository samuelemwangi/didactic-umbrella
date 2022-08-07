package queueing

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
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
	w := &kafka.Writer{
		Addr:                   kafka.TCP(kafkaPproducer.bootStrapServer),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	err := w.WriteMessages(context.Background(), kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Println()
		log.Println(err)
		log.Println()
	}
	w.Close()

}
