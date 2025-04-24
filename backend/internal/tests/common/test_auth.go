package common

import (
	"github.com/gin-gonic/gin"
)

func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", uint(1))
		c.Set("user_email", "test@example.com")

		c.Next()
	}
}
