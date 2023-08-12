// database/mongo.go
package database

import (
	"context"
	"fmt"
	"log"

	"github.com/ntefa/address_webhook/lib"
	"github.com/ntefa/address_webhook/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(username string, password string, clusterURL string) (*mongo.Client, error) {

	// Build MongoDB connection string dynamically
	mongoURI := lib.CreateUri(username, password, clusterURL)
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

func Push2Mongo(client *mongo.Client, data models.WebhookData, dbName string, collectionName string) error {

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	// Insert the webhook data into MongoDB
	_, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return fmt.Errorf("Error inserting data into MongoDB: %v", err)
	}

	// Log a message after successful insertion
	log.Printf("Inserted data into MongoDB!")

	return nil
}
