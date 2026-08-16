# Этап сборки
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o qii .

# Этап запуска
FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/qii .

# Чтобы не было проблем с сертификатами Telegram API
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

CMD ["./qii"]
