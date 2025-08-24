package main

import(
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("home page"))
}

func userView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user view"))
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user Create"))
}

func userCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user Create Post"))
}