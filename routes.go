package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func DeclareRoutes() {

	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/arts", getArts).Methods("GET")
	r.HandleFunc("/arts/{id}", getArt).Methods(("GET"))
	r.HandleFunc("/arts", createArt).Methods("POST")
	r.HandleFunc("/arts/{id}", deleteArt).Methods("DELETE")

	fmt.Printf("Server Started on Port%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err.Error())
	}
}
