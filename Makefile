BINARY_NAME=govid

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

.PHONY: all build clean test coverage lint run migrate help

all: test build

build: ## Build the application
	@echo "Building ${BINARY_NAME}..."
	@go build ${LDFLAGS} -o $(GOBIN)/${BINARY_NAME} ./cmd/govid/main.go

run: build ## Run the application
	@./$(GOBIN)/${BINARY_NAME}

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./internal/...

lint: ## Run golangci-lint
	@echo "Running linter..."
	@golangci-lint run ./...

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

tidy: ## Tidy and verify go modules
	@echo "Tidying modules..."
	@go mod tidy
	@go mod verify

migrate: ## Run database migrations
	@echo "Running migrations..."
	@go run migrate/migrate.go

dev: ## Run with hot reload
	@air
