package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the routes and their handler functions
	http.HandleFunc("/create", Create)  // POST to create an item
	http.HandleFunc("/get", Get)                // GET to retrieve an item by ID

	// Start the server on port 8080
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
