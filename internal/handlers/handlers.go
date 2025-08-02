package handlers

import (
	"net/http"
	"path/filepath"
)

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

func AddClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/add_client" {
		// Serve the add_client.html file from templates director
		path := filepath.Join("templates", "add_client.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}
