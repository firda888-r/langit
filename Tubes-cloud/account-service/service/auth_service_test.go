//UNIT TEST (Tanpa koneksi DB asli)
//Menguji logika bisnis (misal: validasi email, logika hashing password) secara terisolasi

func TestLogin_Success(t *testing.T) {
    // Pakai Mock, jadi tidak butuh database PostgreSQL menyala
    mockRepo := &repository.UserRepositoryMock{} 
    service := NewAuthService(mockRepo)
    
    req := model.LoginRequest{Email: "test@upi.edu", Password: "123"}
    resp, err := service.Login(req)

    // Sekarang kita cek: Harusnya err itu nil (tidak error)
    if err != nil {
        t.Errorf("Harusnya sukses tapi malah error: %v", err)
    }

    if resp.Token == "" {
        t.Errorf("Token tidak boleh kosong")
    }
}