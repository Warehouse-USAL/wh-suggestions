.PHONY: help up-dev run up-prod deploy down build test test-coverage lint format format-fix pr-checks logs clean

.DEFAULT_GOAL := help

COMPOSE_PROD = docker compose -f docker-compose.yml -f docker-compose.prod.yml

help: ## Show available commands
	@echo ""
	@echo "wh-suggestions — available commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""

# ─── Docker targets ───────────────────────────────────────────────────────────

up-dev: ## Start the service and Redis in dev mode
	docker compose up --build

run: up-dev ## Alias for up-dev

up-prod: ## Start services in production mode (requires env vars set)
	$(COMPOSE_PROD) up -d
	@echo "Production services started"

deploy: ## Pull the newest suggestions image and restart it (Redis data untouched)
	$(COMPOSE_PROD) pull suggestions
	$(COMPOSE_PROD) up -d suggestions
	@echo "Suggestions redeployed"

down: ## Stop and remove all containers
	docker compose down

logs: ## Follow service logs
	docker compose logs -f suggestions

# ─── Build / test targets ─────────────────────────────────────────────────────

build: ## Build the binary into bin/
	CGO_ENABLED=0 go build -o bin/wh-suggestions ./cmd/wh-suggestions

clean: ## Remove build artifacts
	rm -rf bin coverage.out

test: ## Run all tests
	go test ./...

test-coverage: ## Run tests and print coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -1

# ─── Quality targets ──────────────────────────────────────────────────────────

lint: ## Run go vet
	go vet ./...

format: ## Check formatting (gofmt)
	@test -z "$$(gofmt -l .)" || { gofmt -l .; echo "Run 'make format-fix'"; exit 1; }

format-fix: ## Apply gofmt to all sources
	gofmt -w .

# ─── CI gate ──────────────────────────────────────────────────────────────────

pr-checks: ## Run all checks required before opening a PR (format + lint + test + build)
	@echo "Running pre-PR checks..."
	$(MAKE) format lint test build
	@echo ""
	@echo "All checks passed. Ready to open a PR."
