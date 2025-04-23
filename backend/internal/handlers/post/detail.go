package post

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
)

// @Summary 게시글 상세 조회
// @Description 특정 ID에 해당하는 게시글을 조회합니다.
// @Tags post
// @Produce json
// @Param id path int true "게시글 ID"
// @Success 200 {object} post.PostResponse
// @Failure 400 {object} post.ErrorResponse
// @Failure 404 {object} post.ErrorResponse
// @Router /posts/{id} [get]
func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid post ID"})
		return
	}

	var post models.Post
	if err := db.DB.Preload("User").First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "post not found"})
		return
	}

	c.JSON(http.StatusOK, PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	})
}
