FROM golang:1.24.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o observer ./cmd/main.go

FROM ubuntu:24.04
WORKDIR /root
COPY --from=builder /app/observer .
COPY logs ./logs

EXPOSE 9000
CMD ["./observer"]