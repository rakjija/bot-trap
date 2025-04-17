package db

import (
	"gorm.io/gorm"
)

// 로그 한 건을 DB에 저장하는 함수
func SaveLog(db *gorm.DB, entry LogEntry) error {
	return db.Create(&entry).Error // 에러 반환 여부로 성공/실패 판단
}
