.PHONY: help test lint fmt docker-up docker-down migrate-up migrate-down

help:
	@echo "BeatFlow development commands"
	@echo ""
	@echo "Usage:"
	@echo "  make help           Show available commands"
	@echo "  make test           Run Go tests for initialized modules"
	@echo "  make lint           Run lint checks placeholder"
	@echo "  make fmt            Format Go code for initialized modules"
	@echo "  make docker-up      Start local Docker Compose services"
	@echo "  make docker-down    Stop local Docker Compose services"
	@echo "  make migrate-up     Run database migrations placeholder"
	@echo "  make migrate-down   Roll back database migrations placeholder"

test:
	@echo "Running Go tests..."
	@for dir in apps/auth-service packages/go/config packages/go/logger; do \
		if find $$dir -name '*.go' | grep -q .; then \
			echo "Testing $$dir"; \
			(cd $$dir && go test ./...); \
		else \
			echo "Skipping $$dir: no Go files yet"; \
		fi; \
	done

lint:
	@echo "lint placeholder: golangci-lint will be added later"

fmt:
	@echo "Formatting Go files..."
	@if find apps packages -name '*.go' | grep -q .; then \
		gofmt -w $$(find apps packages -name '*.go'); \
	else \
		echo "No Go files to format yet"; \
	fi

docker-up:
	docker compose -f deployments/docker-compose/docker-compose.yml up -d

docker-down:
	docker compose -f deployments/docker-compose/docker-compose.yml down

migrate-up:
	@echo "migrate-up placeholder: Goose migrations will be added later"

migrate-down:
	@echo "migrate-down placeholder: Goose migrations will be added later"