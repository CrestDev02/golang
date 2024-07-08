package repository

import (
	"database/sql"
	"my-app/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, user.Username, user.Password)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username=$1"
	row := r.DB.QueryRow(query, username)
	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
