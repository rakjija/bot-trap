# log-service

로그 기록용 마이크로서비스입니다.  
클라이언트의 IP, User-Agent, 요청 경로, 시간을 PostgreSQL에 저장합니다.

## 🔧 실행 방법

1. `.env` 파일 생성:
```
DB_HOST=<localhost>
DB_PORT=5432
DB_USER=bottrap_user
DB_PASSWORD=<yourpassword>
DB_NAME=bottrap_db
```