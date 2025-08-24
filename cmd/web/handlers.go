package main

import(
	"net/http"
	"strconv"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("home page"))
}

func userView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	
	fmt.Fprintf(w, "user with id: %v", id)
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user Create"))
}

func userCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user Create Post"))
}