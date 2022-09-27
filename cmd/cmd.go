package cmd

import (
	"go-elastic-kibana/controller"
	"go-elastic-kibana/usecase"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	configName = "config/config.json"
)

func Run() {
	// Configs
	cfg := LoadConfiguration(configName)
	esClient := LoadElasticsearchClient(cfg)

	// Usecases
	movieUsecase := usecase.NewMovieUsecase(cfg, esClient)

	// Controllers
	movieController := controller.NewMovieController(cfg, movieUsecase)

	// Server
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Router
	r.Get("/", controller.Hello)
	r.Put("/movie", movieController.Insert)
	r.Get("/movie", movieController.SearchByTitle)

	// Start
	server := cfg.Server.Host + ":" + cfg.Server.Port
	println("listen on", "http://"+server)
	if err := http.ListenAndServe(server, r); err != nil {
		log.Fatal(err)
	}
}
