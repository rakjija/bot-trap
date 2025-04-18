package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/internal/metrics"
	"github.com/rakjija/bot-trap/internal/model"
)

// PostLogHandler
// @Summary 로그 저장
// @Description JSON 형식의 로그를 수신하고 저장 (stdout + metrics)
// @Tags logs
// @Accept json
// @Produce json
// @Param log body model.LogPayload true "로그 내용"
// @Success 200 {object} map[string]string "성공 응답"
// @Failure 400 {object} map[string]string "잘못된 요청"
// @Failure 500 {object} map[string]string "서버 에러"
// @Router /logs [post]
func PostLogHandler(c *gin.Context) {
	var req model.LogPayload
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

	log.Printf("[LOG] IP: %s | Path: %s | Msg: %s", req.IP, req.Path, req.Message)

	metrics.LogSaveCounter.Inc()
	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}
