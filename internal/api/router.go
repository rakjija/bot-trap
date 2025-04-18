package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(MetricsMiddleware())

	r.GET("/healthz", HealthzHandler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.POST("/logs", PostLogHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger.json")))
	r.Static("/docs", "/root/docs")

	return r
}

func StartServer() {
	r := NewRouter()
	r.Run(":8080")
}
