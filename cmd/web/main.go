package main

import (
	"net/http"
	"flag"
	"log"
)

func main() {
	mux := http.NewServeMux()

	addr := flag.String("addr", ":4000", "Net Address for app")

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /user/view/{id}", userView)
	mux.HandleFunc("GET /user/create", userCreate)
	mux.HandleFunc("POST /user/create", userCreatePost)

	log.Printf("starting server on port: %v", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}