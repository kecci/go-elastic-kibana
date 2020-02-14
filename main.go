package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Log struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "foo",
		Password: "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}

	es, _ := elasticsearch.NewClient(cfg)

	newLog := Log{
		Name:   "TEST",
		Status: "OK",
	}

	insertLog(es, newLog)

}

func insertLog(es *elasticsearch.Client, newLog Log) {
	dataJSON, err := json.Marshal(newLog)
	js := string(dataJSON)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "logs",
		Body:    strings.NewReader(js),
		Refresh: "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}
