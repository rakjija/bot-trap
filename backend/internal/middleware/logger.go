package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/types"
	"github.com/rakjija/bot-trap/backend/internal/utils"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // 요청 처리 수행

		duration := time.Since(start)
		status := c.Writer.Status()

		level := "info"
		msg := "Request completed"
		if status >= 500 {
			level = "error"
			msg = "Internal server error"
		} else if status >= 400 {
			level = "warn"
			msg = "Client error occurred"
		}

		log := types.LogPayload{
			Timestamp: start.Format(time.RFC3339),
			Duration:  duration.Seconds(),
			Level:     level,
			Service:   "backend",
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Status:    status,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Message:   msg,
			User:      "anonymous",
		}

		if userEmailVal, exists := c.Get("user_email"); exists {
			if userEmail, ok := userEmailVal.(string); ok {
				log.User = userEmail
			}
		}

		utils.SendStructuredLog(log)
	}
}
