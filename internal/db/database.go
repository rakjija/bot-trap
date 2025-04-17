package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=127.0.0.1 user=postgres password=postgres dbname=bottrap port=5432 sslmode=disable"
	}
	fmt.Println("📡 DB 접속 대상 DSN:", dsn)
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
