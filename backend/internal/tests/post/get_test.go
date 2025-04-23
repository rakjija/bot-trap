package post

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakjija/bot-trap/backend/internal/tests/common"
	"github.com/stretchr/testify/assert"
)

func TestGetPost_Success(t *testing.T) {
	common.InitTestDB()
	common.CreateTestUser()
	common.CreateTestPost(1)
	router := common.SetupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/posts/1", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "id")
	assert.Contains(t, w.Body.String(), "title")
	assert.Contains(t, w.Body.String(), "content")
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "username")
	assert.Contains(t, w.Body.String(), "created_at")
}
