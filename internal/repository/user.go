package repository

import (
	"database/sql"
	"my-app/internal/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type UserRepositoryPostgres struct {
	DB *sql.DB
}

func (r *UserRepositoryPostgres) CreateUser(user models.User) error {
	query := `
    INSERT INTO users (username, password, first_name, last_name, email, phone_number, address, city, state, country, zip_code, is_verified)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := r.DB.Exec(query, user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Address, user.City, user.State, user.Country, user.ZipCode, user.IsVerified)
	return err
}

func (r *UserRepositoryPostgres) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, password, first_name, last_name, email, phone_number, address, city, state, country, zip_code, is_verified, created_at, updated_at FROM users WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Address, &user.City, &user.State, &user.Country, &user.ZipCode, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
