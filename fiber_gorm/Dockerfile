FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY fiber_gorm/ ./fiber_gorm/

WORKDIR /app/fiber_gorm
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/fiber_gorm/main .

EXPOSE 8080

CMD ["./main"]

