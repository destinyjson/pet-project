package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message RequestBody
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	DB.Create(&message)
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response []RequestBody
	DB.Find(&response)

	msgSlice := make([]string, len(response))
	for index, value := range response {
		msgSlice[index] = value.Message
	}
	json.NewEncoder(w).Encode(msgSlice)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response RequestBody
	id := mux.Vars(r)["id"]
	DB.First(&response, id)
	DB.Delete(&response, id)

	json.NewEncoder(w).Encode(response)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response RequestBody
	id := mux.Vars(r)["id"]
	DB.First(&response, id)
	json.NewDecoder(r.Body).Decode(&response)
	DB.Save(&response)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	InitDB()
	DB.AutoMigrate(&RequestBody{})

	router := mux.NewRouter()
	router.HandleFunc("/api/write", PostHandler).Methods("POST")
	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/update/{id}", UpdateHandler).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}
