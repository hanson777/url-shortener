package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/hanson777/url-shortener/internal/auth"
	"github.com/hanson777/url-shortener/internal/handler"
	"github.com/hanson777/url-shortener/internal/middleware"
	"github.com/hanson777/url-shortener/internal/service"
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

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	service := service.NewService(queries)
	authService := auth.NewService(queries)

	h := handler.NewHandler(service)
	authHandler := auth.NewHandler(authService)

	rateLimiter := middleware.NewRateLimiter()

	mux := http.NewServeMux()
	mux.Handle("POST /auth/signup", middleware.RateLimitEndpoint(rateLimiter, authHandler.Signup))
	mux.Handle("POST /auth/login", middleware.RateLimitEndpoint(rateLimiter, authHandler.Login))
	mux.Handle("POST /api/shorten", middleware.RateLimitEndpoint(rateLimiter, h.CreateShortURL))
	mux.HandleFunc("GET /{code}", h.Redirect)

	log.Print("Server listening on port 8080")
	err = http.ListenAndServe(":8080", middleware.Cors(mux))
	if err != nil {
		log.Fatal(err)
	}
}
