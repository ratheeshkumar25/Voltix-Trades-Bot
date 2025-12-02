.PHONY: help build run clean docker-build docker-push swagger test lint

# Variables
APP_NAME=trading-bot
DOCKER_REGISTRY?=docker.io
DOCKER_IMAGE_NAME?=voltix-trading-bot
DOCKER_IMAGE_TAG?=latest
BINARY_NAME=trading-bot
GO_VERSION=$(shell go version | awk '{print $$3}')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(shell date -u +"%Y-%m-%d %H:%M:%S UTC")
LD_FLAGS=-ldflags "-X 'main.Version=$(GIT_COMMIT)' -X 'main.BuildTime=$(BUILD_TIME)'"

# Color output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m # No Color

# Default target
.DEFAULT_GOAL := help

help: ## Display this help screen
	@echo "$(GREEN)=== Trading Bot Backend Make Commands ===$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(GREEN)Variables:$(NC)"
	@echo "  GO_VERSION:         $(GO_VERSION)"
	@echo "  GIT_COMMIT:         $(GIT_COMMIT)"
	@echo "  BUILD_TIME:         $(BUILD_TIME)"
	@echo ""

# ============================================================================
# Build Targets
# ============================================================================

build: ## Build the application binary
	@echo "$(GREEN)Building $(APP_NAME)...$(NC)"
	@mkdir -p bin
	@go build $(LD_FLAGS) -o bin/$(BINARY_NAME) ./cmd/main.go
	@echo "$(GREEN)✓ Build complete: bin/$(BINARY_NAME)$(NC)"

