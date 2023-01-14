package main

import (
	"gochi/pkg/middleware"
	"gochi/pkg/routers"
	"gochi/utils"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	middleware.UseMiddlewares(r)
	routers.Public_routers(r)
	utils.StartServer(r)
}
