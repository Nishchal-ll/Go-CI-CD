# ==========================================
# Stage 1: Build the Go binary
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first (optimizes Docker layer caching)
COPY go.mod ./
# Download dependencies (if any)
RUN go mod download

# Copy source code
COPY . .

# Compile static binary (CGO_ENABLED=0 creates a portable standalone binary)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server .

# ==========================================
# Stage 2: Final Minimal Runtime Image
# ==========================================
FROM alpine:3.20

WORKDIR /app

# Add ca-certificates for secure HTTPS requests & curl for health checks
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose port 8080
EXPOSE 8080

# Run the Go binary
CMD ["/app/server"]
