package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(MetricsMiddleware())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/logs", PostLogHandler)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}

func StartServer() {
	r := NewRouter()
	r.Run(":8080")
}
