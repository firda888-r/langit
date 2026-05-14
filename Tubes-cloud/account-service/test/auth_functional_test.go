//FUNCTIONAL TEST (Koneksi ke DB asli)
package test

import (
	"net/http"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
    body := `{"email":"user@mail.com","password":"1234","role":"user"}`

    resp, err := http.Post(
        "http://localhost:8080/api/auth/login",
        "application/json",
        strings.NewReader(body),
    )

    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close() // Tambahkan ini

    if resp.StatusCode != http.StatusOK { // Menggunakan konstanta lebih rapi
        t.Errorf("Expected 200, got %d", resp.StatusCode)
    }
}