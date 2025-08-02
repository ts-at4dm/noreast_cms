package handlers

import (
	"net/http"
	"path/filepath"
)
// Loads Home Page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Serve the index.html file from the templates directory
		path := filepath.Join("templates", "index.html")
		http.ServeFile(w, r, path)
		return
	}
	// Serve 404 for any other path
	http.NotFound(w, r)
}

// Handles links within the Homepage
func ClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/clients" {
		// Server the client.html file from the templates directory
		path := filepath.Join("templates", "clients.html")
		http.ServeFile(w, r, path)
	}
	// Server 404 for any other path
	http.NotFound(w, r)
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/events" {
		// Serve the events.html file from the templates directory
		path := filepath.Join("templates", "events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// Handles links within the clients page
func AddClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/add_client" {
		// Serve the add_client.html file from templates director
		path := filepath.Join("templates", "add_client.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

func ViewClientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/view_clients" {
		// Serve the add_client.html file from templates directory
		path := filepath.Join("templates", "view_clients.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}
 func SearchClientsHandler(w http.ResponseWriter, r*http.Request) {
		if r.URL.Path == "/search_clients" {
		// Serve the add_client.html file from templates directory
		path := filepath.Join("templates", "search_clients.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// Handles links within events page

// Create Events
func CreateEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/create_events" {
		path := filepath.Join("templates", "create_events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// View Events
func ViewEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/view_events" {
		path := filepath.Join("templates", "view_events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// Search Events
func SearchEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/search_events" {
		path := filepath.Join("templates", "search_events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}