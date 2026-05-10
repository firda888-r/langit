package main

import (
	"account-service/handler"
	"account-service/repository"
	"account-service/service"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=1234 dbname=account_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repository.UserRepository{DB: db}
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{Service: authService}

	http.HandleFunc("/api/auth/login", authHandler.Login)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}