# log-service

로그 기록용 마이크로서비스입니다.  
클라이언트의 IP, User-Agent, 요청 경로, 시간을 PostgreSQL에 저장합니다.

## 🔧 실행 방법

1. `.env` 수정

## 실행 명령어

Docker 개발 환경 실행을 편리하게 하기 위해 `Makefile`을 제공합니다.

### 주요 명령어

| 명령어 | 설명 |
|--------|------|
| `make up` | 전체 컨테이너 빌드 및 실행 (`docker-compose up --build -d`) |
| `make down` | 컨테이너 종료 및 정리 |
| `make restart` | 컨테이너 재시작 |
| `make logs` | 로그 스트리밍 확인 (`docker-compose logs -f`) |
| `make clean` | 모든 컨테이너 + 볼륨 삭제 |
| `make build` | log-service 이미지 단독 빌드 |