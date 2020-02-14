[![GoDoc](https://godoc.org/github.com/elastic/go-elasticsearch?status.svg)](http://godoc.org/github.com/elastic/go-elasticsearch)

# Golang Elasticsearch Kibana
This is simple project for golang, elasticsearch, and kibana

## Docker
### Docker & docker-compose Installation
On Mac:
[`Installation Docker with Mac`](https://docs.docker.com/docker-for-mac/install/)

Try to check docker version:
```bash
$ docker --version
$ docker-compose version
```

## Elasticsearch
### Elasticsearch Installation on Docker
Download image of elasticsearch with version `7.6.0`:
```bash
$ docker pull elasticsearch:7.6.0
```

Try to run image of elasticsearch to be a container:
```bash
$ docker run -d --network:elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --name elasticsearch elasticsearch:7.6.0
$ docker ps
```

## Kibana
### Kibana Installation on Docker
Download image of kibana with version `7.6.0`:
```bash
$ docker pull kibana:7.6.0
or
$ docker pull docker.elastic.co/kibana/kibana:7.6.0
```

Try run image of kibana to be a container:
```bash
$ docker run -d --net:elastic -p 5601:5601 -e ELASTICSEARCH_URL=http://localhost:9200 --name kibana kibana:7.6.0
$ docker ps
```

## Docker Compose
### Create docker compose for Elasticsearch and Kibana
Create file with name `docker-compose.yml` and type this example configuration into file:
```
version: '2.2'
services:
  es01:
    image: elasticsearch:7.6.0
    container_name: es01
    environment:
      - node.name=es01
      - cluster.name=es-docker-cluster
      - discovery.seed_hosts=es02
      - cluster.initial_master_nodes=es01,es02
      - bootstrap.memory_lock=true
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - data01:/usr/share/elasticsearch/data
    ports:
      - 9200:9200
    networks:
      - elastic
  es02:
    image: elasticsearch:7.6.0
    container_name: es02
    environment:
      - node.name=es02
      - cluster.name=es-docker-cluster
      - discovery.seed_hosts=es01
      - cluster.initial_master_nodes=es01,es02
      - bootstrap.memory_lock=true
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - data02:/usr/share/elasticsearch/data
    networks:
      - elastic

  kibana:
    image: kibana:7.6.0
    container_name: kibana
    environment:
      SERVER_NAME: kibana.local
      ELASTICSEARCH_HOSTS: http://es01:9200
    ports: 
       - 5601:5601
    networks:
       - elastic
    depends_on:
       - es01
       - es02

volumes:
  data01:
    driver: local
  data02:
    driver: local

networks:
  elastic:
    driver: bridge
```

In this case, I'm creating 3 containers:
* es01 (elasticsearch for node 1)
* es02 (elasticsearch for node 2)
* kibana (as dashboard)

### Run Docker Compose
To run docker compose:
```bash
$ docker-compose up -d
$ docker ps
```

Make sure all containers running

## golang
### Golang Installation
Install golang [`Install golang`](https://golang.org/doc/install)

Try to check version
```bash
$ go version
```

### Setup Project
This is golang library to integrate with elasticsearch

First, download library for go-elasticsearch, in here we are using version v7:
```bash
$ go mod init
$ go get github.com/elastic/go-elasticsearch/v7
```

Then create file `main.go` and insert this example code:
```go
package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

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

	res, err := es.Info()

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)

	log.Println(es.Indices)
}
```

### Run golang elasticsearch Application
Try to run application:
```bash
$ go run main.go
```

See More in here [`Documentation of go-elasticsearch`](https://github.com/elastic/go-elasticsearch/)