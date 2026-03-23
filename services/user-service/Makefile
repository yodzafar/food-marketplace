include .env
export

SERVICE=user_service
MIGRATIONS_DIR=migrations
DB_DSN=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

.PHONY: migrate-up migrate-down migrate-status migrate-create build run wire test

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN) status"

migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

build:
	go build -o bin/(SERVICE) ./cmd/

run:
	go run ./cmd/

wire:
	wire ./wire/

test:
	go test ./..
