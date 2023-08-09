package main

import (
	"AddressListener/handlers" // Update this to the correct package path
	"fmt"
	"net/http"
)

func main() {
	// Define your webhook endpoint and handler
	http.HandleFunc("/webhook-receiver", handlers.WebhookHandler)

	// Start the server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
