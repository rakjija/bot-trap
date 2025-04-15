package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/internal/db"
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
		var req db.LogEntry
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// log.Printf("[LOG] IP: %s | Path: %s | Msg: %s", req.IP, req.Path, req.Message)
		if err := db.DB.Create(&req).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save log"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "saved"})
	})

	r.Run(":8080")
}
