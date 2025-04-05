DOCKER_COMPOSE_FILE := ./docker/docker-compose.yml

MIGRATION_DB_URL := $(shell awk -F '=' '/^MIGRATION_DB_URL/ {print substr($$0, index($$0,$$2))}' ./docker/.env)

.PHONY: up down ps logs build restart migrate-up migrate-down migrate-drop

up:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

ps:
	docker compose -f $(DOCKER_COMPOSE_FILE) ps

logs:
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f

build:
	docker compose -f $(DOCKER_COMPOSE_FILE) build

restart: down up

migrate-up:
	docker compose -f $(DOCKER_COMPOSE_FILE) exec migrations \
	  migrate -path=/migrations \
	          -database="$(MIGRATION_DB_URL)" \
	          up

migrate-down:
	docker compose -f $(DOCKER_COMPOSE_FILE) exec migrations \
	  migrate -path=/migrations \
	          -database="$(MIGRATION_DB_URL)" \
	          down 1

migrate-drop:
	docker compose -f $(DOCKER_COMPOSE_FILE) exec migrations \
	  migrate -path=/migrations \
	          -database="$(MIGRATION_DB_URL)" \
	          drop -f

migrate-create:
ifndef name
	$(error You must specify a migration name, e.g.: make migrate-create name=create_users_table)
endif
	docker compose -f $(DOCKER_COMPOSE_FILE) exec migrations \
	  migrate create -ext sql -dir /migrations -seq $(name)