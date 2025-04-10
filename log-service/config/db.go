package config

import (
	"database/sql"
	"fmt"
	"log"
	"log-service/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB ping 실패:", err)
	}

	models.DB = db
	log.Println("DB 연결 성공")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		ip_address VARCHAR(100),
		user_agent TEXT,
		path VARCHAR(255),
		timestamp TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("테이블 생성 실패:", err)
	}

	log.Println("logs 테이블 확인 완료 (없으면 생성됨)")
}
