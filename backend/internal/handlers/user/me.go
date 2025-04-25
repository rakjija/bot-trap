package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/backend/internal/types"
)

// @Summary Access Token 인증 확인
// @Description Authorization 헤더의 JWT에서 user_id를 추출하여 반환합니다.
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} types.MeResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /users/me [get]
func Me(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	c.JSON(http.StatusOK, types.MeResponse{UserID: userID})
}
