package model

// HealthResponse 서버 상태 응답 구조체
// @Description 서버 헬스 체크 응답 예시
type HealthResponse struct {
	Status string `json:"status" example:"ok"` // 서버 상태
}
