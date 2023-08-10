package database

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

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
