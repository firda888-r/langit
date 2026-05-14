//UNIT TEST (Tanpa koneksi DB asli)
//Menguji logika bisnis (misal: validasi email, logika hashing password) secara terisolasi

package service

import (
	"account-service/model"
	"account-service/repository"
	"testing"
	"errors"
)

// UNIT TEST: Menguji fungsi Login secara terisolasi tanpa database/jaringan
func TestLogin_Unit(t *testing.T) {
	// 1. Setup Mock Repository (Simulasi DB)
	// Untuk tahap 2, kita bisa pakai mock manual jika belum generate gomock
	mockRepo := &repository.UserRepositoryMock{} 
	
	// 2. Inisialisasi Service dengan Mock tersebut
	authSvc := NewAuthService(mockRepo)

	// 3. Data Test
	req := model.LoginRequest{
		Email:    "user@mail.com",
		Password: "1234",
		Role:     "user",
	}

	// 4. Eksekusi fungsi
	resp, err := authSvc.Login(req)

	// 5. Validasi hasil
	// Sesuai perintah dosen: karena kodenya belum selesai, kita bisa ekspektasikan error
	if err == nil {
		t.Errorf("Harusnya FAILED karena logic belum diimplementasi")
	}

	if resp != nil {
		t.Errorf("Response harusnya nil karena fitur belum selesai")
	}
}