package api

import (
	"testing" // Go의 기본 테스트 패키지
)

func TestParseLogJSON(t *testing.T) {
	// ✅ 케이스 1: 정상적인 JSON 입력
	t.Run("정상적인 JSON 입력", func(t *testing.T) {
		// 테스트용 JSON 문자열 (정상 케이스)
		jsonData := []byte(`{"ip": "1.1.1.1", "path": "/", "message": "test"}`)

		// 파싱 함수 호출
		entry, err := ParseLogJSON(jsonData)

		// 에러가 발생하지 않아야 함
		if err != nil {
			t.Fatalf("예상치 못한 에러 발생: %v", err)
		}

		// 파싱된 내용이 예상한 값과 일치하는지 확인
		if entry.IP != "1.1.1.1" || entry.Path != "/" || entry.Message != "test" {
			t.Errorf("파싱된 데이터가 예상과 다릅니다: %+v", entry)
		}
	})

	// ✅ 케이스 2: 필드 누락된 JSON 입력
	t.Run("필드 누락된 JSON 입력", func(t *testing.T) {
		// path, message 필드가 없는 JSON
		jsonData := []byte(`{"ip": "1.1.1.1"}`)

		entry, err := ParseLogJSON(jsonData)

		// 에러는 없어야 함 (누락된 필드는 zero value 처리됨)
		if err != nil {
			t.Fatalf("에러 없어야 함: %v", err)
		}

		// 누락된 필드는 빈 문자열("")이어야 함
		if entry.Path != "" || entry.Message != "" {
			t.Errorf("누락된 필드는 빈 문자열이어야 합니다: %+v", entry)
		}
	})

	// ✅ 케이스 3: 문법이 틀린 JSON 입력
	t.Run("잘못된 JSON 입력", func(t *testing.T) {
		// 중괄호가 닫히지 않은 잘못된 JSON
		jsonData := []byte(`{`)

		_, err := ParseLogJSON(jsonData)

		// 반드시 에러가 발생해야 함
		if err == nil {
			t.Errorf("파싱 에러가 발생해야 합니다")
		}
	})
}
