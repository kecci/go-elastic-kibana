package controller

import "net/http"

// Hello world
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
