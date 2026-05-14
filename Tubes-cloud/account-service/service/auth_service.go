package service

import (
	"account-service/model"
	"account-service/repository"
	"errors"
)

type AuthService interface {
	Login(req model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	repo repository.UserRepository // Mengacu ke interface
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	// Logika sengaja dibuat gagal sesuai instruksi Tahap 2
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid login: data not found (failed by design)")
	}
	return &model.LoginResponse{UserID: user.ID, Token: "success"}, nil
}