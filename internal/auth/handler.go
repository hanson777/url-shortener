// Package auth provides authentication handlers, services, and middleware for user signup and login
package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hanson777/url-shortener/internal/writer"
)

type Handler struct {
	Service ServiceInterface
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
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
		http.Error(w, "Email or password is required", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Signup(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res := SignupResponse{Token: token}
	err = writer.Write(w, http.StatusCreated, res)
	if err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email or password is required", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res := LoginResponse{Token: token}
	err = writer.Write(w, http.StatusCreated, res)
	if err != nil {
		log.Printf("error writing respones: %v", err)
	}
}
