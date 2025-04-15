package main

import (
	"github.com/rakjija/bot-trap/internal/api"
	"github.com/rakjija/bot-trap/internal/db"
)

func main() {
	db.ConnectDB()
	api.StartServer()
}
