package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/bot-trap/backend/internal/config"
	"github.com/rakjija/bot-trap/backend/internal/db"
	"github.com/rakjija/bot-trap/backend/internal/routes"
	"github.com/rakjija/bot-trap/backend/internal/utils"
)

// @title BotTrap API
// @version 1.0
// @description Swagger for BotTrap

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadConfig()

	db.Connect()
	db.Migrate()

	r := gin.Default()

	routes.InitRouter(r)
	utils.RegisterCustomValidators()
	r.Run(":8080")

	log.Printf("[INFO] Server is running on %s..", config.Config.Server.Port)
}
