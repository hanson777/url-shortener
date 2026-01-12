package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/hanson777/url-shortener/internal/service"
	"github.com/hanson777/url-shortener/internal/writer"
)

type Handler struct {
	Service service.ServiceInterface
}

type ShortenURLRequest struct {
	URL string
}

type ShortenURLResponse struct {
	ShortURL string `json:"shortUrl"`
	LongURL  string `json:"longUrl"`
}

func NewHandler(service service.ServiceInterface) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	normalizedURL := normalizeURL(req.URL)

	if !isValidURL(normalizedURL) {
		http.Error(w, "URL is not valid", http.StatusBadRequest)
		return
	}

	code, _ := h.Service.InsertShortURL(r.Context(), normalizedURL)

	response := ShortenURLResponse{
		ShortURL: "http://localhost:8080/" + code,
		LongURL:  normalizedURL,
	}

	err := writer.Write(w, http.StatusCreated, response)
	if err != nil {
		log.Printf("error encoding writer: %v", err)
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if code == "" {
		log.Print("Code must be non-empty")
		return
	}

	url, err := h.Service.GetLongURLByCode(r.Context(), code)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = h.Service.IncrementClicks(r.Context(), url.ID)
	if err != nil {
		log.Printf("failed to increment clicks: %v", err)
	}

	http.Redirect(w, r, url.LongUrl, http.StatusFound)
}
