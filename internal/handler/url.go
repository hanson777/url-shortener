package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hanson777/url-shortener/internal/service"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	response := service.ShortenURL(req.URL)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
