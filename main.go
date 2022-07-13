package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()

	initialize()

	// middleware
	middle(r)

	//routes
	routes(r)

	http.ListenAndServe(":3000", r)
}
