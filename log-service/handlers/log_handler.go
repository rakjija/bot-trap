package handlers

import (
	"log-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func SaveLog(c *gin.Context) {
	var logEntry models.LogEntry
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.InsertLog(logEntry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save log"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}

func GetLogs(c *gin.Context) {
	logs, err := models.FetchLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch logs"})
		return
	}
	c.JSON(http.StatusOK, logs)
}
