package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gochi/app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllUsers(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := env.Db.Query("SELECT id,name,email,password,created_at,updated_at FROM public.users")
		if err != nil {
			w.Write([]byte("Error while getting users"))
			return
		}
		defer rows.Close()
		var users []models.User
		for rows.Next() {
			var user models.User
			err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
			if err != nil {
				w.Write([]byte("Error getting users " + err.Error()))
				return
			}
			users = append(users, user)
		}
		if err = rows.Err(); err != nil {
			w.Write([]byte("Error on rows"))
			return
		}
		j, err := json.Marshal(users)
		if err != nil {
			w.Write([]byte("Error encoding users to JSON"))
			return
		}
		w.Write([]byte(j))
	}
}

func GetUserByID(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		log.Println(id)
		if err != nil {
			w.Write([]byte("ID is not a number"))
			return
		}
		row := env.Db.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM public.users WHERE id=$1", id)
		var user models.User
		switch err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at); err {
		case sql.ErrNoRows:
			w.Write([]byte("No rows were returned"))
		case nil:
			j, err := json.Marshal(user)
			if err != nil {
				w.Write([]byte("error turning to json"))
				return
			}
			w.Write([]byte(j))
		default:
			w.Write([]byte(fmt.Sprintf("Error while mapping user: %v", err)))
		}
	}
}

func CreateUser(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.CreatedUser
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(user)
		var id int
		err = env.Db.QueryRow(`INSERT INTO public.users (name, email, password) VALUES($1,$2,$3) RETURNING id`, user.Name, user.Email, user.Password).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("{\"ID\":%d, \"Success\": %t}", id, true)))
	}
}
