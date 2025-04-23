package common

import (
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
	"github.com/rakjija/bot-trap/backend/internal/utils"
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
