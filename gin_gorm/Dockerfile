# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY gin_gorm/ ./gin_gorm/

# Build the application
WORKDIR /app/gin_gorm
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/gin_gorm/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]


