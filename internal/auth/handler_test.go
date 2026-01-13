package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
	SignupFunc func(ctx context.Context, email string, password string) (string, error)
	LoginFunc  func(ctx context.Context, email string, password string) (string, error)
}

func (m *MockService) Signup(ctx context.Context, email string, password string) (string, error) {
	if m.SignupFunc != nil {
		return m.SignupFunc(ctx, email, password)
	}
	return "", nil
}

func (m *MockService) Login(ctx context.Context, email string, password string) (string, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(ctx, email, password)
	}
	return "", nil
}

func TestSignup(t *testing.T) {
	ms := &MockService{
		SignupFunc: func(ctx context.Context, email string, password string) (string, error) {
			return "token", nil
		},
	}

	h := NewHandler(ms)
	reqBody := SignupRequest{
		Email:    "johndoe@gmail.com",
		Password: "abc123",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	h.Signup(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
	var res SignupResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if res.Token != "token" {
		t.Errorf("expected token abc123, got %s", res.Token)
	}
}
