package main

import (
	"favorite-events-poc/consumer/consumer"
	"log"
)

func main() {
	log.Println("Kafka consumer started")
	consumer.StartConsumer()
}
