.PHONY: build up down dev restart logs clean

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

dev: down build up
	docker-compose logs -f

restart: down up

logs:
	docker-compose logs -f

clean: down
	docker-compose down -v
	docker-compose rm -f

rebuild: down build up

