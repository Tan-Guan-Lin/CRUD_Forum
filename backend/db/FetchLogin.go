package db

import (
	"database/sql"
	"errors"
	
)

func Fetchlogin(username string, db *sql.DB) ([]byte, error) {

	query := "SELECT Password FROM Users WHERE Username = ?"
	var ScannedPassword []byte
	row := db.QueryRow(query, username)
	if err := row.Scan(&ScannedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("Database error: " + err.Error())

	}
	return ScannedPassword, nil

}
