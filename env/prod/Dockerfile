# -------- build stage --------
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Зависимости
COPY go.mod go.sum ./
RUN go mod download

# Исходники
COPY . .

# Сборка
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o ./build/api ./cmd/api


# -------- runtime stage --------
FROM alpine:3.20

WORKDIR /app

# Бинарь
COPY --from=builder /app/build/api /app/api

# Конфиги
COPY config ./config

# Не root (хорошая практика)
RUN chmod +x /app/api &&  adduser -D appuser && chown appuser:appuser -R /app

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/api"]
