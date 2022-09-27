package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-elastic-kibana/model"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type (
	MovieController interface {
		Insert(w http.ResponseWriter, r *http.Request)
	}

	MovieControllerImpl struct {
		config   model.Config
		esConfig elasticsearch.Config
	}
)

func NewMovieController(config model.Config, esConfig elasticsearch.Config) MovieController {
	return MovieControllerImpl{config: config, esConfig: esConfig}
}

func (c MovieControllerImpl) Insert(w http.ResponseWriter, r *http.Request) {
	movieModel := model.MovieModel{
		Title:  "TEST",
		Rating: 10,
	}

	if err := c.insertMovie(movieModel); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success"))
	return
}

func (c MovieControllerImpl) insertMovie(model model.MovieModel) error {
	es, _ := elasticsearch.NewClient(c.esConfig)
	modelByte, err := json.Marshal(model)
	if err != nil {
		return err
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "logs",
		Body:    bytes.NewReader(modelByte),
		Refresh: "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
	return nil
}
