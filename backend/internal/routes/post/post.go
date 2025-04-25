package post

import (
	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/backend/internal/handlers/post"
	"github.com/rakjija/go-board/backend/internal/middleware"
)

func RegisterPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")
	posts.GET("", post.ListPosts)
	posts.GET("/:id", post.GetPost)

	auth := posts.Group("")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("", post.CreatePost)
	auth.PUT("/:id", post.UpdatePost)
	auth.DELETE("/:id", post.DeletePost)
}
