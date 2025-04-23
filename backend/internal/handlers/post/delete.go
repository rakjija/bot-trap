package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
)

// @Summary 게시글 삭제
// @Description 게시글 작성자가 게시글을 삭제합니다.
// @Tags post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "게시글 ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid post ID"})
		return
	}

	var post models.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "post not found"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: "permission denied"})
		return
	}

	if err := db.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}
