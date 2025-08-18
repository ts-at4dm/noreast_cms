package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Database instance
var db *sql.DB

// SetDatabase allows setting the database connection for handlers
func SetDatabase(database *sql.DB) {
	db = database
}

// Home Page Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		path := filepath.Join("templates", "index.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// ===== CLIENT HANDLERS =====

// ClientHandler - Shows client management menu
func ClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/clients" {
		path := filepath.Join("templates", "clients.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// AddClientHandler - Shows add client form
func AddClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/add_client" {
		path := filepath.Join("templates", "add_client.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// CreateClientHandler - Processes client creation
func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Extract form values
	client := struct {
		FirstName   string
		LastName    string
		CompanyName string
		Email       string
		Phone       string
		Address     string
		City        string
		State       string
		Zip         string
	}{
		FirstName:   strings.TrimSpace(r.FormValue("firstname")),
		LastName:    strings.TrimSpace(r.FormValue("lastname")),
		CompanyName: strings.TrimSpace(r.FormValue("company_name")),
		Email:       strings.TrimSpace(r.FormValue("email")),
		Phone:       strings.TrimSpace(r.FormValue("phone")),
		Address:     strings.TrimSpace(r.FormValue("st_address")),
		City:        strings.TrimSpace(r.FormValue("city")),
		State:       strings.TrimSpace(r.FormValue("state")),
		Zip:         strings.TrimSpace(r.FormValue("zip")),
	}

	// Basic validation
	if client.FirstName == "" || client.LastName == "" || client.Email == "" {
		http.Error(w, "First name, last name, and email are required", http.StatusBadRequest)
		return
	}

	// Check database connection
	if db == nil {
		log.Printf("Database connection not initialized")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	// Insert into database
	query := `INSERT INTO clients (firstname, lastname, company_name, email, phone, st_address, city, state, zip) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query,
		client.FirstName, client.LastName, client.CompanyName,
		client.Email, client.Phone, client.Address,
		client.City, client.State, client.Zip)

	if err != nil {
		log.Printf("Error inserting client: %v", err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			http.Error(w, "A client with this email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Error saving client to database", http.StatusInternalServerError)
		return
	}

	clientID, _ := result.LastInsertId()
	log.Printf("Successfully created client with ID: %d", clientID)

	// Redirect to view clients page
	http.Redirect(w, r, "/view_clients", http.StatusSeeOther)
}

// ViewClientsHandler - Shows all clients
// ViewClientsHandler - Shows all clients
func ViewClientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/view_clients" {
		http.NotFound(w, r)
		return
	}

	if db == nil {
		log.Printf("Database connection not initialized")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	// Query all clients
	rows, err := db.Query(`
        SELECT id, firstname, lastname, company_name, email, phone, st_address, city, state, zip, date_added 
        FROM clients ORDER BY date_added DESC`)
	if err != nil {
		log.Printf("Error querying clients: %v", err)
		http.Error(w, "Error retrieving clients", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Struct to hold client data
	type Client struct {
		ID          int
		FirstName   string
		LastName    string
		CompanyName string
		Email       string
		Phone       string
		Address     string
		City        string
		State       string
		Zip         string
		DateAdded   time.Time
	}

	var clients []Client

	// Read rows into the slice
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.CompanyName, &c.Email,
			&c.Phone, &c.Address, &c.City, &c.State, &c.Zip, &c.DateAdded); err != nil {
			log.Printf("Error scanning client: %v", err)
			http.Error(w, "Error reading client data", http.StatusInternalServerError)
			return
		}
		clients = append(clients, c)
	}

	// Check for any row iteration errors
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		http.Error(w, "Error reading client list", http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmplPath := filepath.Join("templates", "view_clients.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, clients); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}
}

// SearchClientsHandler - Shows client search form
func SearchClientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/search_clients" {
		path := filepath.Join("templates", "search_clients.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// ===== EVENT HANDLERS =====

// EventHandler - Shows event management menu
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/events" {
		path := filepath.Join("templates", "events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// CreateEventsHandler - Shows create event form
func CreateEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/create_events" {
		path := filepath.Join("templates", "create_events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// CreateEventHandler - Processes event creation
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Extract form values
	eventName := strings.TrimSpace(r.FormValue("event_name"))
	eventType := strings.TrimSpace(r.FormValue("event_type"))
	eventLocation := strings.TrimSpace(r.FormValue("event_location"))
	ceremonyLocation := strings.TrimSpace(r.FormValue("ceremony_location"))
	packageType := strings.TrimSpace(r.FormValue("package"))
	notes := strings.TrimSpace(r.FormValue("notes"))

	// Parse dates and numbers
	eventDateStr := r.FormValue("event_date")
	startTimeStr := r.FormValue("start_time")
	endTimeStr := r.FormValue("end_time")
	paymentDateStr := r.FormValue("payment_date")

	clientIDStr := r.FormValue("client_id")
	guestCountStr := r.FormValue("guest_count")
	depositAmountStr := r.FormValue("deposit_amount")
	totalPriceStr := r.FormValue("total_price")

	depositReceived := r.FormValue("deposit_received") == "1"
	paymentReceived := r.FormValue("payment_received") == "1"

	// Validation
	if eventName == "" || eventDateStr == "" || clientIDStr == "" || depositAmountStr == "" || totalPriceStr == "" {
		http.Error(w, "Event name, date, client ID, deposit amount, and total price are required", http.StatusBadRequest)
		return
	}

	// Parse numeric values
	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	var guestCount int
	if guestCountStr != "" {
		guestCount, err = strconv.Atoi(guestCountStr)
		if err != nil {
			http.Error(w, "Invalid guest count", http.StatusBadRequest)
			return
		}
	}

	depositAmount, err := strconv.ParseFloat(depositAmountStr, 64)
	if err != nil {
		http.Error(w, "Invalid deposit amount", http.StatusBadRequest)
		return
	}

	totalPrice, err := strconv.ParseFloat(totalPriceStr, 64)
	if err != nil {
		http.Error(w, "Invalid total price", http.StatusBadRequest)
		return
	}

	// Parse date
	eventDate, err := time.Parse("2006-01-02", eventDateStr)
	if err != nil {
		http.Error(w, "Invalid event date", http.StatusBadRequest)
		return
	}

	// Parse payment date (nullable)
	var paymentDate *time.Time
	if paymentDateStr != "" {
		pd, err := time.Parse("2006-01-02", paymentDateStr)
		if err != nil {
			http.Error(w, "Invalid payment date", http.StatusBadRequest)
			return
		}
		paymentDate = &pd
	}

	// Check database connection
	if db == nil {
		log.Printf("Database connection not initialized")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	// Insert into database
	query := `INSERT INTO events (event_date, event_name, event_type, start_time, end_time, client_id, 
              event_location, ceremony_location, package, guest_count, deposit_amount, deposit_received, 
              total_price, payment_received, payment_date, notes) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query,
		eventDate, eventName, eventType, startTimeStr, endTimeStr, clientID,
		eventLocation, ceremonyLocation, packageType, guestCount, depositAmount, depositReceived,
		totalPrice, paymentReceived, paymentDate, notes)

	if err != nil {
		log.Printf("Error inserting event: %v", err)
		http.Error(w, "Error saving event to database", http.StatusInternalServerError)
		return
	}

	eventID, _ := result.LastInsertId()
	log.Printf("Successfully created event with ID: %d", eventID)

	// Redirect to view events page
	http.Redirect(w, r, "/view_events", http.StatusSeeOther)
}

// ViewEventsHandler - Shows all events
func ViewEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/view_events" {
		if db == nil {
			log.Printf("Database connection not initialized")
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}

		// For now, just serve the template
		path := filepath.Join("templates", "view_events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

// SearchEventsHandler - Shows event search form
func SearchEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/search_events" {
		path := filepath.Join("templates", "search_events.html")
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}
