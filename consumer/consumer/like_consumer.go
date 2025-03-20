package consumer

import (
	"context"
	"encoding/json"
	"log"

	"favorite-events-poc/api-server/model"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	InitDB()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "like-consumer-group",
		Topic:   "like_events",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var event model.LikeEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		log.Printf("Consumed message: %+v\n", event)
		insertLikeEvent(event)
	}
}

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "app:password@tcp(localhost:3306)/favorite_db")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}
	log.Println("Connected to MySQL")
}

func insertLikeEvent(event model.LikeEvent) {
	_, err := db.Exec("INSERT INTO likes (user_id, item_id) VALUES (?, ?)", event.UserID, event.ItemID)
	if err != nil {
		log.Println("Error inserting into DB:", err)
	} else {
		log.Printf("Inserted like event into DB: %+v\n", event)
	}
}
