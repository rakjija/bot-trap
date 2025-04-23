package post

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakjija/bot-trap/backend/internal/tests/common"
	"github.com/stretchr/testify/assert"
)

func TestDeletePost_Success(t *testing.T) {
	common.InitTestDB()
	common.CreateTestUser()
	common.CreateTestPost(3)
	router := common.SetupTestRouter()

	req, _ := http.NewRequest("DELETE", "/api/v1/posts/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer mock-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "post deleted"}`, w.Body.String())

	// Post 개수 줄어듦(-1) 확인
	var posts []map[string]any
	getReq, _ := http.NewRequest("GET", "/api/v1/posts", nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusOK, getW.Code)

	err := json.Unmarshal(getW.Body.Bytes(), &posts)
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
}
