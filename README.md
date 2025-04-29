# GoBoard ![GitHub release](https://img.shields.io/github/v/release/rakjija/go-board?style=flat-square) [![codecov](https://codecov.io/gh/rakjija/go-board/branch/main/graph/badge.svg)](https://codecov.io/gh/rakjija/go-board)

간단한 게시판과 사용자 행동 관찰 시스템입니다.
Kubernetes 클러스터에 배포하여, 프론트엔드, 백엔드, 옵저버를 통합 운영합니다.

---

## 아키텍처 및 사용 기술

| 구성 요소      | 사용 기술                                           |
| :------------- | :-------------------------------------------------- |
| **Frontend**   | React + Vite + nginx                                |
| **Backend**    | Go + Gin + Gorm + Swagger                           |
| **Monitoring** | Prometheus + Grafana + Loki + Promtail              |
| **Database**   | MySQL (Helm Chart 기반 설치)                        |
| **Infra**      | Docker, Helm, Kubernetes, Terraform, GitHub Actions |

---

## 디렉터리 구조

```text
├── backend/            # Go 기반 API 서버 (게시판 및 인증 기능)
├── frontend/           # React + Vite 기반 프론트엔드 (nginx 통합 제공)
├── mysql/              # MySQL 설정 파일 (docker-compose 연동용)
├── helm/               # Kubernetes 배포용 Helm Charts (backend, frontend, monitoring 등)
├── infra/              # Terraform 기반 AWS 인프라 코드 (EC2, S3 등 관리)
├── docker-compose.yml  # 로컬 개발용 Docker Compose 설정 (개발 및 테스트 목적)
└── README.md           # 프로젝트 설명 및 실행 가이드
```

---

## 실행 방법

### 1. 로컬 개발용 (Docker Compose 기반)

#### 1-1. Docker Compose 실행

```bash
docker-compose up --build
```

- frontend, backend, mysql, promtail, loki, grafana, prometheus 등이 함께 구동됩니다.

#### 1-2. 서비스 접속

| 서비스   | 주소                                            |
| -------- | ----------------------------------------------- |
| Frontend | http://localhost:3000                           |
| Backend  | http://localhost:8080                           |
| Grafana  | http://localhost:3001 (기본 ID/PW: admin/admin) |

---

### 2. 운영 배포용 (K8s + Helm 기반)

#### 2-1. Docker 이미지 빌드 및 Push

```bash
docker build -t <dockerhub-username>/goboard-backend:<version> ./backend
docker build -t <dockerhub-username>/goboard-frontend:<version> ./frontend

docker push <dockerhub-username>/goboard-backend:<version>
docker push <dockerhub-username>/goboard-frontend:<version>
```

#### 2-2. Helm Chart 수정

Helm 배포 전에 반드시 values.yaml 파일을 수정합니다.

- helm/backend/values.yaml
- helm/frontend/values.yaml

예시:

```yaml
# 수정 전
image:
  repository: rakjija/goboard-backend
  tag: latest

# 수정 후
image:
  repository: <dockerhub-username>/goboard-backend
  tag: <version>
```

#### 2-3. K8s 클러스터에 Helm Chart 설치

```bash
# 1. 네임스페이스 생성
kubectl create namespace goboard

# 2. 서비스별 Helm 설치
helm install frontend ./helm/frontend -n goboard
helm install backend ./helm/backend -n goboard

# 3. MySQL 설치 (bitnami/mysql chart 사용)
helm install mysql bitnami/mysql \
  --namespace goboard \
  --set auth.rootPassword=root \
  --set auth.database=goboard \
  --set auth.username=goboard-user \
  --set auth.password=goboard-pass \
  --set primary.service.type=ClusterIP

# 4. Monitoring Stack 설치 (grafana/loki-stack chart 사용)
helm install monitoring grafana/loki-stack \
  -n goboard \
  --set grafana.enabled=true
```

### 2-4. 서비스 포트포워딩 (로컬 접근용)

```bash
kubectl port-forward svc/frontend 3000:80 -n goboard
kubectl port-forward svc/backend 8080:80 -n goboard
kubectl port-forward svc/grafana 3001:3000 -n goboard
```

#### 2-5. 서비스 접속

| 서비스   | 주소                                            |
| -------- | ----------------------------------------------- |
| Frontend | http://localhost:3000                           |
| Backend  | http://localhost:8080                           |
| Grafana  | http://localhost:3001 (기본 ID/PW: admin/admin) |

---

### (선택) Terraform으로 인프라 구축

```bash
cd infra/terraform
terraform init
terraform apply
```
