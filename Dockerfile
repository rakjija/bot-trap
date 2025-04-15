
# 1단계: 빌드 전용 환경
FROM golang:1.24.2 AS builder

# 작업 디렉토리 생성
WORKDIR /app

# go.mod, go.sum 복사 → 종속성 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 전체 소스 코드 복사
COPY . .

# main.go 빌드 (cmd/bot-trap/main.go)
RUN GOOS=linux GOARCH=amd64 go build -o bot-trap ./cmd/bot-trap

# 파일 확인 로그
RUN ls -al ./bot-trap

# 2단계: 실행용 경량 이미지
FROM alpine:latest

# certs 없으면 https 요청 안됨
RUN apk --no-cache add ca-certificates

# 빌드 결과 복사
WORKDIR /root/
COPY --from=builder /app/bot-trap .

# 실행 가능 여부 확인 로그
RUN ls -al /root/

# 컨테이너 실행 시 실행될 명령
CMD ["./bot-trap"]