# Makefile for gitlab.com/fanligafc-group/fanligafc-backend

BINARY       := fanligafc-backend
CMD          := cmd/main.go
GOFILES      := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
CONFIG_PATH  ?= config/local.yaml
LOCAL_CONFIG_PATH := config/localsetup.yaml

.PHONY: all build run run-local run-config fmt lint test

all: fmt lint test build

# Build the binary into ./bin/
build:
	mkdir -p bin
	go build -o bin/$(BINARY) $(CMD)

# Run the app directly (uses default CONFIG_PATH logic in code)
run:
	CONFIG_PATH=config/local.yaml go run $(CMD)

# Run with CONFIG_PATH pointing at config/local.yaml using env var
run-local:
	go run $(CMD) --config=$(LOCAL_CONFIG_PATH)

# Run passing --config flag to your app
run-config:
	go run $(CMD) --config=$(CONFIG_PATH)

# Format code and fix imports
fmt:
	go fmt ./...
	goimports -w $(GOFILES)

# Lint with golangci-lint
lint:
	golangci-lint run ./...

# Run all tests with verbose output
test:
	go test ./... -v
