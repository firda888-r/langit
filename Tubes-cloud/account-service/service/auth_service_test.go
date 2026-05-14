package service

import (
	"account-service/model"
	"account-service/repository"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	// Menggunakan Mock Repository
	mockRepo := &repository.UserRepositoryMock{}
	svc := NewAuthService(mockRepo)

	req := model.LoginRequest{Email: "user@mail.com", Password: "123"}
	
	// Panggil fungsi langsung
	resp, err := svc.Login(req)

	// Validasi agar PASS
	if err != nil {
		t.Errorf("Harusnya PASS, tapi dapat error: %v", err)
	}
	if resp.Token == "" {
		t.Errorf("Token tidak boleh kosong")
	}
}