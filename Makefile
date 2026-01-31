.PHONY: help migrateup migratedown start build clean run migrate-create

# Variables
DB_CONN ?= postgresql://user:password@database:5432/postgres
MIGRATIONS_PATH ?= migrations
BINARY_NAME ?= main

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/http
	@echo "Build complete: $(BINARY_NAME)"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@echo "Clean complete"

run: ## Build and run the application
	@echo "Building and running..."
	@go run ./cmd/http

start: build ## Build and start the application
	@echo "Starting $(BINARY_NAME)..."
	@./$(BINARY_NAME)

migrateup: ## Run database migrations up
	@echo "Running migrations up..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_CONN)" up
	@echo "Migrations up complete"

migratedown: ## Run database migrations down
	@echo "Running migrations down..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_CONN)" down 1
	@echo "Migrations down complete"

migratedown-all: ## Run database migrations down (all)
	@echo "Running all down migrations..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_CONN)" down -all
	@echo "All migrations down complete"

migrate-force: ## Force a migration version (usage: make migrate-force VERSION=1)
	@echo "Forcing migration to version $(VERSION)..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_CONN)" force $(VERSION)

migrate-create: ## Create a new migration file (usage: make migrate-create NAME=add_table)
	@echo "Creating new migration: $(NAME)..."
	@touch migrations/$$(date +%Y%m%d%H%M%S)_$(NAME).up.sql
	@touch migrations/$$(date +%Y%m%d%H%M%S)_$(NAME).down.sql
	@echo "Migration files created"

install-migrate: ## Install migrate CLI tool
	@echo "Installing golang-migrate..."
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "golang-migrate installed to $$(go env GOPATH)/bin/migrate"

.DEFAULT_GOAL := help
