package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/hanson777/url-shortener/internal/handler"
	"github.com/hanson777/url-shortener/internal/middleware"
	"github.com/hanson777/url-shortener/internal/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/lpernett/godotenv"
)

func main() {
	log.Print("Starting...")
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	h := handler.NewHandler(queries)

	rateLimiter := middleware.NewRateLimiter()

	mux := http.NewServeMux()
	mux.Handle("POST /api/shorten", middleware.RateLimitEndpoint(rateLimiter, h.CreateShortURL))
	mux.HandleFunc("GET /{code}", h.Redirect)

	log.Print("Server listening on port 8080")
	err = http.ListenAndServe(":8080", middleware.Cors(mux))
	if err != nil {
		log.Fatal(err)
	}
}
