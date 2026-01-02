package service

import (
	"context"
	"log"

	"github.com/hanson777/url-shortener/internal/sqlc"
	"github.com/sqids/sqids-go"
)

func GetLongUrlByCode(code string, queries *sqlc.Queries) (sqlc.Url, error) {
	ctx := context.Background()

	id := decodeBase62(code)
	fetchedLongUrl, err := queries.GetLongURL(ctx, id)
	if err != nil {
		log.Printf("error fetching url: %s, id: %d", err, id)
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

func InsertShortURL(longURL string, queries *sqlc.Queries) (string, error) {
	ctx := context.Background()

	insertedShortURL, err := queries.CreateShortURL(ctx, longURL)
	if err != nil {
		log.Printf("error creating ShortURL: %s", err)
		return "", err
	}
	return encodeBase62(insertedShortURL.ID), nil
}

func encodeBase62(id int64) string {
	s, _ := sqids.New()
	code, _ := s.Encode([]uint64{uint64(id)})
	return code
}
