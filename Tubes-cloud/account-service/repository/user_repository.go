package repository

import (
	"account-service/model"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) FindByEmailAndRole(email, role string) (*model.User, error) {
	user := &model.User{}

	query := "SELECT id, email, password, role FROM users WHERE email=$1 AND role=$2"

	err := r.DB.QueryRow(query, email, role).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}