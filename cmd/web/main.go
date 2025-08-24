package main

import (
	"net/http"
	"flag"
)

func main() {
	mux := http.NewServeMux()

	mux.HandlerFunc("GET /{$}", home)
	mux.HandlerFunc("GET /user/view/{id}", userView)
	mux.HandleFunc("GET /user/create", userCreate)
	mux.HandleFunc("POST /user/create", userCreatePost)

	log.Printf("startig server on port: %v", addr)
}