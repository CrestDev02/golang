package models

type User struct {
	ID          int    `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password" db:"password"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Address     string `json:"address" db:"address"`
	City        string `json:"city" db:"city"`
	State       string `json:"state" db:"state"`
	Country     string `json:"country" db:"country"`
	ZipCode     string `json:"zip_code" db:"zip_code"`
	IsVerified  bool   `json:"is_verified" db:"is_verified"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}
