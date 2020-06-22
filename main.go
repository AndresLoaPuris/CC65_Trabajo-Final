package main

import (
	"Tutorial/api"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.IndexRouter)
	router.HandleFunc("/people", api.GetPeople).Methods("GET")
	router.HandleFunc("/people", api.CreatePeople).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(router)))

}
