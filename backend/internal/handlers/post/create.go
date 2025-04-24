package post

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
	"github.com/rakjija/bot-trap/backend/internal/types"
)

// @Summary 게시글 작성
// @Description 인증된 사용자가 새로운 게시글을 작성합니다.
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param post body types.PostCreateRequest true "게시글 작성 요청"
// @Success 201 {object} types.PostResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	var req types.PostCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, types.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	})
}
