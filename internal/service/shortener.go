package service

import (
	"context"
	"log"
	"os"

	"github.com/hanson777/url-shortener/internal/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/sqids/sqids-go"
)

type URLEncoding struct {
	Code string
}

func InsertShortURL(longURL string) (*URLEncoding, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("error connecting to database")
		return &URLEncoding{}, err
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	insertedShortURL, err := queries.CreateShortURL(ctx, longURL)
	if err != nil {
		log.Fatalf("error creating ShortURL: %s", err)
		return &URLEncoding{}, err
	}

	return &URLEncoding{
		Code: encodeBase62(insertedShortURL.ID),
	}, nil
}

func encodeBase62(id int64) string {
	s, _ := sqids.New()
	code, _ := s.Encode([]uint64{uint64(id)})
	return code
}
