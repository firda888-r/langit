package repository

import (
	"account-service/model"
	"database/sql"
)

// Interface ini wajib ada untuk Mocking
type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	// Untuk Tahap 2: Biarkan kosong atau return nil untuk memicu error di test
	return nil, nil 
}