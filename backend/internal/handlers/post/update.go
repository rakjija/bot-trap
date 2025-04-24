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

// @Summary 게시글 수정
// @Description 게시글 작성자가 제목과 내용을 수정합니다.
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "게시글 ID"
// @Param post body types.PostCreateRequest true "수정할 게시글 정보"
// @Success 200 {object} types.PostResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 403 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /posts/{id} [put]
func UpdatePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid post ID"})
		return
	}

	var req types.PostCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	var post models.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse{Error: "post not found"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, types.ErrorResponse{Error: "permission denied"})
		return
	}

	post.Title = req.Title
	post.Content = req.Content

	if err := db.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to update post"})
		return
	}

	c.JSON(http.StatusOK, types.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	})
}
