package main

import (
	"AddressListener/database"
	"AddressListener/handlers" // Update this to the correct package path
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	client, err := database.InitMongoDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Define your webhook endpoint and handler
	http.HandleFunc("/webhook-receiver", handlers.WebhookHandler(client))

	// Start the server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
