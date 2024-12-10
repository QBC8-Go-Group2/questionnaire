# Build stage
FROM golang:1.23-alpine AS builder

# Install git and build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
# Copy config files and keys
COPY --from=builder /app/config.json .
COPY --from=builder /app/Private_key.pem .
COPY --from=builder /app/Public_key.pem .

# Create uploads directory
RUN mkdir -p pkg/data/uploads

# Expose port
EXPOSE 3000

# Command to run the application
CMD ["./main"]