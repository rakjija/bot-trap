package config

import (
	"database/sql"
	"fmt"
	"log"
	"log-service/models"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var db *sql.DB
	var err error

	// 💡 최대 10회 재시도, 2초 간격
	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("[시도 %d] DB 연결 실패: %v", i, err)
		} else if err = db.Ping(); err != nil {
			log.Printf("[시도 %d] DB ping 실패: %v", i, err)
		} else {
			break // 성공
		}

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("DB 연결 최종 실패: %v", err)
	}

	models.DB = db
	log.Println("✅ DB 연결 성공")

	// 테이블 생성 쿼리
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

	log.Println("✅ logs 테이블 확인 완료 (없으면 생성됨)")
}
