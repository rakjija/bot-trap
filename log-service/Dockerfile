# Stage 1: Build
FROM --platform=linux/amd64 golang:1.24.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

# CGO를 끄고 완전 정적 빌드
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o log-service .

# Stage 2: Run
FROM scratch
COPY --from=builder /app/log-service /log-service
CMD ["/log-service"]