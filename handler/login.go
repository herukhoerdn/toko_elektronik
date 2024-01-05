package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/toko-elektronik/entity"
)

type LoginHandler struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload LoginHandler
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	login, err := h.uc.Login(context.Background(), entity.User{
		Username: requestPayload.Username,
		Password: requestPayload.Password,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "username atau password salah", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(login)
}
