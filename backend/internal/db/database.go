package db

import (
	"fmt"
	"log"

	"github.com/rakjija/go-board/backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DB.User,
		config.Config.DB.Password,
		config.Config.DB.Host,
		config.Config.DB.Port,
		config.Config.DB.Database,
	)

	var gormErr error
	DB, gormErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if gormErr != nil {
		log.Fatal("[ERROR] Failed to connect to database:", gormErr)
	}

	log.Println("[INFO] Successfully connected to the database")
}
