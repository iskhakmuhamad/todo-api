# Build stage
FROM golang:1.23.4-alpine AS builder

# Install git for fetching modules (if needed)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary, output as 'app' (adjust main package if needed)
RUN go build -o app ./cmd/main.go

# Final stage: minimal runtime image
FROM alpine:latest

# Install CA certificates for HTTPS if your app uses network calls
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/app .

# Expose the port your API listens on (change if needed)
EXPOSE 3000

# Run the binary
CMD ["./app"]
