package db

import (
	"testing"

	"gorm.io/driver/sqlite" // 테스트용 SQLite 드라이버
	"gorm.io/gorm"
)

func TestSaveLog(t *testing.T) {
	// ✅ 메모리 기반 SQLite DB 연결 (디스크 저장 X)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("테스트용 DB 연결 실패: %v", err)
	}

	// ✅ 테이블 생성 (LogEntry 스키마 기반)
	if err := db.AutoMigrate(&LogEntry{}); err != nil {
		t.Fatalf("테이블 마이그레이션 실패: %v", err)
	}

	// ✅ 테스트용 데이터 준비
	entry := LogEntry{
		IP:      "127.0.0.1",
		Path:    "/test",
		Message: "테스트 메시지",
	}

	// ✅ 실제 저장 함수 호출
	err = SaveLog(db, entry)
	if err != nil {
		t.Errorf("로그 저장 실패: %v", err)
	}

	// ✅ 저장 결과 검증
	var result LogEntry
	tx := db.First(&result)
	if tx.Error != nil {
		t.Errorf("저장된 로그 조회 실패: %v", tx.Error)
	}
	if result.IP != entry.IP || result.Path != entry.Path || result.Message != entry.Message {
		t.Errorf("저장된 값이 다릅니다: %+v", result)
	}
}
