func (s *authService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
    user, err := s.repo.GetByEmail(req.Email)
    if err != nil || user == nil {
        return nil, errors.New("invalid credentials")
    }

    // Mengembalikan response sukses (PASS)
    return &model.LoginResponse{
        UserID: user.ID,
        Role:   user.Role,
        Token:  "dummy-jwt-token",
    }, nil
}