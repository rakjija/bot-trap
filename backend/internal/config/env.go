package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ServerConfig struct {
	Port string
}

var Config AppConfig

func LoadConfig() {
	rootPath, _ := filepath.Abs("/app/.env")
	err := godotenv.Load(rootPath)
	if err != nil {
		log.Println("[WARN] .env 파일을 찾을 수 없습니다. 기본값 환경변수를 사용합니다.")
	}

	Config = AppConfig{
		DB: DBConfig{
			Host:     getEnv("MYSQL_HOST", "localhost"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			User:     getEnv("MYSQL_USER", "root"),
			Password: getEnv("MYSQL_PASSWORD", ""),
			Database: getEnv("MYSQL_DATABASE", "test"),
		},
		Server: ServerConfig{
			Port: getEnv("BACKEND_PORT", "8080"),
		},
	}
}

func getEnv(key string, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	log.Printf("[INFO] 환경변수 %s 가 없어 기본값 %s 를 사용합니다.", key, fallback)
	return fallback
}
