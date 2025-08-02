package server

import (
	"log"
	"net/http"

	"noreast_cms/internal/handlers"
)

func setupHandlers() {
	// Handles Events
	http.HandleFunc("/create_events", handlers.CreateEventsHandler)
	http.HandleFunc("/view_events", handlers.ViewEventsHandler)
	http.HandleFunc("/search_events", handlers.SearchEventsHandler)

	// Handles Clients 
	http.HandleFunc("/search_clients", handlers.SearchClientsHandler)
	http.HandleFunc("/view_clients", handlers.ViewClientsHandler)
	http.HandleFunc("/add_client", handlers.AddClientHandler)
	
	// Handles Home
	http.HandleFunc("/events", handlers.EventHandler)
	http.HandleFunc("/clients", handlers.ClientHandler)
	http.HandleFunc("/", handlers.HomeHandler)
}

func ServerStart() {

	setupHandlers()

	log.Println("Server started on :8080")
	log.Println("Press Ctrl+C to stop the server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}