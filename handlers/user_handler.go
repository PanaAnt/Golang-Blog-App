package handlers

import (
	"blogApp/models"
	"blogApp/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	//the collection of request details
	var signUp models.User
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//calling the service layer
	err = h.Service.RegisterUser(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signUp)

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login models.User
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//call service layer
	token, user, err := h.Service.Login(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	json.NewEncoder(w).Encode(user)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintln("Welcome to the Blog App!")
	w.Write([]byte(message))
}

