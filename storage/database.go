package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Storage]:", "Error loading .env file")
		return nil, err
	}

	// Read environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// Construct connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@localhost/%s?sslmode=disable", dbUser, dbPassword, dbName)

	// Connect to the database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("[Storage]:", "couldn't connect with database")
		return nil, err
	}
	log.Println("[Storage]:", "connected to database")
	return DB, nil
}
