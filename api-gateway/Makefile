.PHONY: build run wire test

run:
	go run ./cmd/

build:
	go build -o bin/api-gateway ./cmd/

wire:
	wire ./wire/

test:
	go test ./..