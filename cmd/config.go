package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-elastic-kibana/model"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

// LoadConfiguration load conf
func LoadConfiguration(file string) (config model.Config) {
	// Openfile
	configFile, err := os.Open(file)
	defer func() {
		if err := configFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Parse
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func LoadElasticsearchClient(cfg model.Config) *elasticsearch.Client {
	esConfig := elasticsearch.Config{
		Addresses: cfg.Elasticsearch.Addresses,
		Username:  cfg.Elasticsearch.Username,
		Password:  cfg.Elasticsearch.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   cfg.Elasticsearch.MaxIdleConnsPerHost,
			ResponseHeaderTimeout: time.Duration(cfg.Elasticsearch.ResponseHeaderTimeout) * time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Duration(cfg.Elasticsearch.DialContextTimeout) * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}

	// Init clieant
	es, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		fmt.Printf("Error elasticsearch.NewClient: %s \n", err)
		return nil
	}

	// Get cluster info
	res, err := es.Info()
	if err != nil {
		fmt.Printf("Error es.Info(): %s \n", err)
		return nil
	}
	defer res.Body.Close()

	return es
}
