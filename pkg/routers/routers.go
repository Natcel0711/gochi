package routers

import (
	"encoding/json"
	"net/http"

	"gochi/app/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Public_routers(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/users", func(r chi.Router) {
		r.With(middleware.RequestID).Get("/", func(w http.ResponseWriter, r *http.Request) {
			j, err := json.Marshal(models.Employee{
				Name: "Natcel",
				Age:  24,
			})
			if err != nil {
				w.Write([]byte("Error while encoding employee"))
			}
			w.Write([]byte(j))
		})
	})
}
