package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pet-project/internal/service"
	"strconv"
)

type Handler struct {
	Service *service.MessageService
}

func NewHandler(service *service.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message service.RequestBody
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message service.RequestBody
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedMsg, err := h.Service.DeleteMessageByID(idInt, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedMsg)
}

func (h *Handler) UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message service.RequestBody
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMessage, err := h.Service.UpdateMessageByID(idInt, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMessage)
}
