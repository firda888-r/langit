package test

import (
	"net/http"
	"strings"
	"testing"
)

func TestLogin_Functional(t *testing.T) {
	// Menembak aplikasi yang sedang running
	body := `{"email":"user@mail.com","password":"123"}`
	resp, err := http.Post("http://localhost:8080/api/auth/login", "application/json", strings.NewReader(body))

	if err != nil {
		t.Fatalf("Server mati: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
}