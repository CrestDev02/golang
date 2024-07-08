package migration

import (
	"database/sql"
	"fmt"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password VARCHAR(100) NOT NULL
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	} else {
		fmt.Println("Database migrated successfully.")
	}
}
