package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"noreast_cms/internal/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func initDatabase() *sql.DB {
	// Get database connection string from environment or use default
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// Default connection string - adjust these values for your MySQL setup
		dsn = "noreast:Renegade2022!@tcp(127.0.0.1:3306)/noreast_cms?parseTime=true"
		log.Printf("Using default DSN. Set DB_DSN environment variable to customize.")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		log.Printf("Please ensure MySQL is running and credentials are correct")
		return nil
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("Error connecting to database: %v", err)
		log.Printf("Please check your MySQL server and database configuration")
		return nil
	}

	log.Println("Database connected successfully")
	return db
}

func setupHandlers() {
	// Client routes
	http.HandleFunc("/clients", handlers.ClientHandler)              // GET: Client menu page
	http.HandleFunc("/add_client", handlers.AddClientHandler)        // GET: Add client form
	http.HandleFunc("/clients/create", handlers.CreateClientHandler) // POST: Process client creation
	http.HandleFunc("/view_clients", handlers.ViewClientsHandler)    // GET: View all clients
	http.HandleFunc("/search_clients", handlers.SearchClientsHandler) // GET: Search clients form

	// Event routes
	http.HandleFunc("/events", handlers.EventHandler)                // GET: Events menu page
	http.HandleFunc("/create_events", handlers.CreateEventsHandler)  // GET: Create event form
	http.HandleFunc("/events/create", handlers.CreateEventHandler)   // POST: Process event creation
	http.HandleFunc("/view_events", handlers.ViewEventsHandler)      // GET: View all events
	http.HandleFunc("/search_events", handlers.SearchEventsHandler)  // GET: Search events form

	// Home route
	http.HandleFunc("/", handlers.HomeHandler) // GET: Home page

	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func ServerStart() {
	// Initialize database connection
	db := initDatabase()
	if db != nil {
		// Set the database connection for handlers
		handlers.SetDatabase(db)
		defer db.Close()
	} else {
		log.Println("Warning: Running without database connection")
	}

	// Setup route handlers
	setupHandlers()

	log.Println("Server starting on http://localhost:8080")
	log.Println("Available routes:")
	log.Println("  GET  / - Home page")
	log.Println("  GET  /clients - Client management")
	log.Println("  GET  /add_client - Add client form")
	log.Println("  POST /clients/create - Create client")
	log.Println("  GET  /view_clients - View all clients")
	log.Println("  GET  /events - Event management")
	log.Println("  GET  /create_events - Create event form")
	log.Println("  POST /events/create - Create event")
	log.Println("  GET  /view_events - View all events")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Server failed to start: %v\n", err)
		return
	}
}