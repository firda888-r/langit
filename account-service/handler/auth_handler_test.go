package handler

import (
	"account-service/model"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAuthService struct{}

func (m MockAuthService) Login(req model.LoginRequest) (*model.LoginResponse, error) {

	return &model.LoginResponse{
		Token: "dummy-token",
	}, nil
}

func TestLoginHandler(t *testing.T) {

	jsonBody := []byte(`{
		"email":"firda@gmail.com",
		"password":"123"
	}`)

	req := httptest.NewRequest(
		"POST",
		"/login",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mockService := MockAuthService{}

	handler := NewAuthHandler(mockService)

	handler.LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("status harus 200")
	}
}