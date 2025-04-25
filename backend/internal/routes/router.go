package routes

import (
	"github.com/gin-gonic/gin"

	docs "github.com/rakjija/go-board/backend/docs"
	"github.com/rakjija/go-board/backend/internal/handlers/meta"
	"github.com/rakjija/go-board/backend/internal/middleware"
	postRoutes "github.com/rakjija/go-board/backend/internal/routes/post"
	userRoutes "github.com/rakjija/go-board/backend/internal/routes/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {

	// Middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestLogger())

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/healthz", meta.HealthzHandler)

	// /api/v1
	v1 := r.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"

	userRoutes.RegisterUserRoutes(v1)
	postRoutes.RegisterPostRoutes(v1)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
