.PHONY: mod lint start-deps

mod:
	go mod tidy

lint:
	golangci-lint run

start-deps:
	cd ./testing && docker-compose up -d && sleep 5

build:
	go build ./cmd/backup