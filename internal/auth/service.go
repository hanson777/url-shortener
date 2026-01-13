package auth

import (
	"context"
	"errors"
	"log"

	"github.com/hanson777/url-shortener/internal/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	queries *sqlc.Queries
}

type ServiceInterface interface {
	Signup(ctx context.Context, email string, password string) (string, error)
	Login(ctx context.Context, email string, password string) (string, error)
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) Signup(ctx context.Context, email string, password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 17)
	if err != nil {
		log.Printf("error encrypting password: %v", err)
	}

	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hashed),
	})
	if err != nil {
		return "", errors.New("error creating user")
	}
	token := GenerateToken(string(user.ID))
	return token, nil
}

func (s *Service) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token := GenerateToken(string(user.ID))
	return token, nil
}
