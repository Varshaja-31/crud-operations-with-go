package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var items = make(map[int]Item)
var idCounter = 1

// Function to handle POST requests and create an item
func Create(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var item Item
	// Decode the request body into an Item struct
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

// Assign an ID to the new item
	item.ID = idCounter
	idCounter++

	// Store the item in the items map
	items[item.ID] = item

	w.Header().Set("Content-Type", "application/json")

	// Respond with the created item as JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Function to handle GET requests and retrieve an item by ID.
func Get(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id") // Get the ID from query parameter
	if idStr == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(idStr) // Convert ID from string to int
	if err != nil {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
		return
	}

	// Look up the item by ID
	if item, exists := items[id]; exists {
		json.NewEncoder(w).Encode(item)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}
