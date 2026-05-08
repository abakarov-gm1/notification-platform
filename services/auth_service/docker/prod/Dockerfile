# ---------- build stage ----------
FROM golang:1.25-bookworm AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./cmd/main.go


# ---------- runtime stage ----------
FROM debian:bookworm-slim

WORKDIR /app

# копируем только бинарник
COPY --from=builder /build/auth .

# запуск
CMD ["./auth"]