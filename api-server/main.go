package main

import (
	"favorite-events-poc/api-server/handler"
	"favorite-events-poc/api-server/producer"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/like", handler.LikeHandler)
	producer.InitKafkaWriter("localhost:9092")
	log.Println("API server started on :8080")
	http.ListenAndServe(":8080", nil)
}
