package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/middleware"
	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	if middleware.LoginAuth(userResponse) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Login failed"))

}
