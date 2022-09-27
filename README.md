[![GoDoc](https://godoc.org/github.com/elastic/go-elasticsearch?status.svg)](http://godoc.org/github.com/elastic/go-elasticsearch)

# Golang Elasticsearch Kibana
This is simple project for golang, elasticsearch, and kibana

## Elastic Installation
Run this docker: https://github.com/kecci/docker-kit/tree/master/elasticsearch

## Golang
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

### Run golang elasticsearch Application
Try to run application:
```bash
$ go run ./main.go
```

See More in here [`Documentation of go-elasticsearch`](https://github.com/elastic/go-elasticsearch/)

## Data Source
Credentials: https://www.themoviedb.org/settings/api
API Docs: https://developers.themoviedb.org/3
Postman: https://www.postman.com/devrel/workspace/tmdb-api/overview
