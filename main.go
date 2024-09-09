package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var message RequestBody

func PostHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	message.ID = 0

	DB.Create(&message)

	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var response []RequestBody
	DB.Find(&response)
	msgSlice := make([]string, len(response))
	for index, value := range response {
		msgSlice[index] = value.Message
	}
	json.NewEncoder(w).Encode(msgSlice)
}

func main() {
	InitDB()
	DB.AutoMigrate(&RequestBody{})

	router := mux.NewRouter()
	router.HandleFunc("/api/write", PostHandler).Methods("POST")
	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
