package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lpernett/godotenv"
)

func getKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("JWT_SECRET")
}

func GenerateToken(userID string) string {
	var (
		key    []byte
		t      *jwt.Token
		signed string
	)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
			"iat":     time.Now().Unix(),
			"jti":     uuid.New().String(),
		})
	key = []byte(getKey())
	signed, _ = t.SignedString(key)
	return signed
}
