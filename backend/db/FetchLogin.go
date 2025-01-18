package db

import (
	"database/sql"
	"errors"
	
)

func Fetchlogin(username string, db *sql.DB) (string, error) {

	query := "SELECT Password FROM Users WHERE Username = ?"
	var ScannedPassword string
	row := db.QueryRow(query, username)
	if err := row.Scan(&ScannedPassword); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", errors.New("Database error: " + err.Error())

	}
	return ScannedPassword, nil

}
