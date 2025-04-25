package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rakjija/go-board/backend/internal/config"
	"github.com/rakjija/go-board/backend/internal/db"
	"github.com/rakjija/go-board/backend/internal/routes"
	"github.com/rakjija/go-board/backend/internal/utils"
)

// @title GoBoard API
// @version 1.0
// @description Swagger for GoBoard

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
