package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/models"
	"github.com/rakjija/bot-trap/backend/internal/types"
	"github.com/rakjija/bot-trap/backend/internal/utils"
)

// @Summary 로그인
// @Description 이메일과 비밀번호로 로그인합니다.
// @Tags users
// @Accept json
// @Produce json
// @Param login body types.LoginRequest true "로그인 요청"
// @Success 200 {object} types.LoginResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /users/login [post]
func Login(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
	}

	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to generate token"})
		return
	}

	c.Set("user_email", user.Email)

	c.JSON(http.StatusOK, types.LoginResponse{
		UserID:      user.ID,
		AccessToken: token,
	})
}
