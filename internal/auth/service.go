package auth

import (
	"context"
	"log"

	"github.com/hanson777/url-shortener/internal/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	queries *sqlc.Queries
}

type ServiceInterface interface {
	Signup(ctx context.Context, email string, password string) string
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) Signup(ctx context.Context, email string, password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 17)
	if err != nil {
		log.Printf("error encrypting password: %v", err)
	}

	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hashed),
	})
	if err != nil {
		log.Printf("error creating user: %v", err)
	}
	token := GenerateToken(string(user.ID))
	return token
}
