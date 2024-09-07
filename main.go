package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
}

var message requestBody

func PostHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": fmt.Sprintf("hello, %s", message.Message),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/write", PostHandler).Methods("POST")
	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
