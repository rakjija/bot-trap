package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakjija/bot-trap/backend/internal/tests/common"
	"github.com/rakjija/bot-trap/backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	common.InitTestDB()
	utils.RegisterCustomValidators()
	common.CreateTestUser()
	router := common.SetupTestRouter()

	body, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "pa55w0rd",
	})

	req, _ := http.NewRequest("POST", "/api/v1/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "access_token")
	assert.Contains(t, w.Body.String(), "user_id")
}
