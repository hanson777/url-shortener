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

var ShortenURLRequest struct {
	URL string
}

type ShortenURLResponse struct {
	ShortURL string
	LongURL  string
}

type RedirectResponse struct {
	Url string
}

func NewHandler(queries *sqlc.Queries) *Handler {
	return &Handler{queries: queries}
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&ShortenURLRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if ShortenURLRequest.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	code, _ := service.InsertShortURL(ShortenURLRequest.URL, h.queries)

	response := ShortenURLResponse{
		ShortURL: "http://localhost:8080/" + code.Code,
		LongURL:  ShortenURLRequest.URL,
	}

	err := writer.Write(w, http.StatusCreated, response)
	if err != nil {
		log.Fatalf("error encoding writer: %s", err)
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	code := r.PathValue("code")

	if code == "" {
		log.Fatal("Code must be non-empty")
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
