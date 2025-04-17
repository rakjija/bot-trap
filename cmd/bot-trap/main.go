package main

import (
	"github.com/rakjija/bot-trap/internal/api"
	"github.com/rakjija/bot-trap/internal/metrics"
)

func main() {
	metrics.Init()
	api.StartServer()
}
