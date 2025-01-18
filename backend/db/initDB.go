package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./db/data/data.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	return database
}
