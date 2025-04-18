package main

import (
	"github.com/rakjija/bot-trap/internal/api"
	"github.com/rakjija/bot-trap/internal/metrics"
)

// @title BotTrap API
// @version 1.0
// @description 백엔드 이상 행동 탐지 시스템 API 문서입니다.
// @contact.name rakjija
// @contact.url https://github.com/rakjija/bot-trap
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	metrics.Init()
	api.StartServer()
}
