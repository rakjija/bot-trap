# BotTrap 🛡️ ![GitHub release](https://img.shields.io/github/v/release/rakjija/bot-trap?style=flat-square) [![codecov](https://codecov.io/gh/rakjija/bot-trap/branch/main/graph/badge.svg)](https://codecov.io/gh/rakjija/bot-trap)

로그 기반 이상 행동 탐지를 위한 백엔드 관찰 시스템  
애플리케이션 로그를 Loki로 수집하고, 
실시간으로 관찰 및 시각화하며,
DevOps 자동화 파이프라인까지 구성한 프로젝트입니다.

---

## 🚀 프로젝트 개요

- 애플리케이션 로그를 수집하여 Loki에 저장
- Prometheus + Grafana로 메트릭 및 로그 기반 이상 행위 관찰 및 시각화
- GitHub Actions를 통해 CI/CD 파이프라인 구성 및 Docker Hub 배포 자동화

---

## 💪 사용 기술 스택

| 영역     | 기술 스택                                 |
|----------|---------------------------------------|
| 언어     | Go (1.24.2)                           |
| 로그 수집 | Loki, Promtail                        |
| 메트릭 수집 | Prometheus                          |
| 시각화   | Grafana                                |
| 컨테이너 | Docker, docker-compose                |
| 자동화   | GitHub Actions (CI + CD + Release)    |
| 배포     | Docker Hub (`rakjija/bottrap`)        |

---

## 📁 디렉토리 구조 요약

```bash
├── cmd/                     # main.go 위치
├── internal/                # 서비스 로직
│   ├── api/                 # API 핸들러, 테스트 코드 포함
│   ├── metrics/             # Prometheus 메트릭 등록 및 미들웨어
│   └── model/               # 요청 구조체 정의
├── grafana/                 # Grafana 대시보드 자동 구성
│   ├── dashboards/          # JSON 대시보드 구성
│   └── provisioning/        # Grafana 프로비저닝 설정
├── prometheus.yml           # Prometheus 수집 설정
├── promtail-config.yaml     # Promtail 로그 수집 설정
├── Dockerfile
├── docker-compose.yml
└── .github/workflows/       # CI/CD GitHub Actions
    ├── ci.yml
    ├── release.yml
    └── cd.yml
```

---

## 🧱 아키텍처

```plaintext
User
 ↓
[log-service] ───▶ stdout ───▶ Promtail ───▶ Loki ┐
       │                                          │
       └────▶ /metrics ─────▶ Prometheus ─────┬───┘
                                              ↓
                                           Grafana
```
- **stdout 로그**를 Promtail이 수집 → Loki 저장  
- `/metrics`는 Prometheus가 scrape
- 둘 다 **Grafana**에서 시각화

---

## ✅ 기능 요약

- **GET /healthz**: 헬스 체크 API
- **POST /logs**: JSON 형식 로그 수신 → 로그는 stdout에 출력되어 Loki로 수집됨
- **GET /metrics**: Prometheus 포맷 메트릭 제공

### 📌 주요 메트릭 목록
| 메트릭 이름                     | 설명                           |
|--------------------------------|--------------------------------|
| `log_save_total`               | 로그 수신 성공 횟수            |
| `log_error_total`              | JSON 파싱 또는 처리 실패 횟수  |
| `log_suspicious_total`         | 의심되는 로그 횟수 (`bot`, `sql`, `/admin` 포함) |
| `http_request_duration_seconds`| 요청 처리 시간 (초 단위)       |

---

## 📦 실행 방법

```bash
git clone https://github.com/rakjija/bot-trap.git
cd bot-trap
docker-compose up --build
```

- http://localhost:8080/logs : 로그 수신 API (POST 요청으로 테스트)
- http://localhost:9090 : Prometheus UI (메트릭 확인)
- http://localhost:3000 : Grafana 대시보드 (ID/PW: admin / admin)

### 🧪 테스트 예시

```bash
curl -X POST http://localhost:8080/logs \
  -H "Content-Type: application/json" \
  -d '{"ip": "1.2.3.4", "path": "/admin", "message": "bot traffic detected"}'
```

---

## 📊 Grafana 대시보드

> 대시보드는 JSON으로 자동 구성되며, 두 개의 주요 대시보드로 구분됩니다:

- 📦 **Log Dashboard (Loki 기반)**: 수신된 로그 수 / 로그 처리 에러
- 📈 **Metrics Dashboard (Prometheus 기반)**: 의심 로그 수 / 요청 응답 시간

> 필요 시 `grafana/dashboards/` 내 JSON 파일을 수정하여 패널 구성 변경 가능

---

## ⚙️ CI/CD 자동화 (GitHub Actions)

| 파일명         | 기능                                                   | 트리거               |
|----------------|--------------------------------------------------------|----------------------|
| `ci.yml`       | 테스트 실행 (`go test`) + Docker 이미지 빌드(유효성 검증 테스트) + Slack 알림 | push, PR             |
| `release.yml`  | GitHub 릴리스 생성 + Docker Hub 이미지 자동 푸시 + Slack 알림 | git tag (`v*`)        |
| `cd.yml`       | 수동 배포 (예: 운영 서버에 SSH 접속 후 배포 스크립트 실행) | Actions > Run workflow |

---

## 🔐 Docker Hub 자동 푸시 예시

```bash
git tag v1.0.0
git push origin v1.0.0
```

태그 푸시 완료 시 다음과 같이 Docker 이미지가 자동으로 올라갑니다:

- `rakjija/bottrap:v1.0.0`
- `rakjija/bottrap:latest`

또한 `GitHub Releases` 탭에도 자동 릴리스가 생성됩니다

