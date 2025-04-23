package healthz

import "github.com/gin-gonic/gin"

// @Summary 헬스 체크
// @Description 서버가 정상 작동 중인지 확인합니다.
// @Tags health
// @Success 200  {object}  HealthResponse
// @Router /healthz [get]
func HealthzHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
