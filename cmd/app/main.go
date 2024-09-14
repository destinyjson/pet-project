package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/service"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&service.RequestBody{})

	repo := service.NewMessageRepository(database.DB)
	serviceMsg := service.NewService(repo)

	handler := handlers.NewHandler(serviceMsg)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/write", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/update/{id}", handler.UpdateMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
