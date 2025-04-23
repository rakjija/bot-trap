package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakjija/bot-trap/backend/internal/tests/common"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePost_Success(t *testing.T) {
	common.InitTestDB()
	common.CreateTestUser()
	common.CreateTestPost(1)
	router := common.SetupTestRouter()

	body, _ := json.Marshal(map[string]string{
		"title":   "test title",
		"content": "test content",
	})

	req, _ := http.NewRequest("PUT", "/api/v1/posts/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer mock-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "id")
	assert.Contains(t, w.Body.String(), "title")
	assert.Contains(t, w.Body.String(), "content")
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "created_at")
}
