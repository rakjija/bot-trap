package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/middleware"
	postRoutes "github.com/rakjija/bot-trap/backend/internal/routes/post"
	userRoutes "github.com/rakjija/bot-trap/backend/internal/routes/user"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	userRoutes.RegisterUserRoutes(api)
	postRoutes.RegisterPostRoutes(api)

	return r
}
