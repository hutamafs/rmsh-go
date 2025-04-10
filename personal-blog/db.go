package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
    const (
        host     = "localhost"
        port     = 5432
        user     = "postgres"
        password = ""
        dbname   = "blog"
    )

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Open DB
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }

    // Test connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    log.Println("✅ Connected to the database!")
    return db
}