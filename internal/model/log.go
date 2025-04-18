package model

// LogPayload 로그 저장 요청 바디
// @Description JSON 형식의 로그를 수신하고 저장합니다 (stdout + metrics)
// @name LogPayload
type LogPayload struct {
	IP      string `json:"ip" example:"192.168.0.1" binding:"required"`                // 클라이언트의 IP 주소
	Path    string `json:"path" example:"/admin" binding:"required"`                   // 접근한 경로
	Message string `json:"message" example:"SQL injection attempt" binding:"required"` // 의심되는 메시지
}
