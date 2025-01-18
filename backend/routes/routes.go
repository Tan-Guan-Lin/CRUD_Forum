package routes

import (
	"net/http"

	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the homepage"))
	})
	r.Post("/login", handlers.LoginHandler)
	r.Post("/register", handlers.RegisterHandler)
	return r

}
