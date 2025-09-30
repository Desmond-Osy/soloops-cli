# SoloOps CLI Makefile
.PHONY: build test install clean lint format help docker-build

VERSION ?= dev
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildDate=$(BUILD_DATE)"

# Build the CLI binary
build:
	@echo "Building SoloOps CLI..."
	go build $(LDFLAGS) -o bin/soloops ./cmd/soloops

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/soloops-linux-amd64 ./cmd/soloops
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/soloops-linux-arm64 ./cmd/soloops
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/soloops-darwin-amd64 ./cmd/soloops
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/soloops-darwin-arm64 ./cmd/soloops
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/soloops-windows-amd64.exe ./cmd/soloops
	@echo "Binaries built in bin/"

# Run tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage report
test-coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Install the CLI locally
install:
	@echo "Installing SoloOps CLI..."
	go install $(LDFLAGS) ./cmd/soloops

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

# Lint the code
lint:
	@echo "Running linters..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run ./...

# Format the code
format:
	@echo "Formatting code..."
	go fmt ./...
	@which goimports > /dev/null && goimports -w . || echo "goimports not found, skipping import formatting"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod verify

# Update dependencies
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t soloops:$(VERSION) .

# Run Docker container
docker-run:
	docker run --rm -it soloops:$(VERSION)

# Show help
help:
	@echo "SoloOps CLI - Makefile commands:"
	@echo ""
	@echo "  make build          - Build the CLI binary"
	@echo "  make build-all      - Build binaries for all platforms"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make install        - Install the CLI locally"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make lint           - Run linters"
	@echo "  make format         - Format code"
	@echo "  make deps           - Download dependencies"
	@echo "  make deps-update    - Update dependencies"
	@echo "  make docker-build   - Build Docker image"
	@echo "  make docker-run     - Run Docker container"
	@echo "  make help           - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  VERSION             - Version to build (default: dev)"

.DEFAULT_GOAL := build