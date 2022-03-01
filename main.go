package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book struct (Model)
type Book struct {
	ID string `json: "id`
}

func main() {

	//Init router
	r := mux.NewRouter()

	//Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe("8000", r))
}
