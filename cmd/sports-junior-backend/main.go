package main

import (
	"fmt"
	"net/http"
	"sports-junior-backend/internal/handlers"
)

func main() {
	http.HandleFunc("/notes", handlers.HandleNotes)
	http.HandleFunc("/notes/", handlers.DeleteNote)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
