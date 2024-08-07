package main

import (
	"log"
	"net/http"
)

var renderer, _ = NewPostRenderer()

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Create a sample post
	post := Post{
		Title:       "Sample Post",
		Description: "This is a sample description",
		Body:        "This is the body of the post",
		Tags:        []string{"go", "programming", "web"},
	}

	// Render the post
	if err := renderer.Render(w, post); err != nil {
		http.Error(w, "Failed to render post", http.StatusInternalServerError)
	}
}

func main() {

	// Define a handler function
	http.HandleFunc("/post", handlePost)

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
