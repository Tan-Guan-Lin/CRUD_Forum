package middleware

import(
	"log"
	"errors"

	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/db"
	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/models"
	"golang.org/x/crypto/bcrypt"

)

func RegisterAuth(User models.UserResponse) error {
	database := db.InitDB()
	defer database.Close()

	var username string = User.Username

	if str, err := db.Fetchlogin(username, database); err != nil && str == ""{
		if err.Error() == "user not found"{			
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err)
			}
			db.Register(username, hashedPassword, database)
			return nil
		}
		return err
	}

	err2 := errors.New("user already exists")
	return err2

	


}