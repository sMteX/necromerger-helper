# --- Stage 1: Build Backend ---
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server main.go

# --- Stage 2: Final Image ---
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/server .

RUN adduser -D -u 1000 necro
USER necro

ENV PORT=8080
EXPOSE 8080

CMD ["./server"]