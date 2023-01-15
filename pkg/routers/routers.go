package routers

import (
	"database/sql"
	"gochi/app/controllers"
	"gochi/app/models"
	"gochi/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		r.With(middleware.RequestID).Get("/", controllers.GetAllUsers(env))
		r.With(middleware.RequestID).Get("/{id}", controllers.GetUserByID(env))
		r.With(middleware.RequestID).Get("/{email}", controllers.GetUserByEmail(env))
		r.With(middleware.RequestID).Post("/", controllers.CreateUser(env))
		r.With(middleware.RequestID).Put("/", controllers.UpdateUser(env))
	})
}
