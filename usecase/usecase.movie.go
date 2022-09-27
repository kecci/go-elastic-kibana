package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-elastic-kibana/model"
	"log"
	"strconv"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	INDEX_MOVIES = "movies"
)

type (
	MovieUsecase interface {
		InsertBulk(datas []model.TmdbMovieData) error
		SearchByTitle(title string) ([]model.TmdbMovieData, error)
	}
	MovieUsecaseImpl struct {
		config   model.Config
		esClient *elasticsearch.Client
	}
)

// NewMovieUsecase movie
func NewMovieUsecase(config model.Config, esClient *elasticsearch.Client) MovieUsecase {
	return MovieUsecaseImpl{config: config, esClient: esClient}
}

// InsertMovie insert movie
func (c MovieUsecaseImpl) InsertBulk(datas []model.TmdbMovieData) error {
	// Setup waitgroup
	var wg sync.WaitGroup

	for _, data := range datas {

		wg.Add(1)

		go func(data model.TmdbMovieData) {
			defer wg.Done()
			// Build the request body.
			dataByte, err := json.Marshal(data)
			if err != nil {
				fmt.Printf("Error json.Marshal: %s \n", err)
				return
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      INDEX_MOVIES,
				DocumentID: strconv.Itoa(data.ID),
				Body:       bytes.NewReader(dataByte),
				Refresh:    "true",
			}

			// Perform the request with the client.
			resDo, err := req.Do(context.Background(), c.esClient)
			if err != nil {
				fmt.Printf("Error getting response: %s \n", err)
				return
			}
			defer resDo.Body.Close()
		}(data)
	}

	wg.Wait()

	return nil
}

// Search for the indexed documents
func (c MovieUsecaseImpl) SearchByTitle(title string) ([]model.TmdbMovieData, error) {
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": map[string]interface{}{
					"query":     title,
					"fuzziness": "AUTO",
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil, err
	}

	// Perform the search request.
	res, err := c.esClient.Search(
		c.esClient.Search.WithContext(context.Background()),
		c.esClient.Search.WithIndex(INDEX_MOVIES),
		c.esClient.Search.WithBody(&buf),
		c.esClient.Search.WithTrackTotalHits(true),
		c.esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
			return nil, err
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return nil, err
		}
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil, err
	}

	var datas []model.TmdbMovieData

	hitsParent := result["hits"].(map[string]interface{})
	hitsList := hitsParent["hits"].([]interface{})

	for _, hit := range hitsList {
		var data model.TmdbMovieData
		source := hit.(map[string]interface{})["_source"]
		sourceByte, _ := json.Marshal(source)
		if err := json.Unmarshal(sourceByte, &data); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}
