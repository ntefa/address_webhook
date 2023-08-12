package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ntefa/address_webhook/database"
	"github.com/ntefa/address_webhook/models" // Update this to the correct package path

	"go.mongodb.org/mongo-driver/mongo"
)

func WebhookHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost { // Make sure the request method is POST
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the incoming JSON data
		var data models.WebhookData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// Insert the webhook data into MongoDB
		err = database.Push2Mongo(client, data)
		if err != nil {
			http.Error(w, "Error inserting data into MongoDB", http.StatusInternalServerError)
			return
		}

		// Respond to the sender
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Webhook received and processed")
	}
}
