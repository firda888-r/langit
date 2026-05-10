package handler

import (
	"account-service/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Service *service.AuthService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	id, role, token, err := h.Service.Login(req.Email, req.Password, req.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := LoginResponse{
		UserID: id,
		Role:   role,
		Token:  token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}