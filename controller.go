package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getArts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ConnectToDB()
	var Arts []Art
	result := DB.Find(&Arts)

	if result.Error != nil {
		panic(result.Error)
	}

	json.NewEncoder(w).Encode(Arts)
}

func getArt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ConnectToDB()

	params := mux.Vars(r)
	var art Art
	DB.Where("id = ?", params["id"]).Find(&art)

	if (Art{} == art) {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(art)

}

func createArt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ConnectToDB()

	var art Art
	json.NewDecoder(r.Body).Decode(&art)

	DB.Create(&art)

	json.NewEncoder(w).Encode(art)

}

func deleteArt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ConnectToDB()

	params := mux.Vars(r)

	var art Art
	DB.Where("id = ?", params["id"]).Find(&art)

	if (Art{} == art) {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	DB.Delete(&Art{}, params["id"])

	fmt.Fprint(w, http.StatusOK)
}
