package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/hanson777/url-shortener/internal/handler"
	"github.com/jackc/pgx/v5"
	"github.com/lpernett/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All systems nominal"))
	})

	mux.HandleFunc("POST /api/shorten", handler.CreateShortURL)

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	log.Print("Server listening on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
