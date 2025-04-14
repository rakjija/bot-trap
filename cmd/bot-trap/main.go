package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() // 기본 설정된 라우터 생성 (로깅, 리커버리 포함)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // JSON 응답 반환
			"status": "ok",
		})
	})

	r.Run(":8080")
}
