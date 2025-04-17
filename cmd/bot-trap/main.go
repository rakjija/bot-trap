package main

import (
	"github.com/rakjija/bot-trap/internal/api"
	"github.com/rakjija/bot-trap/internal/db"
	"github.com/rakjija/bot-trap/internal/metrics"
)

func main() {
	db.ConnectDB()
	metrics.Init()
	api.StartServer()
}
