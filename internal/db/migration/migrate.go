package migration

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password VARCHAR(100) NOT NULL,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        email VARCHAR(100) UNIQUE NOT NULL,
        phone_number VARCHAR(15),
        address VARCHAR(255),
        city VARCHAR(50),
        state VARCHAR(50),
        country VARCHAR(50),
        zip_code VARCHAR(10),
        is_verified BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	} else {
		log.Println("Database migrated successfully.")
	}
}
