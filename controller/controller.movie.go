package controller

import (
	"go-elastic-kibana/model"
	"go-elastic-kibana/usecase"
	"go-elastic-kibana/utility"
	"net/http"
)

type (
	MovieController interface {
		Insert(w http.ResponseWriter, r *http.Request)
		SearchByTitle(w http.ResponseWriter, r *http.Request)
	}

	MovieControllerImpl struct {
		config       model.Config
		movieUsecase usecase.MovieUsecase
	}
)

// NewMovieController movie controller
func NewMovieController(
	config model.Config,
	movieUsecase usecase.MovieUsecase,
) MovieController {
	return MovieControllerImpl{
		config:       config,
		movieUsecase: movieUsecase,
	}
}

// Insert movie
func (c MovieControllerImpl) Insert(w http.ResponseWriter, r *http.Request) {

	// Mapping request body
	var movieModels []model.TmdbMovieData
	if err := utility.RequestBodyToStruct(w, r.Body, &movieModels); err != nil {
		model.MapBaseResponse(w, r, "ERROR", nil, err)
		return
	}

	// Call usecase
	if err := c.movieUsecase.InsertBulk(movieModels); err != nil {
		model.MapBaseResponse(w, r, "ERROR", nil, err)
		return
	}

	model.MapBaseResponse(w, r, "SUCCESS", true, nil)
}

// SearchByTitle title
func (c MovieControllerImpl) SearchByTitle(w http.ResponseWriter, r *http.Request) {

	// Mapping request query
	title := r.URL.Query().Get("title")

	// Call usecase
	res, err := c.movieUsecase.SearchByTitle(title)
	if err != nil {
		model.MapBaseResponse(w, r, "ERROR", nil, err)
		return
	}

	model.MapBaseResponse(w, r, "SUCCESS", res, nil)
}
