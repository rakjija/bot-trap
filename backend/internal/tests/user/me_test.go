package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakjija/bot-trap/backend/internal/tests/common"
	"github.com/stretchr/testify/assert"
)

func TestMe_Success(t *testing.T) {
	common.InitTestDB()
	common.CreateTestUser()
	router := common.SetupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/users/me", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"user_id": 1}`, w.Body.String())
}
