# Makefile Usage Guide

This Makefile provides convenient shortcuts for building, testing, and deploying the Trading Bot backend.

## Quick Start

```bash
# View all available commands
make help

# Setup development environment (install required tools)
make setup

# Run the application in development mode
make run

# Build the application
make build

# Run tests
make test

# Run with hot reload (requires air)
make watch
```

## Available Commands

### Building

| Command | Description |
|---------|-------------|
| `make build` | Build the application binary for your current OS |
| `make build-linux` | Build for Linux (amd64) |
| `make build-mac` | Build for macOS |
| `make build-windows` | Build for Windows |

### Running

| Command | Description |
|---------|-------------|
| `make run` | Run application with `go run` (live reload) |
| `make run-build` | Build then run the binary |
| `make watch` | Run with auto-reload using air (requires: `go install github.com/cosmtrek/air@latest`) |

### Testing & Quality

| Command | Description |
|---------|-------------|
| `make test` | Run all unit tests |
| `make test-coverage` | Run tests and generate coverage report (coverage.html) |
| `make lint` | Run golangci-lint checks |
| `make fmt` | Format code with gofmt and goimports |
| `make vet` | Run go vet checks |

### Swagger / API Documentation

| Command | Description |
|---------|-------------|
| `make swagger` | Generate Swagger documentation from code comments |
| `make swagger-install` | Install swag tool (required for swagger generation) |
| `make swagger-clean` | Remove generated Swagger files |

### Docker

| Command | Description |
|---------|-------------|
| `make docker-build` | Build Docker image |
| `make docker-run` | Build and run Docker container locally |
| `make docker-push` | Push Docker image to registry (requires `docker login`) |
| `make docker-clean` | Remove Docker image |

### Dependency Management

| Command | Description |
|---------|-------------|
| `make deps` | Download all dependencies |
| `make deps-tidy` | Clean up unused dependencies |
| `make deps-vendor` | Vendor dependencies into vendor/ directory |

### Cleanup

| Command | Description |
|---------|-------------|
| `make clean` | Clean build artifacts and temp files |
| `make deep-clean` | Clean everything (build, swagger, docker) |

### Utilities

| Command | Description |
|---------|-------------|
| `make env-setup` | Install development tools (golangci-lint, swag, air) |
| `make version` | Show build version and git info |
| `make all` | Run all checks and build (fmt, lint, test, build) |
| `make setup` | Full setup for development |

## Environment Variables

You can customize Docker registry and image names:

```bash
# Push to custom registry
make docker-push DOCKER_REGISTRY=gcr.io DOCKER_IMAGE_NAME=my-project/bot

# Build with custom tag
make docker-build DOCKER_IMAGE_TAG=v1.0.0
```

## Docker Compose

For local development with MySQL:

```bash
# Start all services (app + database)
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop all services
docker-compose down

# Reset database
docker-compose down -v
docker-compose up -d
```

Create a `.env` file for configuration (copy from `.env.example`):

```bash
cp .env.example .env
# Edit .env with your settings
docker-compose up -d
```

## Swagger API Documentation

After running the app:

1. Generate documentation: `make swagger`
2. Run the app: `make run`
3. Visit: http://localhost:3000/swagger/index.html

## Development Workflow

### First Time Setup

```bash
make setup                  # Install tools and dependencies
cp .env.example .env        # Create config
docker-compose up -d        # Start MySQL
make run                    # Start app
```

### Daily Development

```bash
make watch                  # Run with hot reload
# Edit code...
# Changes auto-reload
```

### Before Committing

```bash
make all                    # Run all checks and tests
make fmt                    # Format code
```

### Building for Production

```bash
make build-linux            # Build for Linux server
make docker-build           # Build Docker image
make docker-push            # Push to registry
```

## Troubleshooting

**Build fails with "command not found: go"**
- Ensure Go 1.24+ is installed: `go version`

**`make watch` doesn't work**
- Install air: `go install github.com/cosmtrek/air@latest`

**Swagger generation fails**
- Install swag: `make swagger-install`

**Docker image build fails**
- Ensure Docker is running: `docker ps`
- Check Docker daemon: `sudo systemctl status docker`

**Database connection errors in docker-compose**
- Check MySQL is healthy: `docker-compose ps`
- View logs: `docker-compose logs mysql`
- Rebuild with fresh DB: `docker-compose down -v && docker-compose up -d`

## Make Tips

- Use `make -B target` to force rebuild (ignore timestamps)
- Use `make -n target` to dry-run (see commands without executing)
- Use `make -j4` to parallelize builds (4 jobs)
- View all rules: `make -p`
