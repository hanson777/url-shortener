package service

import (
	"context"
	"log"

	"github.com/hanson777/url-shortener/internal/sqlc"
)

type ServiceInterface interface {
	GetLongURLByCode(ctx context.Context, code string) (sqlc.Url, error)
	InsertShortURL(ctx context.Context, longURL string) (string, error)
	IncrementClicks(ctx context.Context, ID int64) error
}

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) GetLongURLByCode(ctx context.Context, code string) (sqlc.Url, error) {
	id := decodeBase62(code)
	fetchedLongURL, err := s.queries.GetLongURL(ctx, id)
	if err != nil {
		log.Printf("error fetching url: %v, id: %d", err, id)
	}

	return fetchedLongURL, nil
}

func (s *Service) InsertShortURL(ctx context.Context, longURL string) (string, error) {
	insertedShortURL, err := s.queries.CreateShortURL(ctx, longURL)
	if err != nil {
		log.Printf("error creating ShortURL: %v", err)
		return "", err
	}
	return encodeBase62(insertedShortURL.ID), nil
}

func (s *Service) IncrementClicks(ctx context.Context, ID int64) error {
	err := s.queries.IncrementClicks(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
