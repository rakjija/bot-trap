package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/internal/metrics"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		metrics.HttpDuration.WithLabelValues(c.FullPath(), c.Request.Method).Observe(duration)
	}
}
