// internal/database/connection.go
package database

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

type DB struct {
    *sql.DB
}

func NewConnection(dsn string) (*DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    log.Println("Database connected successfully")
    return &DB{db}, nil
}