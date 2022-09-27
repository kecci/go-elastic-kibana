package main

import (
	"go-elastic-kibana/cmd"
	"go-elastic-kibana/controller"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := cmd.LoadConfiguration("config/config.json")
	esConfig := cmd.LoadElasticsearcConfig(cfg)

	movieController := controller.NewMovieController(cfg, esConfig)

	// Init
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Router
	r.Get("/", controller.Hello)
	r.Get("/movie", movieController.Insert)

	// Start
	server := cfg.Server.Host + ":" + cfg.Server.Port
	println("listen on", "http://"+server)
	if err := http.ListenAndServe(server, r); err != nil {
		log.Fatal(err)
	}
}
