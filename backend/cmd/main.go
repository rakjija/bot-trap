package main

import (
	"log"

	"github.com/rakjija/bot-trap/backend/internal/config"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/routes"
	"github.com/rakjija/bot-trap/backend/internal/utils"
)

func main() {
	config.LoadConfig()

	db.Connect()
	db.Migrate()

	utils.RegisterCustomValidators()
	r := routes.InitRouter()
	r.Run(":8080")

	log.Printf("[INFO] Server is running on %s..", config.Config.Server.Port)
}
