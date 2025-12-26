package main

import (
	"log"
	"net/http"

	"github.com/hanson777/url-shortener/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All systems nominal"))
	})

	mux.HandleFunc("POST /api/shorten", handler.CreateShortURL)

	log.Print("Server listening on port 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
