package common

import (
	"github.com/rakjija/go-board/backend/internal/db"
	"github.com/rakjija/go-board/backend/internal/models"
	"github.com/rakjija/go-board/backend/internal/utils"
)

func CreateTestUser() {
	hashedPassword, _ := utils.HashPassword("pa55w0rd")
	db.DB.Create(&models.User{
		ID:       1,
		Email:    "test@example.com",
		Password: hashedPassword,
		Username: "testuser",
	})
}
