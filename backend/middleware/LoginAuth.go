package middleware

import (
	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/db"
)

func LoginAuth(Username string, Password string) bool {
	database := db.InitDB()
	defer database.Close()

	Stored_password, err := db.Fetchlogin(Username, database)
	if err != nil {
		return false
	}

	return Password == Stored_password

}
