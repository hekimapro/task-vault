.PHONY: dev build run clean tidy migrate migrate-down migrate-status migrate-create vendor

# Load .env variables
include .env
export

dev:
	@echo "Running pending migrations..."
	goose -dir migrations postgres "$(DATABASE_URL)" up
	@echo "Starting development server with Air..."
	@air

build:
	@echo "Building binary..."
	go build -o bin/task-vault ./cmd

run:
	@echo "Running binary..."
	./bin/task-vault

clean:
	@echo "Cleaning up..."
	rm -rf bin/ tmp/

tidy:
	@echo "Tidying Go modules..."
	go mod tidy

vendor:
	@echo "Vendoring Go Modules..."
	go mod vendor

migrate:
	@echo "Running goose up migrations..."
	goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
	@echo "Rolling back last goose migration..."
	goose -dir migrations postgres "$(DATABASE_URL)" down

migrate-status:
	@echo "Showing migration status..."
	goose -dir migrations postgres "$(DATABASE_URL)" status

migrate-create:
	@read -p "Enter table name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	filename="migrations/$${timestamp}_create_$${name}_table.sql"; \
	echo "Creating migration: $$filename"; \
	echo "-- +goose Up" > $$filename; \
	echo "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";" >> $$filename; \
	echo "" >> $$filename; \
	echo "CREATE TABLE $$name (" >> $$filename; \
	echo "    id UUID PRIMARY KEY DEFAULT uuid_generate_v4()," >> $$filename; \
	echo "    created_at TIMESTAMPTZ DEFAULT NOW()," >> $$filename; \
	echo "    updated_at TIMESTAMPTZ DEFAULT NOW()" >> $$filename; \
	echo ");" >> $$filename; \
	echo "" >> $$filename; \
	echo "-- +goose Down" >> $$filename; \
	echo "DROP TABLE IF EXISTS $$name CASCADE;" >> $$filename; \
	echo "Migration created at: $$filename"
