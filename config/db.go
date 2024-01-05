package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// Open a database connection
	db, err := sql.Open("postgres", "postgres://postgres:admin@localhost:5432/toko_elektronik?sslmode=disable")
	if err != nil {
		log.Fatalf("error init db: %v\n", err)
	}
	if err != nil {
		log.Fatal(err)
	}

	return db
}
