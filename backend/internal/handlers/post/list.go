package post

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/backend/internal/db"
	"github.com/rakjija/go-board/backend/internal/models"
	"github.com/rakjija/go-board/backend/internal/types"
)

// @Summary 게시글 목록 조회
// @Description 전체 게시글을 최신순으로 조회합니다.
// @Tags posts
// @Produce json
// @Success 200 {array} types.PostResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /posts [get]
func ListPosts(c *gin.Context) {
	var posts []models.Post

	if err := db.DB.Preload("User").Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to fetch posts"})
		return
	}

	var response []types.PostResponse
	for _, p := range posts {
		response = append(response, types.PostResponse{
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
