package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LogRequest struct {
	IP      string `json:"ip"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

func StartServer() {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/logs", func(c *gin.Context) {
		var req LogRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[LOG] IP: %s | Path: %s | Msg: %s", req.IP, req.Path, req.Message)

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	r.Run(":8080")
}
