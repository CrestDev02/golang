package main

import (
	"log"
	"my-app/config"
	"my-app/internal/db"

	router "my-app/internal/router"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := db.NewPostgresDB(config.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Start server
	s := router.NewServer(config, db)
	if err := s.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
