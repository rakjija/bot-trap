package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/internal/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 테스트용 라우터 구성 함수
func setupTestRouter() *gin.Engine {
	r := gin.Default()

	testDB, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = testDB.AutoMigrate(&db.LogEntry{})
	db.DB = testDB

	r.POST("/logs", func(c *gin.Context) {
		var req db.LogEntry
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.DB.Create(&req).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "saved"})
	})

	return r
}

func TestLogAPI(t *testing.T) {
	r := setupTestRouter()

	t.Run("성공: 정상 요청", func(t *testing.T) {
		body := []byte(`{"ip": "1.1.1.1", "path": "/health", "message": "hi"}`)
		req, _ := http.NewRequest("POST", "/logs", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("got %d, want 200", resp.Code)
		}
	})

	t.Run("실패: 필드 누락", func(t *testing.T) {
		body := []byte(`{"ip": "1.1.1.1"}`) // path, message 없음
		req, _ := http.NewRequest("POST", "/logs", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("got %d, want 400", resp.Code)
		}
	})

	t.Run("실패: 잘못된 JSON", func(t *testing.T) {
		body := []byte(`{`) // 문법 오류
		req, _ := http.NewRequest("POST", "/logs", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("got %d, want 400", resp.Code)
		}
	})
}
