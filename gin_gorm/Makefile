.PHONY: build up down dev restart logs clean

# Build Docker images
build:
	docker-compose build

# Start services in detached mode
up:
	docker-compose up -d

# Stop and remove containers
down:
	docker-compose down

# Development mode: rebuild and start with logs
dev: down build up
	docker-compose logs -f

# Restart services
restart: down up

# View logs
logs:
	docker-compose logs -f

# Stop and remove containers, volumes, and images
clean: down
	docker-compose down -v
	docker-compose rm -f

# Rebuild and restart (useful for code changes)
rebuild: down build up


