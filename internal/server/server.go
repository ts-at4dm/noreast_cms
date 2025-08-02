package server

import (
	"log"
	"net/http"

	"noreast_cms/internal/handlers"
)


func Start() {
	http.HandleFunc("/events", handlers.EventHandler)
	http.HandleFunc("/clients", handlers.ClientHandler)
	http.HandleFunc("/", handlers.HomeHandler)
	log.Println("Server started on :8080")
	log.Println("Press Ctrl+C to stop the server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}