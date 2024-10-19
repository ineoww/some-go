package main

import (
	"log"
	"net/http"
)


func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	


	log.Println("Starting server on port 9000")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}