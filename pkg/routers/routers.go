package routers

import (
	"database/sql"
	"gochi/app/controllers"
	"gochi/app/models"
	"gochi/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func Public_routers(r *chi.Mux) {
	db, err := sql.Open("postgres", utils.ConnectionBuilder("postgres"))
	if err != nil {
		log.Fatalf("Error establishing connection to DB: %v", err)
	}
	env := &models.Env{Db: db}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", controllers.GetAllUsers(env))
		r.Get("/{id}", controllers.GetUserByID(env))
		r.Get("/ByEmail/{email}", controllers.GetUserByEmail(env))
		r.Post("/", controllers.CreateUser(env))
		r.Put("/", controllers.UpdateUser(env))
	})
}
