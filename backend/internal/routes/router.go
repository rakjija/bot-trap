package routes

import (
	"github.com/gin-gonic/gin"

	docs "github.com/rakjija/bot-trap/backend/docs"
	"github.com/rakjija/bot-trap/backend/internal/handlers/healthz"
	"github.com/rakjija/bot-trap/backend/internal/middleware"
	postRoutes "github.com/rakjija/bot-trap/backend/internal/routes/post"
	userRoutes "github.com/rakjija/bot-trap/backend/internal/routes/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {

	// Middleware
	r.Use(middleware.CORSMiddleware())

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/healthz", healthz.HealthzHandler)

	// /api/v1
	v1 := r.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"

	userRoutes.RegisterUserRoutes(v1)
	postRoutes.RegisterPostRoutes(v1)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
