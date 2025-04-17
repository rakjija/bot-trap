# BotTrap 🛡️ ![GitHub release](https://img.shields.io/github/v/release/rakjija/bot-trap?style=flat-square)

로그 기반 이상행동 탐지를 위한 백엔드 관찰 시스템  
수집된 요청 로그를 저장하고, 실시간으로 모니터링하며  
DevOps 자동화 파이프라인까지 구성한 프로젝트입니다.

---

## 🚀 프로젝트 개요

- 사용자 요청 로그를 수집하여 PostgreSQL에 저장
- Prometheus + Grafana로 실시간 이상 행위 관찰 및 시각화
- GitHub Actions를 통해 CI/CD 파이프라인 구성 및 Docker Hub 배포 자동화

---

## 💪 사용 기술 스택

| 영역     | 기술                              |
|----------|-----------------------------------|
| 언어     | Go (1.24.2)                       |
| DB       | PostgreSQL                        |
| 관찰     | Prometheus, Grafana               |
| 컨테이너 | Docker, docker-compose            |
| 자동화   | GitHub Actions (CI + CD + Release) |
| 배포     | Docker Hub (`rakjija/bottrap`)    |

---

## 📁 디렉토리 구조 요약

```bash
├── cmd/                  # main.go 위치
├── internal/             # 서비스 로직
│   ├── api/              # API 핸들러, 테스트 코드 포함
│   ├── db/               # DB 모델 정의
│   └── metrics/          # Prometheus 메트릭 등록 및 미들웨어
├── grafana/              # Grafana 대시보드 자동 구성
│   ├── dashboards/       # JSON 대시보드 구성
│   └── provisioning/     # Grafana 프로비저닝 설정
├── Dockerfile
├── docker-compose.yml
└── .github/workflows/    # CI/CD GitHub Actions
    ├── ci.yml
    ├── release.yml
    └── cd.yml
```

---

## 🧱 아키텍처

```plaintext
User
 ↓
[log-service] ────▶ PostgreSQL
     │
     └────────▶ /metrics ────▶ Prometheus ─▶ Grafana
```

---

## ✅ 기능 요약

- **GET /healthz**: 헬스 체크 API
- **POST /logs**: JSON 형식 로그 저장 API
- **GET /metrics**: Prometheus 포맷 메트릭 제공

### 📌 주요 메트릭 목록
| 메트릭 이름                   | 설명                  |
|------------------------------|-----------------------|
| `log_save_total`             | 저장 성공 횟수       |
| `log_error_total`            | 저장 실패 횟수       |
| `log_suspicious_total`       | 의심되는 로그 횟수   |
| `http_request_duration_seconds` | 요청 응답 시간 측정 |

---

## 📦 실행 방법

```bash
git clone https://github.com/rakjija/bot-trap.git
cd bot-trap
docker-compose up --build
```

- `http://localhost:8080/logs`: 로그 저장 API
- `http://localhost:9090`: Prometheus UI
- `http://localhost:3000`: Grafana 대시보드 (ID/PW: admin / admin)

---

## 📊 Grafana 대시보드

> 대시보드는 JSON으로 자동 구성되며, 두 개의 주요 대시보드로 나뉩니다:

- 📦 **Log Dashboard**: 로그 저장 / 에러 수
- 📈 **Metrics Dashboard**: 의심 로그 수 / 평균 응답 시간

> 필요 시 `grafana/dashboards/`에서 JSON 수정 및 반영 가능

---

## ⚙️ CI/CD 자동화 (GitHub Actions)

| 파일명         | 기능                                 | 트리거        |
|----------------|--------------------------------------|---------------|
| `ci.yml`       | go test + Docker 이미지 빌드 (테스트용) | push, PR      |
| `release.yml`  | 릴리스 생성 + Docker Hub 이미지 배포 | git tag (`v*`) |
| `cd.yml`       | 수동 배포 (예: SSH 기반 운영 서버 전개) | Actions 버튼 실행 |

---

## 🔐 Docker Hub 자동 푸시 예시

```bash
git tag v1.0.0
git push origin v1.0.0
```

태그 푸시 완료 시 다음과 같이 Docker 이미지가 자동으로 올라갑니다:

- `rakjija/bottrap:v1.0.0`
- `rakjija/bottrap:latest`

또한 `GitHub Releases` 탭에도 자동 릴리스가 생성됩니다 🥳

