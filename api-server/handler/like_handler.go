package handler

import (
	"encoding/json"
	"favorite-events-poc/api-server/model"
	"favorite-events-poc/api-server/producer"
	"fmt"
	"log"
	"net/http"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method Not Allowed")
		return
	}

	var event model.LikeEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid JSON")
		return
	}

	log.Printf("Received like event: %+v\\n", event)
	producer.ProduceLikeEvent(event)

	fmt.Fprint(w, "Like event received!")
}