build-linux: ## Build for Linux (amd64)
	@echo "$(GREEN)Building for Linux amd64...$(NC)"
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build $(LD_FLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/main.go
	@echo "$(GREEN)✓ Linux build complete: bin/$(BINARY_NAME)-linux-amd64$(NC)"

build-mac: ## Build for macOS
	@echo "$(GREEN)Building for macOS...$(NC)"
	@mkdir -p bin
	@GOOS=darwin GOARCH=amd64 go build $(LD_FLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/main.go
	@echo "$(GREEN)✓ macOS build complete: bin/$(BINARY_NAME)-darwin-amd64$(NC)"

build-windows: ## Build for Windows
	@echo "$(GREEN)Building for Windows...$(NC)"
	@mkdir -p bin
	@GOOS=windows GOARCH=amd64 go build $(LD_FLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/main.go
	@echo "$(GREEN)✓ Windows build complete: bin/$(BINARY_NAME)-windows-amd64.exe$(NC)"

# ============================================================================
# Run Targets
# ============================================================================

run: ## Run the application
	@echo "$(GREEN)Starting $(APP_NAME)...$(NC)"
	@go run ./cmd/main.go
	@echo "$(GREEN)✓ Application stopped$(NC)"

run-build: build ## Build and run the application
	@echo "$(GREEN)Running $(APP_NAME) from binary...$(NC)"
	@./bin/$(BINARY_NAME)

watch: ## Run with auto-reload (requires air: go install github.com/cosmtrek/air@latest)
	@echo "$(GREEN)Starting $(APP_NAME) with hot reload...$(NC)"
	@which air > /dev/null || (echo "$(RED)air not found. Install with: go install github.com/cosmtrek/air@latest$(NC)" && exit 1)
	@air

# ============================================================================
# Testing & Code Quality
# ============================================================================

test: ## Run unit tests
	@echo "$(GREEN)Running tests...$(NC)"
	@go test -v -cover ./...
	@echo "$(GREEN)✓ Tests complete$(NC)"

test-coverage: ## Run tests with coverage report
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✓ Coverage report: coverage.html$(NC)"

lint: ## Run golangci-lint (requires: golangci-lint)
	@echo "$(GREEN)Running linter...$(NC)"
	@which golangci-lint > /dev/null || (echo "$(RED)golangci-lint not found. Install from: https://golangci-lint.run/$(NC)" && exit 1)
	@golangci-lint run ./...
	@echo "$(GREEN)✓ Linting complete$(NC)"

fmt: ## Format code (gofmt + goimports)
	@echo "$(GREEN)Formatting code...$(NC)"
	@go fmt ./...
	@which goimports > /dev/null && goimports -w . || echo "$(YELLOW)goimports not installed (optional)$(NC)"
	@echo "$(GREEN)✓ Formatting complete$(NC)"

vet: ## Run go vet
	@echo "$(GREEN)Running vet...$(NC)"
	@go vet ./...
	@echo "$(GREEN)✓ Vet check complete$(NC)"

# ============================================================================
# Swagger / API Documentation
# ============================================================================

swagger-install: ## Install swag tool (required for swagger generation)
	@echo "$(GREEN)Installing swag...$(NC)"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "$(GREEN)✓ swag installed$(NC)"

swagger: swagger-install ## Generate Swagger documentation
	@echo "$(GREEN)Generating Swagger documentation...$(NC)"
	@swag init -g cmd/main.go --parseVendor --dir .
	@echo "$(GREEN)✓ Swagger generated in docs/ directory$(NC)"
	@echo "$(YELLOW)Access Swagger UI at: http://localhost:3000/swagger/index.html (when app running)$(NC)"
	@echo "$(YELLOW)Swagger JSON: http://localhost:3000/swagger/doc.json$(NC)"

swagger-clean: ## Clean generated Swagger files
	@echo "$(YELLOW)Removing Swagger files...$(NC)"
	@rm -rf docs/
	@echo "$(GREEN)✓ Swagger files removed$(NC)"

# ============================================================================
# Dependency Management
# ============================================================================

deps: ## Download dependencies
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	@go mod download
	@echo "$(GREEN)✓ Dependencies downloaded$(NC)"

deps-tidy: ## Clean up dependencies
	@echo "$(GREEN)Tidying dependencies...$(NC)"
	@go mod tidy
	@echo "$(GREEN)✓ Dependencies tidied$(NC)"

deps-vendor: ## Vendor dependencies
	@echo "$(GREEN)Vendoring dependencies...$(NC)"
	@go mod vendor
	@echo "$(GREEN)✓ Dependencies vendored$(NC)"

# ============================================================================
# Docker Targets
# ============================================================================

docker-build: ## Build Docker image
	@echo "$(GREEN)Building Docker image: $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)$(NC)"
	@docker build -t $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) \
		--build-arg GO_VERSION=$(GO_VERSION:go%=%) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		-f Dockerfile .
	@echo "$(GREEN)✓ Docker image built successfully$(NC)"

docker-run: docker-build ## Build and run Docker container
	@echo "$(GREEN)Running Docker container...$(NC)"
	@docker run -it --rm \
		-p 3000:3000 \
		-e ENV=development \
		$(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

docker-push: docker-build ## Push Docker image to registry (requires: docker login)
	@echo "$(GREEN)Pushing image to registry...$(NC)"
	@docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
	@echo "$(GREEN)✓ Image pushed to $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)$(NC)"

docker-clean: ## Remove Docker image
	@echo "$(YELLOW)Removing Docker image...$(NC)"
	@docker rmi $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) 2>/dev/null || true
	@echo "$(GREEN)✓ Docker image removed$(NC)"

# ============================================================================
# Cleanup Targets
# ============================================================================

clean: ## Clean build artifacts and temporary files
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rf bin/
	@rm -rf coverage.out coverage.html
	@go clean -cache -testcache
	@echo "$(GREEN)✓ Clean complete$(NC)"

deep-clean: clean swagger-clean docker-clean ## Deep clean (build artifacts, swagger, docker)
	@echo "$(GREEN)✓ Deep clean complete$(NC)"

# ============================================================================
# Development Helpers
# ============================================================================

env-setup: ## Setup development environment (install tools)
	@echo "$(GREEN)Setting up development environment...$(NC)"
	@echo "Installing golangci-lint..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Installing swag..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Installing air (hot reload)..."
	@go install github.com/cosmtrek/air@latest
	@echo "$(GREEN)✓ Development environment setup complete$(NC)"

version: ## Show build version info
	@echo "$(GREEN)Build Information:$(NC)"
	@echo "  Binary:       $(BINARY_NAME)"
	@echo "  Go Version:   $(GO_VERSION)"
	@echo "  Git Commit:   $(GIT_COMMIT)"
	@echo "  Build Time:   $(BUILD_TIME)"

# ============================================================================
# Database Targets (if using migrations)
# ============================================================================

db-migrate: ## Run database migrations (when implemented)
	@echo "$(GREEN)Running database migrations...$(NC)"
	@go run ./cmd/main.go migrate
	@echo "$(GREEN)✓ Migrations complete$(NC)"

# ============================================================================
# All-in-one targets
# ============================================================================

all: clean deps fmt lint test build ## Run all checks and build
	@echo "$(GREEN)✓ All tasks complete$(NC)"

setup: env-setup deps swagger ## Full setup for development
	@echo "$(GREEN)✓ Setup complete. Run 'make run' to start development$(NC)"

.PHONY: all setup
