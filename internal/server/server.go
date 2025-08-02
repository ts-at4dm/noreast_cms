package server

import (
	"log"
	"net/http"

	"noreast_cms/internal/handlers"
)

func setupHandlers() {
	// Handles Events page
	http.HandleFunc("/create_events", handlers.CreateEventsHandler)
	http.HandleFunc("/view_events", handlers.ViewEventsHandler)
	http.HandleFunc("/search_events", handlers.SearchEventsHandler)

	// Handles Clients page
	http.HandleFunc("/search_clients", handlers.SearchClientsHandler)
	http.HandleFunc("/view_clients", handlers.ViewClientsHandler)
	http.HandleFunc("/add_client", handlers.AddClientHandler)
	
	// Handles Home page
	http.HandleFunc("/events", handlers.EventHandler)
	http.HandleFunc("/clients", handlers.ClientHandler)
	http.HandleFunc("/", handlers.HomeHandler)
}

func ServerStart() {
	// Calls Handlers
	setupHandlers()

	log.Println("Server running on Port 8080...")

	// Start the server and handle any errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Server failed to start: %v\n", err)
		return
	}
}