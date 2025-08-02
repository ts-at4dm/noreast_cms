package server

import (
	"log"
	"net/http"
	"path/filepath"
)



func homhandler(w http.ResponseWriter, r *http.Request) {
	homepage := r.URL.Path
	if homepage == "/" {
		// Serve the index.html file from the templates directory
		filepath := filepath.Join("templates", "index.html")
		http.ServeFile(w, r, filepath)
		return
	}
	// Serve 404 for any other path
	http.NotFound(w, r)
}

func Start() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :8080")
	log.Println("Press Ctrl+C to stop the server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}