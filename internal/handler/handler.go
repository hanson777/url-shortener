package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/hanson777/url-shortener/internal/service"
	"github.com/hanson777/url-shortener/internal/sqlc"
	"github.com/hanson777/url-shortener/internal/writer"
)

type Handler struct {
	queries *sqlc.Queries
}

type ShortenURLRequest struct {
	URL string
}

type ShortenURLResponse struct {
	ShortURL string `json:"shortUrl"`
	LongURL  string `json:"longUrl"`
}

type RedirectResponse struct {
	Url string `json:"url"`
}

func NewHandler(queries *sqlc.Queries) *Handler {
	return &Handler{queries: queries}
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

	normalizedUrl := normalizeUrl(req.URL)

	if !isValidURL(normalizedUrl) {
		http.Error(w, "URL is not valid", http.StatusBadRequest)
		return
	}

	code, _ := service.InsertShortURL(normalizedUrl, h.queries)

	response := ShortenURLResponse{
		ShortURL: "http://localhost:8080/" + code,
		LongURL:  normalizedUrl,
	}

	err := writer.Write(w, http.StatusCreated, response)
	if err != nil {
		log.Printf("error encoding writer: %s", err)
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if code == "" {
		log.Print("Code must be non-empty")
		return
	}

	url, err := service.GetLongUrlByCode(code, h.queries)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	response := &RedirectResponse{
		url.LongUrl,
	}

	http.Redirect(w, r, response.Url, http.StatusPermanentRedirect)
}
