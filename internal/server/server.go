package server

import (
	"log"
	"net/http"
	"path/filepath"
)



func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Serve the index.html file from the templates directory
		path := filepath.Join("templates", "index.html")
		http.ServeFile(w, r, path)
		return
	}
	// Serve 404 for any other path
	http.NotFound(w, r)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/clients" {
		// Server the client.html file from the templates directory
		path := filepath.Join("templates", "clients.html")
		http.ServeFile(w, r, path)
	}
	// Server 404 for any other path
	http.NotFound(w, r)
}

func Start() {
	http.HandleFunc("/clients", clientHandler)
	http.HandleFunc("/", homeHandler)
	log.Println("Server started on :8080")
	log.Println("Press Ctrl+C to stop the server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}