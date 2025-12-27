package service

import (
	"time"

	"github.com/deatil/go-encoding/encoding"
)

type ShortURL struct {
	Code        string
	OriginalURL string
	CreatedAt   time.Time
}

func ShortenURL(originalURL string) *ShortURL {
	code := generateShortURL(originalURL)
	return &ShortURL{
		Code:        code,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
	}
}

func generateShortURL(url string) string {
	code := encoding.FromString(url).Base62Encode().ToString()
	return code
}
