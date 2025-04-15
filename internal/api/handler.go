package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakjija/bot-trap/internal/db"
)

type LogRequest struct {
	IP      string `json:"ip"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

// 커스텀 카운터 메트릭 선언
var logSaveCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "log_save_total",
		Help: "총 저장된 로그 수",
	},
)

// 메트릭 등록
func init() {
	prometheus.MustRegister(logSaveCounter)
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

		logSaveCounter.Inc() // 로그 저장 시 메트릭 증가

		c.JSON(http.StatusOK, gin.H{"status": "saved"})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8080")
}
