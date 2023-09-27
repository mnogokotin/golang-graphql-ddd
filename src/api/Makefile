include .env

help: ### display help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

compose-up: ### run docker compose
	docker compose --env-file=../../docker/.env up -d --build
.PHONY: compose-up

compose-down: ### down docker compose
	docker compose down --remove-orphans
.PHONY: compose-down

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

mock: ### run mockgen
	mockgen -source ./internal/usecase/interfaces.go -package usecase_test > ./internal/usecase/mocks_test.go
.PHONY: mock

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	go run ./cmd/migrate/main.go
.PHONY: migrate-up

fixtures-load: ### load fixtures
	go run ./cmd/fixtures_loader/main.go
.PHONY: fixtures-load

fixtures-dump: ### dump fixtures
	go run ./cmd/fixtures_dumper/main.go
.PHONY: fixtures-dump
