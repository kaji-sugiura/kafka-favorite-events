package producer

import (
	"context"
	"encoding/json"
	"favorite-events-poc/api-server/model"
	"log"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitKafkaWriter(brokerAddress string) {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    "like_events",
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("Kafka writer initialized")
}

func ProduceLikeEvent(event model.LikeEvent) {
	msg, err := json.Marshal(event)
	if err != nil {
		log.Println("Failed to marshal event:", err)
		return
	}

	err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	})

	if err != nil {
		log.Println("Failed to produce message to Kafka:", err)
	} else {
		log.Println("Produced message to Kafka:", string(msg))
	}
}
