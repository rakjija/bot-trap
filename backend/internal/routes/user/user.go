package user

import (
	"github.com/gin-gonic/gin"
	handler "github.com/rakjija/bot-trap/backend/internal/handlers/user"
	"github.com/rakjija/bot-trap/backend/internal/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")

	user.POST("/signup", handler.Signup)
	user.POST("/login", handler.Login)

	auth := user.Group("")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/me", handler.Me)
}
