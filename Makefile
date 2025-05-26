include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

build:
	tailwindcss -i views/css/styles.css -o public/styles.css
	@templ generate view
	@go build -o bin/fullstackgo cmd/main.go 

test:
	@go test -v ./...
	
run: build
	@./bin/fullstackgo

tailwind:
	@tailwindcss -i views/css/styles.css -o public/styles.css --watch

templ:
	@templ generate -watch -proxy=http://localhost:8080

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))