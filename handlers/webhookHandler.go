package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WebhookHandler is responsible for handling incoming webhook requests
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Process the data (you can replace this with your actual processing logic)
	// For example, print the received data
	fmt.Printf("Received webhook data: %+v\n", data)

	// Respond to the sender
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Webhook received and processed")
}
