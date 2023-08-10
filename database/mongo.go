// database/mongo.go
package database

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
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

func createUri(username string, password string, clusterUrl string) string {

	// Build MongoDB connection string dynamically
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", username, password, clusterUrl)
	return mongoURI
}

func formaEnvVariables(usernameString string, passwordString string, clusterUrlString string) error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		errString := "Error loading .env file : " + err.Error()
		return fmt.Errorf(errString)
	}

	// Get MongoDB connection details from environment variables
	username = url.QueryEscape(os.Getenv(usernameString))
	password = url.QueryEscape(os.Getenv(passwordString))
	clusterURL = os.Getenv(clusterUrlString)
	return nil
}
