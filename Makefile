.PHONY: install build test coverage local-up local-down server-logs database-logs

PACKAGES := $(shell go list ./... )

install:
	@go mod download

build:
	@go build -o ./build/ cmd/api/main.go

tests:
	@go test -count=1 -cover -timeout 60s $(PACKAGES)

coverage:
	@rm -rf ./coverage
	@mkdir ./coverage
	@go test -count=1 -cover -coverprofile ./coverage/coverage.out -timeout 60s $(PACKAGES)
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

local-up:
	@docker-compose up -d

local-down:
	@docker-compose down

server-logs:
	@docker-compose logs server_go

database-logs:
	@docker-compose logs mysql
