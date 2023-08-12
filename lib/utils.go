package lib

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var Username string
var Password string
var ClusterURL string

var DBName string
var CollectionName string

var WebhookUrl string

func FormatEnvVariables(usernameString string, passwordString string, clusterUrlString string, dbNameString string, collectionNameString string, webhookUrl string) error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		errString := "Error loading .env file : " + err.Error()
		return fmt.Errorf(errString)
	}

	// Get MongoDB connection details from environment variables
	Username = url.QueryEscape(os.Getenv(usernameString))
	Password = url.QueryEscape(os.Getenv(passwordString))
	ClusterURL = os.Getenv(clusterUrlString)
	DBName = os.Getenv(dbNameString)
	CollectionName = os.Getenv(collectionNameString)
	WebhookUrl = os.Getenv(webhookUrl)

	return nil
}

func CreateUri(username string, password string, clusterUrl string) string {

	// Build MongoDB connection string dynamically
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", username, password, clusterUrl)
	return mongoURI
}
