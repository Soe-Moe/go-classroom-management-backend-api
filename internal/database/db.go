package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB // Global DB variable

// InitDB connect to database
func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal("Failed to open a DB connection: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping DB: ", err)
    }

    log.Println("Database connection established")
}