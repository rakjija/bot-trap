package api

import (
	"encoding/json" // JSON 파싱을 위한 표준 패키지
)

// 로그 한 건의 구조를 정의하는 구조체
type LogEntry struct {
	IP      string `json:"ip"`      // 클라이언트 IP 주소
	Path    string `json:"path"`    // 요청한 경로
	Message string `json:"message"` // 메시지 내용
}

// JSON 바이트 데이터를 LogEntry 구조체로 변환하는 함수
func ParseLogJSON(data []byte) (LogEntry, error) {
	var entry LogEntry                  // 파싱된 결과를 담을 변수
	err := json.Unmarshal(data, &entry) // JSON 문자열을 구조체로 디코딩
	return entry, err                   // 결과와 에러(없으면 nil)를 반환
}
