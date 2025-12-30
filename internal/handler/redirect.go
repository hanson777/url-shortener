package handler

import (
	"log"
	"net/http"

	"github.com/hanson777/url-shortener/internal/service"
)

type RedirectResponse struct {
	Url string
}

func Redirect(w http.ResponseWriter, r *http.Request) {

	code := r.PathValue("code")

	if code == "" {
		log.Fatal("Code must be non-empty")
		return
	}

	url, err := service.GetLongUrlByCode(code)
	if err != nil {
		log.Fatalf("error fetching url: %s", err)
	}

	response := &RedirectResponse{
		url.LongUrl,
	}

	http.Redirect(w, r, response.Url, http.StatusPermanentRedirect)
}
