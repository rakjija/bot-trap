package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/observer/internal/models"
)

func HandleLog(c *gin.Context) {
	var entry models.LogEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid log format"})
		return
	}

	entry.Timestamp = time.Now().Format(time.RFC3339)

	file, err := os.OpenFile("/app/logs/logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot write log"})
		return
	}
	defer file.Close()

	jsonData, _ := json.Marshal(entry)
	file.Write(append(jsonData, '\n'))

	c.JSON(http.StatusOK, gin.H{"status": "log received"})
}
