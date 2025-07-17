package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	// Load .env file. Will error out if not found
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error: .env file not found. Check source.")
	}

	// GET server IP from .env
	serverIP := os.Getenv("SERVER_IP")
	if serverIP == "" {
		log.Fatal("Server IP not found.")
	}
	
	// URL Connection
	url := fmt.Sprintf("http://%s:8080", serverIP)

    // Make GET request
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Failed to connect to server: %v", err)
    }
    defer resp.Body.Close()

    // Read response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response: %v", err)
    }

    fmt.Printf("Server response:\n%s\n", string(body))
}