package common

import (
	"log"

	"github.com/rakjija/go-board/backend/internal/db"
	"github.com/rakjija/go-board/backend/internal/models"
	"github.com/rakjija/go-board/backend/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDB() *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test DB: %v", err)
	}

	err = testDB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("failed to migrate test DB: %v", err)
	}

	utils.RegisterCustomValidators()

	testDB.Exec("DELETE FROM posts")
	testDB.Exec("DELETE FROM users")

	db.DB = testDB
	return testDB
}
