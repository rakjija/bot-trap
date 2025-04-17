# BotTrap 🛡️

로그 기반 이상행동 탐지를 위한 백엔드 관찰 시스템
수집된 요청 로그를 저장하고, 실시간으로 모니터링하며
DevOps 자동화 파이프라인까지 구성한 프로젝트입니다.

## 🚀 프로젝트 개요

- 사용자 요청 로그를 수집하여 PostgreSQL에 저장
- Prometheus + Grafana로 실시간 이상 행위 관찰 및 시각화
- GitHub Actions를 통해 CI/CD 파이프라인 구성 및 Docker Hub 배포 자동화

---

## 💪 사용 기술 스택

| 영역 | 기술 |
|------|------|
| 언어 | Go (1.24.2) |
| DB | PostgreSQL |
| 관찰 | Prometheus, Grafana |
| 컨테이너 | Docker, docker-compose |
| 자동화 | GitHub Actions (CI + CD) |
| 배포 | Docker Hub (rakjija/bottrap) |

---

## 📁 디렉토리 구조 요약

```
├── cmd/                  # main.go 위치
├── internal/             # 서비스 로직
│   ├── api/              # API 핸들러, 테스트 코드 포함
│   ├── db/               # DB 모델 정의
├── grafana/              # Grafana 대시보드 자동 구성
│   ├── dashboards/
│   └── provisioning/
├── Dockerfile
├── docker-compose.yml
└── .github/workflows/    # CI/CD GitHub Actions
    ├── ci.yml
    ├── release.yml
    └── cd.yml
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

- GET /healthz: 헬스 체크 API
- POST /logs: JSON 형식 로그 저장 API
- GET /metrics: Prometheus 포맷 메트릭 제공
- 주요 메트릭:
    - log_save_total: 저장 성공 횟수
    - log_error_total: 저장 실패 횟수

---

## 📦 실행 방법

```bash
git clone https://github.com/rakjija/bot-trap.git
cd bot-trap
docker-compose up --build
```

- `localhost:8080/logs`: 로그 저장 API
- `localhost:9090`: Prometheus UI
- `localhost:3000`: Grafana 대시보드 (ID/PW: admin/admin)

---

## 📊 Grafana 대시보드

> 자동으로 패널이 구성되도록 하였습니다.

- Total Logs Saved
- Log Save Errors

---

## ⚙️ CI/CD 자동화 (GitHub Actions)

| 용도         | 기능                                | 트리거        |
|--------------|-------------------------------------|---------------|
| `ci.yml`     | go test + docker build (test only) | push, PR      |
| `release.yml`| docker build & push to Docker Hub  | tag (`v*`)    |
| `cd.yml`     | 수동 배포 (예: SSH 배포 스크립트)   | Actions 버튼 실행 |

---

### 🔐 Docker Hub 자동 푸시 예시

```bash
git tag v1.0.0
git push origin v1.0.0
```
완료 시 Docker Hub에 다음과 같은 이미지가 자동 푸시됩니다:

•	rakjija/bottrap:v1.0.0
•	rakjija/bottrap:latest
