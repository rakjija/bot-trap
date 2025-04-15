# BotTrap 🛡️

로그 기반 이상행동 탐지를 위한 백역된 관찰 시스템

## 🚀 프로젝트 개요

- 사용자의 요청 로그를 수집 및 저장
- Prometheus + Grafana로 관찰 및 시각화
- GitHub Actions를 통한 CI 파이프라인 구성

---

## 💪 사용 기술 스택

| 영역 | 기술 |
|------|------|
| 언어 | Go (1.24.2) |
| DB | PostgreSQL |
| 관찰 | Prometheus, Grafana |
| 컨테이너 | Docker, docker-compose |
| 자동화 | GitHub Actions |

---

## 📁 디렉토리 구조 요약

```
├── cmd/
├── internal/
│   ├── api/
│   ├── db/
├── grafana/
│   ├── dashboards/
│   └── provisioning/
├── Dockerfile
├── docker-compose.yml
└── .github/workflows/ci.yml
```

---

## 🧱 아키텍처

```
User
 ↓
[log-service] ────▶ PostgreSQL
     │
     └────────▶ /metrics ────▶ Prometheus ─▶ Grafana
```

---

## ✅ 기능 요약

- `/logs`: JSON 형식 로그 저장 API
- `/healthz`: 헬스 체크 API
- `/metrics`: Prometheus 포맷 메트릭 제공
- `log_save_total`: 저장 성공 횟수
- `log_error_total`: 저장 실패 횟수

---

## 📦 실행 방법

```bash
git clone https://github.com/rakjija/bot-trap.git
cd bot-trap
docker-compose up --build
```

- `localhost:8080/logs`: 로그 API 테스트
- `localhost:9090`: Prometheus UI
- `localhost:3000`: Grafana 대시보드 (ID/PW: admin/admin)

---

## 📊 Grafana 대시보드

자동 구성된 패널 예시:
- Total Logs Saved
- Log Save Errors

---

## ⚙️ CI 자동화 (GitHub Actions)

- `push` 또는 `PR` 발생 시 자동으로 `go build ./cmd/bot-trap` 실행
