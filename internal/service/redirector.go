package service

import (
	"context"
	"log"
	"os"

	"github.com/hanson777/url-shortener/internal/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/sqids/sqids-go"
)

type URLDecoding struct {
	LongUrl string
}

func GetLongUrlByCode(code string) (sqlc.Url, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	id := decodeBase62(code)

	fetchedLongUrl, err := queries.GetLongURL(ctx, id)
	if err != nil {
		log.Fatalf("error fetching url: %s, id: %d", err, id)
	}

	return fetchedLongUrl, nil
}

func decodeBase62(code string) int64 {
	s, _ := sqids.New()
	idArray := s.Decode(code)
	var joinedId uint64
	for _, digit := range idArray {
		joinedId = joinedId*10 + digit
	}
	return int64(joinedId)
}
