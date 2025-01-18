package handlers

import (
	"net/http"
	"encoding/json"


	
	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/models"
	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/middleware"

)

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	// Only handle POST requests for login
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}


	var userResponse models.UserResponse
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userResponse); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	
	err2 := middleware.RegisterAuth(userResponse)
	if err2 == nil{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Registration successful"))
		return
	}
	if err2.Error() == "user already exists"{
		http.Error(w, "Username already exists. Please choose a different one.", http.StatusConflict)
		return
	}
	http.Error(w, "Internal server error", http.StatusInternalServerError)



}