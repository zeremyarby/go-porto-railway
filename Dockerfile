# Stage 1: Build
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy go modules dulu biar cache efisien
COPY go.mod ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary (main ada di cmd/web/main.go)
RUN go build -o portfolio ./cmd/web

# Stage 2: Runtime
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/portfolio .

# Copy static assets & templates
COPY --from=builder /app/static ./static

EXPOSE 8080
CMD ["./portfolio"]
