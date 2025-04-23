package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakjija/bot-trap/backend/internal/tests/common"
	"github.com/stretchr/testify/assert"
)

func TestSignup_Success(t *testing.T) {
	common.InitTestDB()
	router := common.SetupTestRouter()

	body, _ := json.Marshal(map[string]string{
		"email":    "newuser@example.com",
		"password": "pa55w0rd",
		"username": "newuser",
	})

	req, _ := http.NewRequest("POST", "/api/v1/users/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "user created successfully")
}
