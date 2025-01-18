package middleware

import (

	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/db"
	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginAuth(User models.UserResponse) bool {
	database := db.InitDB()
	defer database.Close()

	Stored_password, err := db.Fetchlogin(User.Username, database)
	if err != nil {
		return false
	}

	if bcrypt.CompareHashAndPassword(Stored_password, []byte(User.Password)) == nil{
		return true
	}
	return false

}
