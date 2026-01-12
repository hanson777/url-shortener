package auth

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service ServiceInterface
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and/or password is required", http.StatusBadRequest)
		return
	}
}
