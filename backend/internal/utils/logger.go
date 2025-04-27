package utils

import (
	"encoding/json"
	"log"

	"github.com/rakjija/go-board/backend/internal/types"
)

func SendStructuredLog(payload types.LogPayload) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[LOG] 로그 직렬화 실패: %v\n", err)
		return
	}

	log.Println(string(data))
}
