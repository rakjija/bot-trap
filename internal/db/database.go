package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type LogEntry struct {
	gorm.Model
	IP      string
	Path    string
	Message string
}

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=bottrap port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&LogEntry{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	fmt.Println("✅ Connected to PostgreSQL and migrated LogEntry table.")
}
