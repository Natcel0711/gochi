package utils

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartServer(r *chi.Mux) {
	log.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
