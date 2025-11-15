FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy redis basic_operations project
COPY redis/basic_operations/ ./redis/basic_operations/

WORKDIR /app/redis/basic_operations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/redis/basic_operations/main .

CMD ["./main"]

