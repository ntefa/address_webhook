package handlers

import (
	"AddressListener/models" // Update this to the correct package path
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func WebhookHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Make sure the request method is POST
		if r.Method != http.MethodPost {
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

		database := client.Database("your_database_name2")
		collection := database.Collection("your_collection_name")

		// Insert the webhook data into MongoDB
		_, err = collection.InsertOne(context.Background(), data)
		if err != nil {
			log.Println("Error inserting data into MongoDB:", err)
			http.Error(w, "Error inserting data into MongoDB", http.StatusInternalServerError)
			return
		}
		// // Process the data
		// fmt.Printf("Received webhook ID: %s\n", data.WebhookID)
		// fmt.Printf("Received event type: %s\n", data.Type)
		// fmt.Printf("Received event timestamp: %s\n", data.CreatedAt)
		// fmt.Printf("Received event network: %s\n", data.Event.Network)

		// for _, activity := range data.Event.Activity {
		// 	fmt.Printf("Received activity block number: %s\n", activity.BlockNum)
		// 	fmt.Printf("Received activity hash: %s\n", activity.Hash)
		// 	fmt.Printf("Received activity from address: %s\n", activity.FromAddress)
		// 	fmt.Printf("Received activity to address: %s\n", activity.ToAddress)
		// 	// ... Print other activity details ...
		// }

		// Respond to the sender
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Webhook received and processed")
	}
}
