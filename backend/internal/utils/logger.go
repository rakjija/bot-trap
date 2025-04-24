package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rakjija/bot-trap/backend/internal/config"
	"github.com/rakjija/bot-trap/backend/internal/types"
)

func SendLog(level, message, user string) {

	payload := types.LogPayload{
		Level:   level,
		Message: message,
		User:    user,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[LOGGER] 로그 직렬화 실패: %v\n", err)
		return
	}

	res, err := http.Post(config.Config.Log.Endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("[LOGGER] 로그 전송 실패: %v\n", err)
		return
	}
	defer res.Body.Close()

	log.Printf("[LOGGER] 로그 전송 완료 (%s): %s\n", level, message)
}
