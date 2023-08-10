// database/mongo.go
package database

import (
	"AddressListener/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var username string
var password string
var clusterURL string

func InitMongoDB() (*mongo.Client, error) {
	// Load environment variables from .env file
	err := formaEnvVariables("USERNAME", "PASSWORD", "CLUSTER_URL")
	// Build MongoDB connection string dynamically
	mongoURI := createUri(username, password, clusterURL)
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func Push2Mongo(client *mongo.Client, data models.WebhookData) error {

	database := client.Database("your_database_name2")
	collection := database.Collection("your_collection_name")

	// Insert the webhook data into MongoDB
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Errorf("Error inserting data into MongoDB: ", err)
		return err
	}
	return nil
}
