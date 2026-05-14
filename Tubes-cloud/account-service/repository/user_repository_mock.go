package repository

import "account-service/model"

type UserRepositoryMock struct{}

func (m *UserRepositoryMock) GetByEmail(email string) (*model.User, error) {
	// Simulasi data ditemukan
	return &model.User{ID: 1, Email: "user@mail.com", Password: "123", Role: "user"}, nil
}