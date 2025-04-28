FROM golang:1.24.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go run github.com/swaggo/swag/cmd/swag@latest init -g ./cmd/main.go -o ./docs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

FROM ubuntu:24.04
RUN apt-get update
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]