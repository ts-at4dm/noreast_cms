package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Update the connection string with the new user and password
    dsn := "noreast:Renegade2022!@tcp(127.0.0.1:3306)/noreast_cms?parseTime=true"


    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }
    defer db.Close()

    // Test the connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    fmt.Println("Successfully connected to the database!")
}