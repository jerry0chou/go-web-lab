FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy cache_patterns project
COPY redis/cache_patterns/ ./redis/cache_patterns/

WORKDIR /app/redis/cache_patterns
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/redis/cache_patterns/main .

CMD ["./main"]

