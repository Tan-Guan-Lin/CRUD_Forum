package main

import (
	"log"
	"net/http"

	"github.com/Tan-Guan-Lin/CRUD_Forum/backend/routes"
)

func main() {

	r := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
