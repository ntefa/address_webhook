package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ntefa/address_webhook/database"
	"github.com/ntefa/address_webhook/handlers" // Update this to the correct package path
	"github.com/ntefa/address_webhook/lib"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	err := lib.FormatEnvVariables("USERNAME", "PASSWORD", "CLUSTER_URL", "DBNAME", "COLLECTIONNAME", "WEBHOOKURL")

	client, err := database.InitMongoDB(lib.Username, lib.Password, lib.ClusterURL)
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
	http.ListenAndServe(":8080", nil) //Note: ListenAndServe is a blocking call, if needed to receive on different ports make use of goroutines.
}
