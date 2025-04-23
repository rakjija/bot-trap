package common

import (
	"fmt"

	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
)

func CreateTestPost(n int) {
	for i := 1; i <= n; i++ {
		db.DB.Create(&models.Post{
			Title:   fmt.Sprintf("test%d", i),
			Content: fmt.Sprintf("test%d", i),
			UserID:  1,
		})
	}
}
