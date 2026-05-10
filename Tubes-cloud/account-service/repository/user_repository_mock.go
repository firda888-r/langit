//Hasil generate Mock
// Di file repository/user_repository_mock.go
package repository

import "account-service/model"

type UserRepositoryMock struct{}

func (m *UserRepositoryMock) GetByEmail(email string) (*model.User, error) {
    // Kita simulasikan: jika emailnya 'test@upi.edu', maka ADA datanya
    if email == "test@upi.edu" {
        return &model.User{ID: 1, Email: "test@upi.edu", Role: "user"}, nil
    }
    return nil, nil
}