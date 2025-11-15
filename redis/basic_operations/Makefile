.PHONY: up down build run clean cli

# Start Redis and Go app with Docker Compose
up:
	docker-compose up --build -d

# Stop and remove containers
down:
	docker-compose down

# Build Docker images
build:
	docker-compose build

# Run the application (assumes Redis is running)
run:
	go run main.go

# Clean up Docker volumes
clean:
	docker-compose down -v

# Connect to Redis CLI
cli:
	docker exec -it redis_server redis-cli

# Run locally (requires Redis running on localhost:6379)
run-local:
	@echo "Make sure Redis is running on localhost:6379"
	@sed -i.bak 's/redis:6379/localhost:6379/g' main.go || sed -i '' 's/redis:6379/localhost:6379/g' main.go
	go run main.go
	@sed -i.bak 's/localhost:6379/redis:6379/g' main.go || sed -i '' 's/localhost:6379/redis:6379/g' main.go
	@rm -f main.go.bak

