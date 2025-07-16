# Go image
FROM golang:1.24.4 AS builder

WORKDIR /app

# Go mod fayllarni alohida kiritib kechikmasin
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Source kod
COPY . .

# Build
RUN go build -o main ./cmd/main.go

# Final stage — minimal image
FROM debian:bookworm-slim

WORKDIR /app

# tzdata (vaqt) qo‘shamiz
RUN apt-get update && apt-get install -y tzdata && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
