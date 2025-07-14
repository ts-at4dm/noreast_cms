package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Connection Established...")
}

func main() {
	http.HandleFunc("/", handler)
	
	port := ":8080"
	
	log.Printf("Connected using port %v", port)
	log.Fatal(http.ListenAndServe(port, nil))	
}