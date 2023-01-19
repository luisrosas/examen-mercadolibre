.PHONY: build test coverage

PACKAGES := $(shell go list ./... )

build:
	@go build -o ./build/ cmd/api/main.go

tests:
	@go test -count=1 -cover -timeout 60s $(PACKAGES)

coverage:
	@rm -rf ./coverage
	@mkdir ./coverage
	@go test -count=1 -cover -coverprofile ./coverage/coverage.out -timeout 60s $(PACKAGES)
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
