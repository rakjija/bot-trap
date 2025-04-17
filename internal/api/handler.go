package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/internal/db"
	"github.com/rakjija/bot-trap/internal/metrics"
)

type LogRequest struct {
	IP      string `json:"ip"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

func PostLogHandler(c *gin.Context) {
	var req db.LogEntry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 의심 로그 탐지
	msg := strings.ToLower(req.Message)
	path := strings.ToLower(req.Path)
	if strings.Contains(msg, "bot") || strings.Contains(msg, "sql") || strings.Contains(path, "/admin") {
		metrics.SuspiciousLogCounter.Inc()
	}

	// 로그 저장
	if err := db.DB.Create(&req).Error; err != nil {
		metrics.LogErrorCounter.Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save log"})
		return
	}

	metrics.LogSaveCounter.Inc()
	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}
