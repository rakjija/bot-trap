package db

import (
	"log"

	"github.com/rakjija/go-board/backend/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to migrate database schema: %v", err)
	}
	log.Println("[INFO] Database schema migration completed successfully")
}
