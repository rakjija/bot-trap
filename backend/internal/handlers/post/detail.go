package post

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
	"github.com/rakjija/bot-trap/backend/internal/types"
)

// @Summary 게시글 상세 조회
// @Description 특정 ID에 해당하는 게시글을 조회합니다.
// @Tags post
// @Produce json
// @Param id path int true "게시글 ID"
// @Success 200 {object} types.PostResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /posts/{id} [get]
func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid post ID"})
		return
	}

	var post models.Post
	if err := db.DB.Preload("User").First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse{Error: "post not found"})
		return
	}

	c.JSON(http.StatusOK, types.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	})
}
