package api

import "github.com/gin-gonic/gin"

// HealthzHandler handles the health check request
// @Summary 헬스 체크
// @Description 애플리케이션의 동작 여부를 확인합니다.
// @Tags system
// @Produce json
// @Success 200 {object} model.HealthResponse "서버 정상 동작 확인"
// @Router /healthz [get]
func HealthzHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
