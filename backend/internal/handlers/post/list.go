package post

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
)

func ListPosts(c *gin.Context) {
	var posts []models.Post

	if err := db.DB.Preload("User").Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
		return
	}

	var response []PostResponse
	for _, p := range posts {
		response = append(response, PostResponse{
			ID:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			UserID:    p.UserID,
			Username:  p.User.Username,
			CreatedAt: p.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
