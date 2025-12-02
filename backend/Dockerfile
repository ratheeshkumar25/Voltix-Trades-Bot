# Multi-stage Dockerfile for Trading Bot Backend

# Stage 1: Builder
FROM golang:1.24-alpine AS builder

# Set build arguments
ARG GO_VERSION=1.24
ARG GIT_COMMIT=unknown
ARG BUILD_TIME=unknown

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-X 'main.Version=${GIT_COMMIT}' -X 'main.BuildTime=${BUILD_TIME}'" \
    -o /app/trading-bot ./cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /home/appuser

# Copy binary from builder
COPY --from=builder /app/trading-bot /usr/local/bin/trading-bot

# Copy configuration files if needed
# COPY --from=builder /app/config ./config

# Change ownership
RUN chown -R appuser:appuser /home/appuser

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD /usr/local/bin/trading-bot --health-check || exit 1

# Run the application
ENTRYPOINT ["/usr/local/bin/trading-bot"]
CMD []
