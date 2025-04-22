package routes

import (
	"github.com/gin-gonic/gin"
	postRoutes "github.com/rakjija/bot-trap/backend/internal/routes/post"
	userRoutes "github.com/rakjija/bot-trap/backend/internal/routes/user"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	userRoutes.RegisterUserRoutes(api)
	postRoutes.RegisterPostRoutes(api)

	return r
}
