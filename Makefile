# .PHONY tells Make that these are not filenames but tasks
.PHONY: build test lint docker run

include .env

# Default target: run the application
run: postgres postgres_ready go_run

# Build the Go application
build:
	@echo "Building the Go application..."
	go build -o api ./cmd/api


# Run tests
test:
	@echo "Running Go tests..."
	go test ./...


# Set up the PostgreSQL container (Docker setup)
postgres:
	@echo "Checking if PostgreSQL container exists..."
	@if ! docker ps -q -f name=postgresdb; then \
		echo "PostgreSQL container not found. Creating a new one..."; \
		docker run --name postgresdb -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p 5432:5432 -d postgres:latest; \
	else \
		echo "PostgreSQL container already running."; \
	fi

# Ensure PostgreSQL is ready before running the application
postgres_ready:
	@echo "Waiting for PostgreSQL to start..."
	@until docker exec postgresdb pg_isready -U postgres; do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 1; \
	done

# Run the Go application
go_run:
	@echo "Running the Go application..."
	go run ./cmd/api
