package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hanson777/url-shortener/internal/service"
	"github.com/hanson777/url-shortener/internal/writer"
)

var ShortenURLRequest struct {
	URL string
}

type ShortenURLResponse struct {
	ShortURL string
	LongURL  string
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&ShortenURLRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if ShortenURLRequest.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	code, _ := service.InsertShortURL(ShortenURLRequest.URL)

	response := ShortenURLResponse{
		ShortURL: "http://localhost:8080/" + code.Code,
		LongURL:  ShortenURLRequest.URL,
	}

	err := writer.Write(w, http.StatusCreated, response)
	if err != nil {
		log.Fatalf("error encoding writer: %s", err)
	}
}
