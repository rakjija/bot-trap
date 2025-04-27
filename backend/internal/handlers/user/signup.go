package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/backend/internal/db"
	"github.com/rakjija/go-board/backend/internal/models"
	"github.com/rakjija/go-board/backend/internal/types"
	"github.com/rakjija/go-board/backend/internal/utils"
)

// @Summary 회원가입
// @Description 새 유저를 등록합니다.
// @Tags users
// @Accept json
// @Produce json
// @Param signup body types.SignupRequest true "회원가입 요청"
// @Success 201 {object} types.SignupResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /users/signup [post]
func Signup(c *gin.Context) {
	var req types.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to hash password"})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Username: req.Username,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, types.SignupResponse{
		UserID:  user.ID,
		Message: "user created successfully",
	})
}
