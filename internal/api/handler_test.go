package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/internal/db"
	"github.com/rakjija/bot-trap/internal/metrics"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	metrics.Init() // 메트릭 초기화

	// 인메모리 DB
	testDB, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = testDB.AutoMigrate(&db.LogEntry{})
	db.DB = testDB

	r.POST("/logs", PostLogHandler)

	return r
}

func TestLogAPI(t *testing.T) {
	r := setupTestRouter()

	t.Run("성공: 정상 요청", func(t *testing.T) {
		body := []byte(`{"ip": "1.1.1.1", "path": "/test", "message": "hello bot"}`)
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
		body := []byte(`{`) // JSON 문법 오류
		req, _ := http.NewRequest("POST", "/logs", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("got %d, want 400", resp.Code)
		}
	})
}
