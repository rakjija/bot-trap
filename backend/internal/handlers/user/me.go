package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Access Token 인증 확인
// @Description Authorization 헤더의 JWT에서 user_id를 추출하여 반환합니다.
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} MeResponse
// @Failure 401 {object} ErrorResponse
// @Router /users/me [get]
func Me(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "user_id not found in context"})
		return
	}
	userID := userIDVal.(uint)

	c.JSON(http.StatusOK, MeResponse{UserID: userID})
}
