package main

// Item represents an item to be stored or retrieved.
type Item struct {
	ID   int    `json:"id"`   // ID is a unique identifier for each item
	Name string `json:"name"` // Name is the name of the item
}
