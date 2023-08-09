package main

import (
	"AddressListener/handlers" // Update this to the correct package path
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB connection details from environment variables
	username := url.QueryEscape(os.Getenv("USERNAME"))
	password := url.QueryEscape(os.Getenv("PASSWORD"))
	clusterURL := os.Getenv("CLUSTER_URL")

	// Build MongoDB connection string dynamically
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", username, password, clusterURL)
	fmt.Println(mongoURI)
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
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
	http.HandleFunc("/webhook-receiver", handlers.WebhookHandler)

	// Start the server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
