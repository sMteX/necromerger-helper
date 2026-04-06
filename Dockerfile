# --- Stage 1: Build Frontend ---
FROM node:22-alpine AS frontend-builder

WORKDIR /app/ui

# Copy package.json and package-lock.json
COPY ui/package*.json ./

# Install dependencies
RUN npm install

# Copy frontend source
COPY ui/ ./

# Build the frontend (Vite)
RUN npm run build

# --- Stage 2: Build Backend ---
FROM golang:1.26-alpine AS backend-builder

WORKDIR /app

# Install necessary build tools if needed
RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the backend source
COPY . .

# Copy the built frontend assets from the previous stage
# This is crucial because of the //go:embed ui/dist
RUN rm -rf ui/dist
COPY --from=frontend-builder /app/ui/dist ./ui/dist

# Build the Go server
RUN go build -o server main.go

# --- Stage 3: Final Image ---
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the binary from the backend-builder stage
COPY --from=backend-builder /app/server .

# Create a non-root user for security
RUN adduser -D -u 1000 necro
USER necro

# Set the port the server listens on
ENV PORT=8080
EXPOSE 8080

# Command to run the application
CMD ["./server"]
