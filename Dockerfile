# Stage 1: Build
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Run
FROM debian:bullseye-slim

# Set timezone (optional)
ENV TZ=Asia/Jakarta

# Set working directory
WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
