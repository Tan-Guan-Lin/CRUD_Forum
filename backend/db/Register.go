package db

import(
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"log"
)

func Register(Username string, HashedPassword []byte, database *sql.DB){
	query := "INSERT INTO Users (Username, Password) VALUES (?, ?);"
	_, err := database.Exec(query, Username, HashedPassword)
	if err != nil {
		log.Fatal(err)
	}

}