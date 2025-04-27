# GoBoard ![GitHub release](https://img.shields.io/github/v/release/rakjija/go-board?style=flat-square) [![codecov](https://codecov.io/gh/rakjija/go-board/branch/main/graph/badge.svg)](https://codecov.io/gh/rakjija/go-board)

간단한 게시판과 사용자 행동 관찰 시스템입니다.
Kubernetes 클러스터에 배포하여, 프론트엔드, 백엔드, 옵저버를 통합 운영합니다.

---

## 아키텍처 및 사용 기술

| 구성 요소    | 사용 기술                                                |
| :----------- | :------------------------------------------------------- |
| **Frontend** | React + Vite + nginx                                     |
| **Backend**  | Go + Gin + Gorm + Swagger                                |
| **Observer** | Go + Gin, ELK Stack                                      |
| **Database** | MySQL (Helm Chart 기반 설치)                             |
| **Infra**    | Docker, Helm, Kubernetes(k3s), Terraform, GitHub Actions |

---

## 상세 작업 내역

### 0. 공통 작업

- Dockerfile 작성 및 Docker 이미지 빌드
- Helm Chart 작성 → Kubernetes에 배포
- GitHub Actions로 CI 구성
  - Backend 테스트 및 커버리지 측정 (Codecov 연동)
  - Docker Build 테스트
  - Slack 알림 연동 (Build/Release 결과 실시간 알림)

### 1. Frontend (React + Vite + nginx)

- React + Vite로 프로젝트 구성
- nginx를 이용한 정적 파일 제공

### 2. Backend (Go + Gin + Gorm)

- Gin 기반 REST API 서버 개발
- JWT 인증 및 MySQL 연동 기능 구현
- Swagger 문서화 자동화 (`Dockerfile` 내 `swag init` 적용)
- 커스텀 Validator 추가 (예: 이메일, 비밀번호 형식 등)
- 주요 API 핸들러 테스트 코드 작성 (유닛 테스트 기반)
- Makefile을 통한 프로젝트 자동화 환경 구성
  - 빌드, 테스트, 문서화 명령어 일원화

### 3. Observer (Go + ELK Stack)

- 사용자 행동 수집을 위한 Observer 서비스 개발
- 로그를 ELK(Elasticsearch, Logstash, Kibana)로 전송

### 4. Database (MySQL)

- Helm 공식 Chart를 활용하여 MySQL 설치
- 사용자, 게시판 데이터 저장용 DB 구성
- Kubernetes Persistent Volume Claim(PVC) 사용 (선택사항)

### 5. Infra & DevOps

- Terraform을 사용하여 AWS EC2 인스턴스 생성
- EC2 인스턴스에 k3s 설치하여 경량 Kubernetes 클러스터 구축
- Helm을 통한 서비스 배포 자동화
- GitHub Actions를 통한 자동 Build & Release 파이프라인 구축
- 포트포워딩을 통한 로컬 테스트 환경 구성

---

## 디렉터리 구조

```text
├── backend/            # Go 기반 API 서버
├── frontend/           # React 기반 프론트엔드
├── observer/           # 사용자 행동 기록 Observer (Go + ELK)
├── mysql/              # MySQL 설정 파일 (docker-compose 연동용)
├── infra/              # Terraform 기반 AWS 인프라 코드
├── k8s/                # 수동 Kubernetes 배포용 YAML 파일 (초기 작업)
├── helm/               # Kubernetes 배포용 Helm Charts
├── docker-compose.yml  # 로컬 개발용 Docker Compose 파일
└── README.md           # 프로젝트 설명 문서
```

---

## 실행 방법

### 1. Docker 이미지 빌드 및 Push

```bash
docker build -t <dockerhub-username>/goboard-backend:latest ./backend
docker build -t <dockerhub-username>/goboard-frontend:latest ./frontend
docker build -t <dockerhub-username>/goboard-observer:latest ./observer

docker push <dockerhub-username>/goboard-backend:latest
docker push <dockerhub-username>/goboard-frontend:latest
docker push <dockerhub-username>/goboard-observer:latest
```

### 2. Kubernetes 클러스터에 Helm Chart 설치

```bash
helm upgrade --install go-board-backend helm/backend -n goboard
helm upgrade --install go-board-frontend helm/frontend -n goboard
helm upgrade --install go-board-observer helm/observer -n goboard
```

### 3. 서비스 포트포워딩

```bash
kubectl port-forward svc/go-board-frontend 3000:80 -n goboard
kubectl port-forward svc/go-board-backend 8080:80 -n goboard
kubectl port-forward svc/go-board-observer 9000:9000 -n goboard
```

### (선택) Terraform으로 인프라 구축

```bash
cd infra/terraform
terraform init
terraform apply
```
