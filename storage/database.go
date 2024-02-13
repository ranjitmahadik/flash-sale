package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("postgres", "postgresql://user:user@123@localhost/flash-sale?sslmode=disable")
	if err != nil {
		log.Fatal("[Storage]:", "couldn't connect with database")
		return nil, err
	}
	log.Println("[Storage]:", "connected to database")
	return DB, nil
}
