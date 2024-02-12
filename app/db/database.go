package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
