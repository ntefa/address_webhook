package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ntefa/address_webhook/database"
	lib "github.com/ntefa/address_webhook/lib"
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
		err = database.Push2Mongo(client, data, lib.DBName, lib.CollectionName)
		if err != nil {
			http.Error(w, "Error inserting data into MongoDB", http.StatusInternalServerError)
			return
		}
		// sendDiscordNotification()

		// Respond to the sender
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Webhook received and processed")
	}
}

// func sendDiscordNotification() {
// 	webhookURL := "YOUR_DISCORD_WEBHOOK_URL"
// 	notificationMessage := "New data received and processed!"

// 	payload := map[string]interface{}{
// 		"content": notificationMessage,
// 	}

// 	body, _ := json.Marshal(payload)
// 	http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
// }
