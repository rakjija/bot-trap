package common

import (
	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/backend/internal/handlers/post"
	"github.com/rakjija/go-board/backend/internal/handlers/user"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.POST("/users/login", user.Login)
	v1.POST("/users/signup", user.Signup)
	v1.GET("/users/me", MockAuthMiddleware(), user.Me)

	v1.GET("/posts", post.ListPosts)
	v1.POST("/posts", MockAuthMiddleware(), post.CreatePost)
	v1.PUT("/posts/:id", MockAuthMiddleware(), post.UpdatePost)
	v1.GET("/posts/:id", post.GetPost)
	v1.DELETE("/posts/:id", MockAuthMiddleware(), post.DeletePost)

	return r
}
